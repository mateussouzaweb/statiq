# STATIQ - Simple WebServer for Static Files

*Statiq* gives you a simple and efficient server to develop and deploy web projects containing static files only like Websites, PWA, SPA, etc... fair and simple, zero configurations.

In production, this project aims to be used behind a reverse proxy (like *traefik*, *cloudflare*, ...), where you will handle advanced configurations like HTTPS, HTTP Auth, ACL, and more.

---

## Features

- Written in Go Language
- Based on ``labstack/echo`` project - <https://echo.labstack.com/>
- Includes Logs, Recovery, Request ID and CORS
- Contains basic security features like XSS, Content Type Sniffing, X-Frame Options, ...
- Auto rewrite to SPA project
- Auto remove trailing slash
- Gzip enabled

---

## Installation and Usage

To install, just download the binary file and place it on the binaries folder:

```bash
sudo wget https://raw.githubusercontent.com/mateussouzaweb/statiq/master/bin/statiq -O /usr/local/bin/statiq && sudo chmod +x /usr/local/bin/statiq
```

To check command flags, use:

```bash
statiq --help
```

To start the server, run:

```bash
statiq --port 8080 --root /path/to/root/server/
```

Kill the process when you need and that is it!
