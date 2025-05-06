//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"english_checkin_backend/internal/repository"
	"english_checkin_backend/internal/server"
	"english_checkin_backend/pkg/app"
	"english_checkin_backend/pkg/log"
	"english_checkin_backend/pkg/sid"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewCasbinEnforcer,
)
var serverSet = wire.NewSet(
	server.NewMigrateServer,
)

// build App
func newApp(
	migrateServer *server.MigrateServer,
) *app.App {
	return app.NewApp(
		app.WithServer(migrateServer),
		app.WithName("demo-migrate"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serverSet,
		sid.NewSid,
		newApp,
	))
}
