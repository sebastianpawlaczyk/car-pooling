package cfg

type Config struct {
	Port         int    `required:"false" env:"PORT" envDefault:"9091"`
	MongoURI     string `required:"true" env:"MONGO_URI"`
	DatabaseName string `required:"true" env:"DATABASE_NAME"`
}
