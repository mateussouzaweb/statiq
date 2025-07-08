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
	SPA   bool
}

// FilesMiddleware handle file server
func FilesMiddleware(config *Config) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			p := c.Request().URL.Path
			p, err := url.PathUnescape(p)

			if err != nil {
				return err
			}

			// Determine resource
			// Append "/" + for security reasons
			name := filepath.Join(config.Root, filepath.Clean("/"+p))
			isHTML := strings.Contains(p, ".htm")
			fileInfo, err := os.Stat(name)

			if err != nil {

				// On SPA model, rewrite to index page
				if os.IsNotExist(err) && config.SPA {
					name = filepath.Join(config.Root, config.Index)
					isHTML = true
				} else {
					return err
				}

			} else if fileInfo.IsDir() {

				// If it's a directory, append index file and check if it exists
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

			// Add security headers
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

			// Set Mimetype && UTF-8 header
			for format, mime := range mimes {
				if strings.HasSuffix(name, format) {
					c.Response().Header().Set(
						echo.HeaderContentType,
						mime,
					)
					break
				}
			}

			// Set cache control
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
