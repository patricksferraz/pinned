/*
Copyright © 2022 Patrick Ferraz <patrick.ferraz@outlook.com>

*/
package cmd

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Netflix/go-env"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	appKafka "github.com/patricksferraz/guest-check/app/kafka"
	"github.com/patricksferraz/guest-check/app/rest"
	"github.com/patricksferraz/guest-check/infra/client/kafka"
	"github.com/patricksferraz/guest-check/infra/client/kafka/topic"
	"github.com/patricksferraz/guest-check/infra/db"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
func allCmd() *cobra.Command {
	var conf Config

	_, err := env.UnmarshalFromEnviron(&conf)
	if err != nil {
		log.Fatal(err)
	}

	allCmd := &cobra.Command{
		Use:   "all",
		Short: "Run both gRPC and rest servers",

		Run: func(cmd *cobra.Command, args []string) {
			pg, err := db.NewPostgreSQL(*conf.Db.DsnType, *conf.Db.Dsn)
			if err != nil {
				log.Fatal(err)
			}

			if *conf.Db.Debug {
				pg.Debug(true)
			}

			if *conf.Db.Migrate {
				pg.Migrate()
			}
			defer pg.Db.Close()

			kc, err := kafka.NewKafkaConsumer(*conf.Kafka.Servers, *conf.Kafka.GroupId, topic.CONSUMER_TOPICS)
			if err != nil {
				log.Fatal("cannot start kafka consumer", err)
			}

			deliveryChan := make(chan ckafka.Event)
			kp, err := kafka.NewKafkaProducer(*conf.Kafka.Servers, deliveryChan)
			if err != nil {
				log.Fatal("cannot start kafka producer", err)
			}

			go kp.DeliveryReport()
			go appKafka.StartKafkaServer(pg, kc, kp)
			rest.StartRestServer(pg, kp, *conf.RestPort)
		},
	}

	return allCmd
}

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load(basepath + "/../.env")
		if err != nil {
			log.Printf("Error loading .env files")
		}
	}

	rootCmd.AddCommand(allCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
