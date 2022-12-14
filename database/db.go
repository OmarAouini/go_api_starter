package database

import (
	"fmt"

	"github.com/OmarAouini/golang_api_starter/config"
	gorm_logrus "github.com/onrik/gorm-logrus"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// global db connection
var DB *gorm.DB

// connect to database and set global DB connection
//
// ssl mode eg: disable
//
// timezone eg: Asia/Shanghai
func ConnectDb() {

	// for additional configuration and docs, refer to: https://gorm.io/docs/
	logrus.Infof("connecting to db %s on schema %s...\n", config.AppConfiguration.DbName, config.AppConfiguration.DbSchema)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", config.AppConfiguration.DbHost, config.AppConfiguration.DbUsername,
		config.AppConfiguration.DbPassword, config.AppConfiguration.DbName, config.AppConfiguration.DbPort, config.AppConfiguration.DbSslMode, config.AppConfiguration.DbTimezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gorm_logrus.New(),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s.", config.AppConfiguration.DbSchema), // schema name here
			SingularTable: false,
		}})

	if err != nil {
		logrus.Fatal(err)
	}

	dbConf, _ := db.DB()

	dbConf.SetMaxIdleConns(config.AppConfiguration.DbMinConn)
	dbConf.SetMaxOpenConns(config.AppConfiguration.DbMaxConn)

	DB = db

	logrus.Info("connected to database!")
}
