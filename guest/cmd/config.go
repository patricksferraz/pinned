package cmd

type Config struct {
	RestPort *int `env:"REST_PORT,default=8080"`

	Db struct {
		Debug   *bool   `env:"DB_DEBUG,default=false"`
		Migrate *bool   `env:"DB_MIGRATE,default=false"`
		DsnType *string `env:"DSN_TYPE,required=true"`
		Dsn     *string `env:"DSN,required=true"`
	}

	Kafka struct {
		Servers *string `env:"KAFKA_BOOTSTRAP_SERVERS,default=kafka:9092"`
		GroupId *string `env:"KAFKA_CONSUMER_GROUP_ID,default=guest-check"`
	}
}
