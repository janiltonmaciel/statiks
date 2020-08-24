package lib

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"time"
)

const (
	rootName = "rootCA.pem"
	keyName  = "rootCA-key.pem"
)

var userAndHostname string

// nolint
func init() {
	u, err := user.Current()
	if err == nil {
		userAndHostname = u.Username + "@"
	}
	if h, err := os.Hostname(); err == nil {
		userAndHostname += h
	}
	if err == nil && u.Name != "" && u.Name != u.Username {
		userAndHostname += " (" + u.Name + ")"
	}
}

func GetMkCert(host string) (certArray []byte, keyArray []byte) {
	carootPath := getCAROOT()
	caCert, caKey := loadCA(carootPath)

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	fatalIfErr(err, "failed to generate certificate key")

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	fatalIfErr(err, "failed to generate serial number")

	tpl := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization:       []string{"mkcert development certificate"},
			OrganizationalUnit: []string{userAndHostname},
		},

		NotAfter: time.Now().AddDate(10, 0, 0),

		// Fix the notBefore to temporarily bypass macOS Catalina's limit on
		// certificate lifespan. Once mkcert provides an ACME server, automation
		// will be the recommended way to guarantee uninterrupted functionality,
		// and the lifespan will be shortened to 825 days. See issue 174 and
		// https://support.apple.com/en-us/HT210176.
		NotBefore: time.Date(2019, time.June, 1, 0, 0, 0, 0, time.UTC),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	if ip := net.ParseIP(host); ip != nil {
		tpl.IPAddresses = append(tpl.IPAddresses, ip)
	} else {
		tpl.DNSNames = append(tpl.DNSNames, host)
	}

	pub := priv.Public()
	cert, err := x509.CreateCertificate(rand.Reader, tpl, caCert, pub, caKey)
	fatalIfErr(err, "failed to generate certificate")

	privDER, err := x509.MarshalPKCS8PrivateKey(priv)
	fatalIfErr(err, "failed to marshal private key")

	keyArray = pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privDER})
	certArray = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert})

	return certArray, keyArray
}

func getCAROOT() string {
	if env := os.Getenv("CAROOT"); env != "" {
		return env
	}

	var dir string
	switch {
	case runtime.GOOS == "windows":
		dir = os.Getenv("LocalAppData")
	case os.Getenv("XDG_DATA_HOME") != "":
		dir = os.Getenv("XDG_DATA_HOME")
	case runtime.GOOS == "darwin":
		dir = os.Getenv("HOME")
		if dir == "" {
			return ""
		}
		dir = filepath.Join(dir, "Library", "Application Support")
	default: // Unix
		dir = os.Getenv("HOME")
		if dir == "" {
			return ""
		}
		dir = filepath.Join(dir, ".local", "share")
	}
	return filepath.Join(dir, "mkcert")
}

// loadCA will load or create the CA at CAROOT.
func loadCA(carootPath string) (caCert *x509.Certificate, caKey crypto.PrivateKey) {
	var err error
	var certPEMBlock []byte
	var keyPEMBlock []byte

	f := filepath.Join(carootPath, rootName)
	if _, err = os.Stat(f); os.IsNotExist(err) {
		logger.Print("ERROR: failed to find the default CA location, set one as the CAROOT env var")
	} else {
		logger.Printf("Using the local CA at \"%s\" ✨", carootPath)
	}

	certPEMBlock, err = ioutil.ReadFile(f)
	fatalIfErr(err, "failed to read the CA certificate")
	certDERBlock, _ := pem.Decode(certPEMBlock)
	if certDERBlock == nil || certDERBlock.Type != "CERTIFICATE" {
		logger.Fatalln("ERROR: failed to read the CA certificate: unexpected content")
	}
	caCert, err = x509.ParseCertificate(certDERBlock.Bytes)
	fatalIfErr(err, "failed to parse the CA certificate")

	f = filepath.Join(carootPath, keyName)
	if _, err = os.Stat(f); os.IsNotExist(err) {
		return // keyless mode, where only -install works
	}

	keyPEMBlock, err = ioutil.ReadFile(f)
	fatalIfErr(err, "failed to read the CA key")
	keyDERBlock, _ := pem.Decode(keyPEMBlock)
	if keyDERBlock == nil || keyDERBlock.Type != "PRIVATE KEY" {
		logger.Fatalln("ERROR: failed to read the CA key: unexpected content")
	}
	caKey, err = x509.ParsePKCS8PrivateKey(keyDERBlock.Bytes)
	fatalIfErr(err, "failed to parse the CA key")

	return caCert, caKey
}

func fatalIfErr(err error, msg string) {
	if err != nil {
		logger.Fatalf("ERROR: %s: %s", msg, err)
	}
}
