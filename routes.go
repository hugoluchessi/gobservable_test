package main

import (
	"net/http"

	"github.com/hugoluchessi/badger"
	"github.com/hugoluchessi/gobservable/metrics"
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

	reqCountMw := metrics.NewRequestCountMiddleware(ms.MetricService)
	reqTimeMw := metrics.NewRequestTimeMiddleware(ms.MetricService)

	mainRouter.Use(reqCountMw.Handler)
	mainRouter.Use(reqTimeMw.Handler)

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
