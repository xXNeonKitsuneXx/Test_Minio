package main

import (
	"fmt"
	"log"
	"strings"
	"test_Minio/internal/handler"
	"test_Minio/internal/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/spf13/viper"
)

func main() {
	initTimeZone()
	initConfig()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)
	log.Println(dsn)

	// Initialize minio client object.
	minioClient, err := minio.New(viper.GetString("minio.host")+":"+viper.GetString("minio.port"), &minio.Options{
		Creds:  credentials.NewStaticV4("EneudJTZpjRIZuFYq6MF", "eo1uLxOg33oeuOHbI7kokmSNSp1AgJTfw1QMWLda", ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Minio connected")

	uploadSvc := service.NewUploadService(minioClient)
	storageHandler := handler.NewStorageHandler(uploadSvc)

	app := fiber.New()

	//Endpoint ###########################################################################
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/upload", storageHandler.UploadFile)
	//#####################################################################################

	log.Printf("TestMinio run at port:  %v", viper.GetInt("app.port"))
	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}
