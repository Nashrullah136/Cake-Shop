package app

import (
	"Pretests/controller/cake"
	"Pretests/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rubenv/sql-migrate"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func Run() {
	router := Initialize()
	router.Run(":" + os.Getenv("APP_PORT"))
}

func loadEnv() {
	if err := godotenv.Load(basepath + "/../.env"); err != nil {
		log.Fatalln("Fail to load .env")
	}
}

func migration() {
	isMigrate, err := strconv.ParseBool(os.Getenv("MIGRATE"))
	if err != nil {
		log.Fatalln("Wrong type for MIGRATE in .env")
	}
	if isMigrate {
		migrations := &migrate.FileMigrationSource{
			Dir: basepath + "/../database/migrations",
		}
		config := database.DefaultConfig()
		config.ParseTime = true
		dbConnection := database.SetUpDbConnection(config)
		numOfMigrations, err := migrate.Exec(dbConnection, "mysql", migrations, migrate.Up)
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Printf("Successfully migrate %d migrations", numOfMigrations)
	}
}

func Initialize() *gin.Engine {
	log.Printf(basepath)
	loadEnv()
	database.GetDbConnection()
	migration()
	router := setUpRouter()
	return router
}

func setUpRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/cakes", cake.GetAll)
	router.GET("/cakes/:id", cake.GetById)
	router.POST("/cakes", cake.Create)
	router.PATCH("/cakes/:id", cake.Update)
	router.DELETE("/cakes/:id", cake.Delete)
	return router
}
