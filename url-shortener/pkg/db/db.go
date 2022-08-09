package db

import (
    "log"
    "strconv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
    "github.com/BabouZ17/url-shortener/pkg/model"
    "github.com/BabouZ17/url-shortener/pkg/config"
)


func New(config config.Config) *gorm.DB {
    host := config.PostgresConfig.Host
    port := strconv.Itoa(config.PostgresConfig.Port)
    username := config.PostgresConfig.Username
    password := config.GetEnvValue(config.PostgresConfig.Password).(string)
    database := config.PostgresConfig.Db
    databaseUrl := "postgres://" + username + ":" + password + "@" + host + ":" + port + "/" + database

    db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&model.Url{})

    return db
}
