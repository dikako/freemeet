package migrations

import (
	"github.com/dikako/free-meet/helpers"
	"github.com/dikako/free-meet/interfaces"
)

// type User struct {
// 	gorm.Model
// 	Username string
// 	Email    string
// 	Password string
// }

// type Account struct {
// 	gorm.Model
// 	Type    string
// 	Name    string
// 	Balance uint
// 	UserID  uint
// }

//db
// func connectDB() *gorm.DB {
// 	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=freemeet password=freemeet sslmode=disable")
// 	helpers.HandleErr(err)

// 	return db
// }

func createAccounts() {
	db := helpers.ConnectDB()

	users := &[2]interfaces.User{
		{Username: "Dika", Email: "dika@mailnesia.com"},
		{Username: "Dika Koko", Email: "dikakoko@mailnesia.coms"},
	}

	for i := 0; i < len(users); i++ {
		generatePassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{
			Username: users[i].Username,
			Email:    users[i].Email,
			Password: generatePassword,
		}
		db.Create(&user)

		account := &interfaces.Account{
			Type:    "Daily Account",
			Name:    string(users[i].Username + "'s" + " account"),
			Balance: uint(10000 + int(i+1)),
			UserId:  user.ID,
		}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate() {
	User := &interfaces.User{}
	Account := interfaces.Account{}
	db := helpers.ConnectDB()
	db.AutoMigrate(&User, &Account)
	defer db.Close()

	createAccounts()
}
