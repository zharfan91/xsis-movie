package postgres

import (
	"fmt"
	"log"
	"xsis-movie/config"
	"xsis-movie/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

var (
	psqlConn *gorm.DB
	err      error
)

// // initialize database
//
//	func init() {
//		setupPostsql()
//	}
func setupPostsql(cfg *config.Config) *gorm.DB {
	// cfg := &config.Config{}
	// if err != nil {
	// 	log.Fatalf("Error getting env, %v", err)
	// }
	dns := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s TimeZone=%s",
		cfg.PG.DB_HOST, cfg.PG.DB_PORT, cfg.PG.DB_USER, cfg.PG.DB_NAME, cfg.PG.DB_PASS, cfg.PG.TIME_ZONE)
	log.Println("dns : ", dns)
	psqlConn, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to create a connection to database Registry error:%s", err)
	}
	psqlConn.AutoMigrate(entity.Movies{})
	return psqlConn
}

func Postsql(cfg *config.Config) *gorm.DB {
	return setupPostsql(cfg)
}
