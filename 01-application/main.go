package main

import (
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"
)

// getOutboundIP возвращает «локальный» IP-адрес хоста.
// Это адрес интерфейса, через который хост ходит во внешнюю сеть.
// Если определить не удалось — возвращаем "unknown".
func getOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "unknown"
	}
	defer conn.Close()

	localAddr, ok := conn.LocalAddr().(*net.UDPAddr)
	if !ok {
		return "unknown"
	}
	return localAddr.IP.String()
}

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	author := os.Getenv("AUTHOR")
	if author == "" {
		author = "unknown"
	}

	ip := getOutboundIP()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// Главная страница с информацией
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `<!DOCTYPE html>
<html lang="ru">
<head>
  <meta charset="utf-8">
  <title>Info</title>
  <style>
    body { font-family: sans-serif; margin: 40px; }
    table { border-collapse: collapse; }
    td { padding: 6px 14px; border-bottom: 1px solid #ddd; }
    td:first-child { font-weight: bold; }
  </style>
</head>
<body>
  <h1>Информация о хосте</h1>
  <table>
    <tr><td>Hostname</td><td>%s</td></tr>
    <tr><td>IP-адрес</td><td>%s</td></tr>
    <tr><td>Author</td><td>%s</td></tr>
  </table>
</body>
</html>`, hostname, ip, author)
	})

	// Kubernetes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.GET("/live", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.GET("/ready", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Printf("Starting server on :8080 (hostname=%s, ip=%s, author=%s)", hostname, ip, author)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server error: %v", err)
	}
}