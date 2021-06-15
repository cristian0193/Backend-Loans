package persistence

import (
	"Backend-Loans/domain/entity"
	"Backend-Loans/infrastructure/persistence"
	"Backend-Loans/infrastructure/repository"
	"Backend-Loans/utils"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DbHelper struct {
	LoansRepository     repository.LoansRepository
	PaymentsRepository  repository.PaymentsRepository
	InterestsRepository repository.InterestsRepository
	UsersRepository     repository.UsersRepository
	ClientsRepository   repository.ClientsRepository
	TypesRepository     repository.TypesRepository
	db                  *gorm.DB
}

func InitDbHelper() (*DbHelper, error) {

	var host = utils.GetEnv("DB_HOST", "localhost")
	var port = utils.GetEnv("DB_PORT", "5432")
	var user = utils.GetEnv("DB_USER", "postgres")
	var password = utils.GetEnv("DB_PASSWORD", "mysecretpassword")
	var dbname = utils.GetEnv("DB_NAME", "postgres")
	var drive = utils.GetEnv("DB_DRIVER", "postgres")

	// Heroku
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)

	// Local
	/* psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname) */
	db, err := gorm.Open(drive, psqlInfo)

	if err != nil {
		panic(err)
	}
	sqlDB := db.DB()
	sqlDB.SetMaxIdleConns(600)
	sqlDB.SetMaxOpenConns(0)
	db.LogMode(true)
	db.AutoMigrate()
	return &DbHelper{
		LoansRepository:     persistence.InitLoansRepositoryImpl(db),
		PaymentsRepository:  persistence.InitPaymentsRepositoryImpl(db),
		InterestsRepository: persistence.InitInterestsRepositoryImpl(db),
		UsersRepository:     persistence.InitUsersRepositoryImpl(db),
		ClientsRepository:   persistence.InitClientsRepositoryImpl(db),
		TypesRepository:     persistence.InitTypesRepositoryImpl(db),
		db:                  db,
	}, nil
}

func (s *DbHelper) Close() error {
	return s.db.Close()
}

func (s *DbHelper) Automigrate() error {
	return s.db.AutoMigrate(
		&entity.Loans{},
		&entity.Payments{},
		&entity.Interests{},
		&entity.Types{},
		&entity.Users{},
		&entity.Clients{},
		&entity.Types{}).Error
}
