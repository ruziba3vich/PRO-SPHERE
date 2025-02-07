package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ruziba3vich/prosphere/internal/items/config"
	"github.com/ruziba3vich/prosphere/internal/items/http/app"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	cfg, err := config.New()
	if err != nil {
		logger.Fatalln(err)
	}

	go app.Run(cfg, logger)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	logger.Println("Server stopped gracefully")
}
