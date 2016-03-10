package utils


import (
	"log"
	"os"
	"fmt"
)

var logger *log.Logger

func InitLog() {

	f, _ := os.OpenFile("nazar.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	defer f.Close()

	logger = log.New(f, "", 0)
}

func LogError(v ...interface{}) {
	logger.Println("Error: " + fmt.Sprint(v))
}

func LogInfo(v ...interface{}) {
	logger.Println(fmt.Sprint(v))
}