package main

import (
	"io"
	"log"
	"os"

	"github.com/Strike-official/global-getting-started/configmanager"
	"github.com/gin-gonic/gin"
)

func main() {
	// Read Config
	err := configmanager.InitAppConfig("configs/config.json")
	if err != nil {
		log.Fatal("[Initialize] Failed to start APIs. Error: ", err)
	}
	conf := configmanager.GetAppConfig()

	// Init LogFile
	logFile := initLogger(conf.LogFilePath)

	// Init Routes
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	router := gin.Default()
	routes(router)

	// Start serving the application
	err = router.Run(conf.Port)
	if err != nil {
		log.Fatal("[Initialize] Failed to start server. Error: ", err)
	}
}

func initLogger(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	return file
}
