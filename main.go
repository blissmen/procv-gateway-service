package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
	"gopkg.in/yaml.v3"
)

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

var (
	registry map[string]Service
	breaker  *gobreaker.CircuitBreaker
)

func main() {
	loadRegistry("config/services.yaml")

	breaker = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "ServiceBreaker",
		MaxRequests: 3,
		Interval:    60 * time.Second,
		Timeout:     30 * time.Second,
	})

	r := gin.Default()
	r.Any("/api/:service/*path", proxyHandler)
	r.Run(":8080")
}

func loadRegistry(path string) {
	absPath, _ := filepath.Abs(path)
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Fatalf("Failed to read registry file: %v", err)
	}

	err = yaml.Unmarshal(data, &registry)
	if err != nil {
		log.Fatalf("Failed to parse registry file: %v", err)
	}

	log.Println("Service registry loaded")
}

func proxyHandler(c *gin.Context) {
	service := c.Param("service")
	path := c.Param("path")

	targetService, ok := registry[service]
	if !ok {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Unknown service"})
		return
	}

	targetURL := "http://" + targetService.Host + ":" + strconv.Itoa(targetService.Port)
	target, err := url.Parse(targetURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid service address"})
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.URL.Path = path
		req.Host = target.Host
	}

	breaker.Execute(func() (interface{}, error) {
		proxy.ServeHTTP(c.Writer, c.Request)
		return nil, nil
	})
}
