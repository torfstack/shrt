package http

import (
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/torfstack/shrt/pkg/service"
	"net/http"
	"strings"
)

func Shrt(s service.ShrtService) func(c echo.Context) error {
	return func(c echo.Context) error {
		switch p := c.Param("something"); {
		case p == "":
			return c.String(http.StatusBadRequest, "url is required")
		case isBase64UrlEncoded(p):
			url, err := s.Unshorten(p)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			return c.Redirect(http.StatusMovedPermanently, url)
		default:
			shortUrl, err := s.Shorten(sanitizeUrl(p))
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			return c.String(http.StatusOK, shortUrl)
		}
	}
}

func isBase64UrlEncoded(s string) bool {
	_, err := base64.RawURLEncoding.DecodeString(s)
	return err == nil
}

func sanitizeUrl(s string) string {
	if strings.HasPrefix(s, "http://") ||
		strings.HasPrefix(s, "https://") {
		return s
	}

	_, err := http.DefaultClient.Get(fmt.Sprintf("https://%s", s))
	if err == nil {
		return fmt.Sprintf("https://%s", s)
	}

	return fmt.Sprintf("http://%s", s)
}
