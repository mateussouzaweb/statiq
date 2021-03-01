package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Command line flags
	portFlag := flag.Int64("port", 8080, "Port to the server")
	rootFlag := flag.String("root", "/var/www/", "Root folder where the files are")

	// Parse values
	flag.Parse()

	// Basic validation
	if *portFlag < 1 {
		log.Fatal("ERROR: Missing port flag")
	}

	if *rootFlag == "" {
		log.Fatal("ERROR: Missing root path")
	}

	// Create server
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	// Static files + SPA rewrite
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  *rootFlag,
		HTML5: true,
	}))

	// Start server
	port := fmt.Sprintf(":%d", *portFlag)
	err := e.Start(port)

	if err != nil {
		log.Fatal(err)
	}

}
