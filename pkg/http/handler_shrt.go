package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shrt/pkg/service"
	"strings"
)

func Shrt(s service.ShrtService) func(c echo.Context) error {
	return func(c echo.Context) error {
		switch p := c.Param("something"); {
		case p == "":
			return c.String(http.StatusBadRequest, "url is required")
		case isUrl(p):
			shortUrl, err := s.Shorten(p)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			return c.String(http.StatusOK, shortUrl)
		default:
			url, err := s.Unshorten(p)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			return c.Redirect(http.StatusMovedPermanently, url)
		}
	}
}

func isUrl(s string) bool {
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")
}
