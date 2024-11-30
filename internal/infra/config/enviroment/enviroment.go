package enviroment

type Config struct {
	MongoURL    string `var:"MONGO_URL"`
	MongoDBName string `var:"MONGO_DB_NAME"`
}