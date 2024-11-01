package service

type ShrtService interface {
	Shorten(url string) (string, error)
	Unshorten(url string) (string, error)
}

type Store interface {
	Save(key string, value string) error
	Load(key string) (string, error)
	Contains(key string) (bool, error)
	Delete(key string) error
}
