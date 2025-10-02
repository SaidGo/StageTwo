//go:build disable_extras

package web

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	legalEntitiesRequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "legal_entities_requests_total",
		Help: "Total number of HTTP requests to /legal-entities (and /legal_entities) endpoints",
	})
	bankAccountsRequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "bank_accounts_requests_total",
		Help: "Total number of HTTP requests to /bank-accounts (and /bank_accounts) endpoints",
	})
)

func init() {
	prometheus.MustRegister(legalEntitiesRequestsTotal, bankAccountsRequestsTotal)
}

func isLegalEntitiesPath(p string) bool {
	return strings.HasPrefix(p, "/legal-entities") || strings.HasPrefix(p, "/legal_entities")
}

func isBankAccountsPath(p string) bool {
	return strings.HasPrefix(p, "/bank-accounts") || strings.HasPrefix(p, "/bank_accounts")
}

// MetricsByPath — middleware. Инкремент только для валидно сопоставленного маршрута (исключаем 404).
func MetricsByPath() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // дождаться сопоставления маршрута и выполнения хендлера

		path := c.FullPath()
		if path == "" {
			// нерегистрированный маршрут (404) — не считаем
			return
		}

		if isLegalEntitiesPath(path) {
			legalEntitiesRequestsTotal.Inc()
		} else if isBankAccountsPath(path) {
			bankAccountsRequestsTotal.Inc()
		}
	}
}
