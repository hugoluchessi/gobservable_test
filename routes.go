package main

import (
	"net/http"
	"time"

	"github.com/hugoluchessi/badger"
	prom "github.com/hugoluchessi/gobservable/metrics/providers/prometheus"
	"github.com/hugoluchessi/gobservable_test/config"
	"github.com/hugoluchessi/gobservable_test/controllers"

	log "github.com/hugoluchessi/gobservable/logging"
)

func ConfigureRoutes(ms *config.MonitorServices) *badger.Mux {
	// Create new Mux
	mux := badger.NewMux()

	// Create new router group
	mainRouter := mux.AddRouter("/")

	// Metrics middleware idea #1
	mainRouter.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			ms.MetricService.IncrCounter([]string{"req", "c"}, 1)
			h.ServeHTTP(res, req)
		})
	})

	// Metrics middleware idea #2
	mainRouter.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			start := time.Now()
			h.ServeHTTP(res, req)
			ms.MetricService.MeasureSince([]string{"req", "t"}, start)
		})
	})

	loggerMw := log.NewContextLoggerMiddleware(ms.ContextLogger)

	mainRouter.Use(loggerMw.Handler)

	mainRouter.Get("", http.HandlerFunc(healthCheckHandler))

	promSink := ms.MetricService.Sink.(*prom.Sink)
	mainRouter.Get("metrics", prom.HTTPHandlerFor(promSink))

	mainRouter.Get("someDomain", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		controllers.HandleSomeDomainGetSomething(ms, res, req)
	}))

	return mux
}

func healthCheckHandler(res http.ResponseWriter, req *http.Request) {
	// Check external dependencies and return 500 in case anything goes wrong
}
