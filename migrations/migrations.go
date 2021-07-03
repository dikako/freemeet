package migrations

import (
	"github.com/dikako/free-meet/helpers"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

//db connection
func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=freemeet password=freemeet sslmode=disable")
	helpers.HandleErr(err)

	return db
}

func createAccounts() {
	db := connectDB()

	users := [2]User{
		{Username: "Dika", Email: "dika@mailnesia.com"},
		{Username: "Dika Koko", Email: "dikakoko@mailnesia.com"},
	}

	for i := 0; i < len(users); i++ {
		generatePassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := User{
			Username: users[i].Username,
			Email:    users[i].Email,
			Password: generatePassword,
		}
		db.Create(&user)

		account := Account{
			Type:    "Daily Account",
			Name:    string(users[i].Username + "'s" + " account"),
			Balance: uint(10000 + int(i+1)),
			UserID:  user.ID,
		}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate() {
	db := connectDB()
	db.AutoMigrate(&User{}, &Account{})
	defer db.Close()

	createAccounts()
}
