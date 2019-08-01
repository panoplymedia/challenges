package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"sales-portal-api/environment"
	"sales-portal-api/models"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type DbEngine struct {
	Logger *logrus.Logger
	DB     *gorm.DB
}

func NewEngine(env *environment.Environment, logger *logrus.Logger) (*DbEngine, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		env.DbUser,
		env.DbPass,
		env.DbHost,
		env.DbPort,
		env.DbName,
	)

	dbMaxLife, _ := strconv.Atoi(env.DbMaxLifetime)
	dbMaxOpen, _ := strconv.Atoi(env.DbMaxOpenConnections)
	dbMaxIdle, _ := strconv.Atoi(env.DbMaxIdleConnections)
	dbMaxAttempts, _ := strconv.Atoi(env.DbMaxAttempts)

	var db *gorm.DB
	var err error
	for i := 0; i < dbMaxAttempts; i++ {
		db, err = gorm.Open("postgres", dbUrl)
		if err != nil {
			logger.Infoln(err)
		}
	}

	db.SingularTable(true)
	db.DB().SetConnMaxLifetime(time.Duration(dbMaxLife) * time.Second)
	db.DB().SetMaxOpenConns(dbMaxOpen)
	db.DB().SetMaxIdleConns(dbMaxIdle)

	db.SetLogger(logger)

	engine := &DbEngine{
		Logger: logger,
		DB:     db,
	}

	return engine, nil
}

func (e *DbEngine) VerifyUser(username, password string) bool {
	usr := &models.User{}
	err := e.DB.Model(models.User{}).
		Where("username = $1", username).
		Find(usr).Error

	if err != nil {
		e.Logger.Debugln(err)
		return false
	}

	row := e.DB.Model(models.User{}).
		Select("password_hash").
		Where("username = $1", username).
		Where("password_hash=(crypt($2, $3))", password, usr.PasswordHash).
		Row()

	var hash string
	err = row.Scan(&hash)
	if err != nil {
		e.Logger.Debugln(err)
		return false
	}

	return hash == usr.PasswordHash
}

func(e *DbEngine) LoadCsvRecords(path string) error {

}