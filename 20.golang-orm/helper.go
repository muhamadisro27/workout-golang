package golangorm

import "gorm.io/gorm"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func BrokeWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance = ?", 0)
}

func SultanWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance > ?", 1000000)
}
