package main

import (
	"HR-monitor/pkg/config"
	"HR-monitor/pkg/repository"
	"context"
	"log"
)


func main() {
	ctx := context.Background()
	config := config.LoadConfig()
	err := repository.InitDB(ctx, config)
	if err != nil {
		log.Fatal("DB init failed ", err)

	}
	defer repository.CloseDB()
	
	
}


