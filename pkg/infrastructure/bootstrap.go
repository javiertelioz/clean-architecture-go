package infrastructure

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"github.com/javiertelioz/clean-architecture-go/config"
	"github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/hello"
	domainService "github.com/javiertelioz/clean-architecture-go/pkg/domain/services"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/logger"
	customMiddleware "github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/middleware"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/repository"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/routes"
	infrastructureService "github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/services"
)

type AppDependencies struct {
	Router *chi.Mux
}

// Bootstrap godoc
// @title						Swagger Clean Architecture Go
// @version					1.0
// @description				This is a sample. You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/)
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.url				http://www.swagger.io/support
// @contact.email				support@docs.io
//
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host						localhost:8080
// @BasePath					/
// @Schemes					http https
//
// @securityDefinitions.apikey	bearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and the access token.
//
// @accept						json
// @produce					json
//
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func Bootstrap(db *gorm.DB) *AppDependencies {
	appConfig, _ := config.GetConfig[string]("AppName")
	profilingConfig, _ := config.GetConfig[config.ProfilingConfig]("Profiling")
	salt, _ := config.GetConfig[int]("Crypto.salt")

	sayHelloUseCase := hello.NewSayHelloUseCase()

	loggerService := logger.NewLogger()
	userRepository := repository.NewUserRepository(db)
	cryptoService := infrastructureService.NewBcryptService(salt)
	userService := domainService.NewUserService(userRepository, loggerService)
	jwtService, _ := infrastructureService.NewJWTService("YELLOW SUBMARINE, BLACK WIZARDRY", loggerService)

	services := &controllers.Services{
		CryptoService: cryptoService,
		UserService:   userService,
		LoggerService: loggerService,
	}

	applicationController := controllers.NewApplicationController(appConfig)
	helloController := controllers.NewHelloController(*sayHelloUseCase, loggerService)
	userController := controllers.NewUserController(services)
	authController := controllers.NewAuthController(cryptoService, jwtService, userService, loggerService)

	resolvers := graphql.NewResolverRegistry(
		cryptoService,
		userService,
		loggerService,
		appConfig,
	)

	schema := graphql.NewSchema(resolvers)
	graphQLController := controllers.NewGraphQLController(schema)

	router := chi.NewRouter()
	router.Use(customMiddleware.CORSMiddleware())
	//router.Use(customMiddleware.AuthorizeJWT())
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/ping"))

	if profilingConfig.Enabled {
		router.Mount("/debug", middleware.Profiler())
	}

	router.Mount("/swagger", routes.NewSwaggerRoutes().Mount())
	router.Mount("/graphql", routes.NewGraphQLRoutes(graphQLController).Mount())
	router.Mount("/application", routes.NewApplicationRoutes(applicationController).Mount())

	router.Route("/api", func(r chi.Router) {
		r.Mount("/v1/hello", routes.NewHelloRoutes(helloController).Mount())
		r.Mount("/v1/auth", routes.NewAuthRoutes(authController).Mount())
		r.Mount("/v1/users", routes.NewUserRoutes(userController).Mount())
	})

	return &AppDependencies{
		Router: router,
	}
}
