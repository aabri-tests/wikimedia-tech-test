package ports

type WikiMediaFactory interface {
	GetWikiMediaAdapter() (WikiMedia, error)
}

type LoggerFactory interface {
	GetLogger() (LogInfoFormat, error)
}

type CacheFactory interface {
	GetCache() (Cache, error)
}
