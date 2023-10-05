package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewConnection(ctx context.Context, log *slog.Logger, cfg Config) (*gorm.DB, error) {
	var err error
	var conn *gorm.DB

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfg.DatabaseHost, cfg.DatabaseUsername, cfg.DatabasePassword, cfg.DatabaseDBName, cfg.DatabasePort)

	conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormlogger.Discard,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s", cfg.DatabaseSchema),
			SingularTable: false,
		},
	})

	if err != nil {
		log.Error(fmt.Sprintf("Error to connect to schema %s", cfg.DatabaseSchema))
		return nil, err
	}

	db, err := conn.DB()

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Info(fmt.Sprintf("Successfuly connected to %s database", cfg.DatabaseDBName))

	return conn, nil

}
