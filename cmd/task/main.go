package main

import (
	"context"
	"flag"
	"english_checkin_backend/cmd/task/wire"
	"english_checkin_backend/pkg/config"
	"english_checkin_backend/pkg/log"
)

func main() {
	var envConf = flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger := log.NewLog(conf)
	logger.Info("start task")
	app, cleanup, err := wire.NewWire(conf, logger)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}

}
