package app

import (
	"log"
	"xsis-movie/config"
	v1 "xsis-movie/internal/controller/http/v1"
	"xsis-movie/internal/usecase"
	"xsis-movie/internal/usecase/repo"
	"xsis-movie/pkg/postgres"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {

	// Repository
	// pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	// }
	// defer pg.Close()

	postsql := postgres.Postsql(cfg)
	// Use case
	moviesUseCase := usecase.New(
		repo.New(postsql),
	)

	// HTTP Server
	handler := gin.Default()
	v1.NewRouter(handler, moviesUseCase)
	// server
	if err := handler.Run(":" + cfg.HTTP.Port); err != nil {
		log.Println(err)
		return
	}

}
