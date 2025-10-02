package web

import (
	"example.com/local/Go2part/internal/web/obankaccount"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"example.com/local/Go2part/internal/cache"
	"example.com/local/Go2part/internal/legalentities"
	"example.com/local/Go2part/internal/web/olegalentity"
	rds "example.com/local/Go2part/pkg/redis"
)

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	// health
	r.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })

	// --- Redis + Cache wiring ---
	addr := env("REDIS_ADDR", "127.0.0.1:6300")
	dbStr := env("REDIS_DB", "0")
	pw := env("REDIS_PASSWORD", "")
	db, _ := strconv.Atoi(dbStr)

	rd, err := rds.New(addr, db, pw)
	if err == nil && rd != nil {
		// ок
	} else {
		rd = nil // отключим кеш, если редис недоступен
	}

	ttl := 5 * time.Minute
	cacheSvc := cache.NewService(rd, ttl, ttl)

	// Сервис юрлиц с кешем
	leSvc := legalentities.NewService(cacheSvc)

	// Регистрация CRUD хендлеров
	olegalentity.NewLegalEntityHandler(leSvc).Register(r)
	obankaccount.NewHandler(leSvc).Register(r)

	return r
}
