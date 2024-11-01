package service

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/sha3"
	"shrt/pkg/config"
)

type shrtService struct {
	cfg   config.Config
	store Store
}

var _ ShrtService = &shrtService{}

func NewShrtService(cfg config.Config) ShrtService {
	return &shrtService{cfg, NewInMemoryStore()}
}

func (s *shrtService) Shorten(url string) (string, error) {
	digest := make([]byte, 16)
	sha3.ShakeSum128(digest, []byte(url))
	encodedLength := base64.RawURLEncoding.EncodedLen(16)
	encoded := make([]byte, encodedLength)
	base64.RawURLEncoding.Encode(encoded, digest)
	encoded = encoded[:8]

	shortUrl := fmt.Sprintf("%s/%s", s.cfg.BaseUrl, string(encoded))

	err := s.store.Save(string(encoded), url)
	if err != nil {
		return "", err
	}

	return shortUrl, nil
}

func (s *shrtService) Unshorten(uuid string) (string, error) {
	url, err := s.store.Load(uuid)
	if err != nil {
		return "", err
	}

	return url, nil
}
