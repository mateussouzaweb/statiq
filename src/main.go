package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config := Config{
		Port:  8080,
		Root:  "/var/www/",
		Index: "index.html",
		SPA:   true,
	}

	// Command line flags
	flag.Int64Var(
		&config.Port,
		"port",
		config.Port,
		"Port to the server",
	)
	flag.StringVar(
		&config.Root,
		"root",
		config.Root,
		"Root folder where the files are",
	)

	// Parse values
	version := flag.Bool("version", false, "Print program version")
	flag.Parse()

	if *version {
		fmt.Println("STATIQ version 0.0.3")
		return
	}

	// Basic validation
	if config.Port < 1 {
		log.Fatal("ERROR: Missing port flag")
	}

	if config.Root == "" {
		log.Fatal("ERROR: Missing root path")
	}

	// Create server
	e := echo.New()

	// Default middlewares
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())

	// Static files + SPA rewrite
	e.Use(FilesMiddleware(&config))

	// Start server
	port := fmt.Sprintf(":%d", config.Port)
	err := e.Start(port)

	if err != nil {
		log.Fatal(err)
	}

}
