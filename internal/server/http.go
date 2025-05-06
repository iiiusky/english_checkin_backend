package server

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	nethttp "net/http"
	"english_checkin_backend/docs"
	"english_checkin_backend/internal/handler"
	"english_checkin_backend/internal/middleware"
	"english_checkin_backend/pkg/jwt"
	"english_checkin_backend/pkg/log"
	"english_checkin_backend/pkg/server/http"
	"english_checkin_backend/web"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	e *casbin.SyncedEnforcer,
	adminHandler *handler.AdminHandler,
	userHandler *handler.UserHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)
	// 设置前端静态资源
	s.Use(static.Serve("/", static.EmbedFolder(web.Assets(), "dist")))
	s.NoRoute(func(c *gin.Context) {
		indexPageData, err := web.Assets().ReadFile("dist/index.html")
		if err != nil {
			c.String(nethttp.StatusNotFound, "404 page not found")
			return
		}
		c.Data(nethttp.StatusOK, "text/html; charset=utf-8", indexPageData)
	})
	// swagger doc
	docs.SwaggerInfo.BasePath = "/"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)

	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.POST("/login", adminHandler.Login)
		}

		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger), middleware.AuthMiddleware(e))
		{
			strictAuthRouter.GET("/users", userHandler.GetUsers)

			strictAuthRouter.GET("/menus", adminHandler.GetMenus)
			strictAuthRouter.GET("/admin/menus", adminHandler.GetAdminMenus)
			strictAuthRouter.POST("/admin/menu", adminHandler.MenuCreate)
			strictAuthRouter.PUT("/admin/menu", adminHandler.MenuUpdate)
			strictAuthRouter.DELETE("/admin/menu", adminHandler.MenuDelete)

			strictAuthRouter.GET("/admin/users", adminHandler.GetAdminUsers)
			strictAuthRouter.GET("/admin/user", adminHandler.GetAdminUser)
			strictAuthRouter.PUT("/admin/user", adminHandler.AdminUserUpdate)
			strictAuthRouter.POST("/admin/user", adminHandler.AdminUserCreate)
			strictAuthRouter.DELETE("/admin/user", adminHandler.AdminUserDelete)
			strictAuthRouter.GET("/admin/user/permissions", adminHandler.GetUserPermissions)
			strictAuthRouter.GET("/admin/role/permissions", adminHandler.GetRolePermissions)
			strictAuthRouter.PUT("/admin/role/permission", adminHandler.UpdateRolePermission)
			strictAuthRouter.GET("/admin/roles", adminHandler.GetRoles)
			strictAuthRouter.POST("/admin/role", adminHandler.RoleCreate)
			strictAuthRouter.PUT("/admin/role", adminHandler.RoleUpdate)
			strictAuthRouter.DELETE("/admin/role", adminHandler.RoleDelete)

			strictAuthRouter.GET("/admin/apis", adminHandler.GetApis)
			strictAuthRouter.POST("/admin/api", adminHandler.ApiCreate)
			strictAuthRouter.PUT("/admin/api", adminHandler.ApiUpdate)
			strictAuthRouter.DELETE("/admin/api", adminHandler.ApiDelete)

		}
	}
	return s
}
