package main

import (
	"github.com/sandronister/auction-go/internal/infra/config/enviroment"
	"github.com/sandronister/auction-go/internal/infra/database/mongodb"
	"github.com/sandronister/auction-go/internal/infra/logger"
	"github.com/sandronister/enviroment-go/pkg/load"
)

var config =enviroment.Config{};

func init(){
	logger.Init()
	varEnv,err := load.New("../.env")

	if err != nil {
		logger.Error("Error loading environment variables",err)
		return
	}

	
	err = varEnv.LoadVariable(&config)

	if err != nil {
		logger.Error("Error loading environment variables",err)
		return
	}
}


func main(){
	_,err:=mongodb.NewConnection(config.MongoURL,config.MongoDBName)	

	if err != nil {
		logger.Error("Error connecting to database",err)
		return
	}
	
	logger.Info("Connected to database")
}