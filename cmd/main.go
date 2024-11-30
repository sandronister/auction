package main

import (
	"github.com/sandronister/auction-go/internal/infra/config/enviroment"
	"github.com/sandronister/auction-go/internal/infra/database/mongodb"
	"github.com/sandronister/auction-go/internal/infra/logger"
	"github.com/sandronister/enviroment-go/pkg/load"
)

var config enviroment.Config;

func init(){
	varEnv := load.New("../.env")
	errs := varEnv.Load(config)

	if len(errs) > 0 {
		for _, err := range errs {
			logger.Error("Error loading environment variables", err)
		}
		return
	}
}


func main(){
	database,err:=mongodb.NewConnection(config.MongoURL,config.MongoDBName)	

	if err != nil {
		logger.Error("Error connecting to database",err)
		return
	}
	
}