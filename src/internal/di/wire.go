//go:build wireinject
// +build wireinject

package di

import (
	articleApp "go_2601_04/internal/application/article"
	userApp "go_2601_04/internal/application/user"
	articleDomain "go_2601_04/internal/domain/article"
	userDomain "go_2601_04/internal/domain/user"
	mysqlRepo "go_2601_04/internal/infrastructure/persistence/mysql"
	"go_2601_04/internal/interfaces/http"
	handler "go_2601_04/internal/interfaces/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// repository
var repositorySet = wire.NewSet(
	mysqlRepo.NewUserRepository,
	wire.Bind(new(userDomain.Repository), new(*mysqlRepo.UserRepository)),

	mysqlRepo.NewArticleRepository,
	wire.Bind(new(articleDomain.Repository), new(*mysqlRepo.ArticleRepository)),
)

// service
var serviceSet = wire.NewSet(
	userApp.NewUserService,
	articleApp.NewArticleService,
)

// handler
var handlerSet = wire.NewSet(
	handler.NewUserHandler,
	handler.NewArticleHandler,
)

func InitializeApp(dsn string) (*gin.Engine, error) {
	wire.Build(
		mysqlRepo.NewDatabase,
		repositorySet,
		serviceSet,
		handlerSet,
		setupRouter,
	)
	return nil, nil
}

func setupRouter(u *http.UserHandler, a *http.ArticleHandler) *gin.Engine {
	r := gin.Default()
	u.Register(r)
	a.Register(r)
	return r
}
