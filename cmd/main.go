package main

import (
	"io"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("/logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Ошибка открытия лог-файла: %v", err)
	}

	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	log.Println("Приложение запущено")
}
