package server

import (
	"gold/src/configs"
	"gold/src/logger"
	"gold/src/router"
	"net/http"
)

var conf = configs.Get()
var log = logger.Get()

// Start : Starts the HTTP Server.
func Start() error {
	log.Sugar().Infof(
		"Starting %s@%s HTTP Server at %s:%s",
		conf.Application.Name,
		conf.Application.Version,
		conf.Server.Host,
		conf.Server.Port,
	)

	return http.ListenAndServe(conf.Server.Host+":"+conf.Server.Port, router.Get())
}
