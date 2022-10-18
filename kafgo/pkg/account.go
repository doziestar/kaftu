package pkg

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
  )

type IAccount interface {
	GetAccount() (Account, error)
	CreateAccount() (Account, error)
	Deposit(amount float64) (Account, error)
	Withdraw(amount float64) (Account, error)
	GetBalance() (float64, error)
}

type Account struct {
	gorm.Model
	Username string `gorm:"unique index"`
	UserID   string `gorm:"unique index"`
	BookBalance  float64 `gorm:"default:0"`
	BookBalanceCurrency string `gorm:"default:USD"`
	AvailableBalance float64 `gorm:"default:0"`
	AvailableBalanceCurrency string `gorm:"default:USD"`
}

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository() *AccountRepository {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"))
	CheckError(err)
	db.AutoMigrate(&Account{})
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) GetAccount() (Account, error) {
	var account Account
	result := r.db.First(&account)
	return account, result.Error
}

func (r *AccountRepository) CreateAccount() (Account, error) {
	var account Account
	result := r.db.Create(&account)
	return account, result.Error
}

func (r *AccountRepository) Deposit(amount float64) (Account, error) {
	var account Account
	result := r.db.Model(&account).Update("book_balance", amount)
	return account, result.Error
}

func (r *AccountRepository) Withdraw(amount float64) (Account, error) {
	var account Account
	result := r.db.Model(&account).Update("book_balance", amount)
	return account, result.Error
}

func (r *AccountRepository) GetBalance() (float64, error) {
	var account Account
	result := r.db.First(&account)
	return account.BookBalance, result.Error
}

func (r *AccountRepository) GetAvailableBalance() (float64, error) {
	var account Account
	result := r.db.First(&account)
	return account.AvailableBalance, result.Error
}

func (r *AccountRepository) GetBookBalance() (float64, error) {
	var account Account
	result := r.db.First(&account)
	return account.BookBalance, result.Error
}