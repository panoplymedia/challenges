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
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
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
			time.Sleep(1 * time.Second)
		} else {
			break
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

func (e *DbEngine) VerifyUser(email, enteredPass string) (*models.User, bool) {
	usr := &models.User{}
	err := e.DB.Model(models.User{}).
		Where("email = $1", email).
		Find(usr).Error

	if err != nil {
		e.Logger.Errorln(err)
		return nil, false
	}

	if usr.Email == "" {
		e.Logger.Errorln(fmt.Errorf("no user found with email %s", email))
	}

	row := e.DB.Model(models.User{}).
		Select("password").
		Where("email = $1", email).
		Where("password = crypt($2, $3)", enteredPass, usr.Password).Row()

	var hash string
	err = row.Scan(&hash)
	if err != nil {
		e.Logger.Println(err)
		return nil, false
	}

	return usr, hash == usr.Password
}
