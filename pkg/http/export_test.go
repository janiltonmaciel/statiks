package http

var (
	NoCacheHandler       = noCacheHandler
	CacheHandler         = cacheHandler
	DelayHandler         = delayHandler
	WriteNotModified     = writeNotModified
	ToHTTPError          = toHTTPError
	CheckIfModifiedSince = checkIfModifiedSince
)
