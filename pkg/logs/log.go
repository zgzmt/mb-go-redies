package logs

import (
	"fmt"
	"log"
	configs "mb-go-redis/configs"
	"os"
)
var G_Logger *log.Logger
func init()  {
	logFile, err := os.Open(configs.G_LogConfig.Path)
	if err != nil{
		log.Logger.Fatal(fmt.Sprintf("listen err: %v", err))
	}
	G_Logger = log.New(logFile, "redist", 1)
}
