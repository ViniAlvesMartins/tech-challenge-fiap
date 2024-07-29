package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseHost     string `envconfig:"database_host"`
	DatabasePort     string `envconfig:"database_port"`
	DatabaseDBName   string `envconfig:"database_name"`
	DatabaseSchema   string `envconfig:"database_schema"`
	DatabaseUsername string `envconfig:"database_username"`
	DatabasePassword string `envconfig:"database_password"`
	MigrationsDir    string `envconfig:"migrations_dir"`

	SnsRegion               string `envconfig:"sns_region"`
	SnsUrl                  string `envconfig:"sns_url"`
	SnsAccessKey            string `envconfig:"sns_access_key"`
	OrderCreatedTopic       string `envconfig:"order_created_topic"`
	OrderStatusQueue        string `envconfig:"order_status_queue"`
	OrderPaymentStatusQueue string `envconfig:"order_payment_status_queue"`
}

func NewConfig() (cfg Config, err error) {
	err = envconfig.Process("", &cfg)
	return
}
