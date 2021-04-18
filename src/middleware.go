package main

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

// Config struct
type Config struct {
	Port  int64
	Root  string
	Index string
	HTML5 bool
}

// filesMiddleware handle file server
func filesMiddleware(config *Config) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			p := c.Request().URL.Path
			p, err := url.PathUnescape(p)

			if err != nil {
				return err
			}

			// Determine resource
			name := filepath.Join(config.Root, filepath.Clean("/"+p)) // "/"+ for security
			isHTML := strings.Contains(p, ".htm")

			fi, err := os.Stat(name)

			if err != nil {

				if os.IsNotExist(err) && config.HTML5 {
					name = filepath.Join(config.Root, config.Index)
					isHTML = true
				} else {
					return err
				}

			} else if fi.IsDir() {

				name = filepath.Join(name, config.Index)
				isHTML = true

				_, err = os.Stat(name)

				if err != nil {
					if os.IsNotExist(err) {
						return next(c)
					}
					return err
				}

			}

			// Security headers
			c.Response().Header().Set(
				echo.HeaderXContentTypeOptions,
				"nosniff",
			)

			if isHTML {
				c.Response().Header().Set(
					echo.HeaderXXSSProtection,
					"1; mode=block",
				)
				c.Response().Header().Set(
					echo.HeaderContentSecurityPolicy,
					"frame-ancestors",
				)
			}

			// Mimetypes && UTF-8 header
			for format, mime := range mimes {
				if strings.HasSuffix(name, format) {
					c.Response().Header().Set(
						echo.HeaderContentType,
						mime,
					)
					break
				}
			}

			// Cache control
			if !isHTML {
				c.Response().Header().Set(
					"Cache-Control",
					"max-age=31536000, immutable",
				)
			} else {
				c.Response().Header().Set(
					"Cache-Control",
					"max-age=180, private",
				)
			}

			return c.File(name)
		}
	}
}
