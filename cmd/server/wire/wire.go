//go:build wireinject
// +build wireinject

package wire

import (
	"english_checkin_backend/internal/handler"
	"english_checkin_backend/internal/job"
	"english_checkin_backend/internal/repository"
	"english_checkin_backend/internal/server"
	"english_checkin_backend/internal/service"
	"english_checkin_backend/pkg/app"
	"english_checkin_backend/pkg/jwt"
	"english_checkin_backend/pkg/log"
	"english_checkin_backend/pkg/server/http"
	"english_checkin_backend/pkg/sid"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
	repository.NewCasbinEnforcer,
	repository.NewAdminRepository,
	repository.NewStudentRepository,
	repository.NewStudentProgressRepository,
	repository.NewTaskRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewAdminService,
	service.NewStudentService,
	service.NewStudentProgressService,
	service.NewTaskService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewAdminHandler,
	handler.NewStudentHandler,
	handler.NewStudentProgressHandler,
	handler.NewTaskHandler,
)

var jobSet = wire.NewSet(
	job.NewJob,
	job.NewUserJob,
)
var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJobServer,
)

// build App
func newApp(
	httpServer *http.Server,
	jobServer *server.JobServer,
	// task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, jobServer),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		jobSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
