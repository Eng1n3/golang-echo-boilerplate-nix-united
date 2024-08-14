package db

import (
	"echo-demo-project/config"
	"echo-demo-project/db/seeders"
	"fmt"

	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func Init(cfg *config.Config) *gorm.DB {
	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	cfg.DB.User,
	// 	cfg.DB.Password,
	// 	cfg.DB.Host,
	// 	cfg.DB.Port,
	// 	cfg.DB.Name)

	dataSourceNameSqlite := fmt.Sprintf("%s.sqlite", cfg.DBSqlite.Name)

	fmt.Println(dataSourceNameSqlite)

	db, err := gorm.Open(sqlite.Open(dataSourceNameSqlite), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	//Build seeders stack
	usersSeeder := seeders.NewUsersSeeder(gorm_seeder.SeederConfiguration{Rows: 10})
	seedersStack := gorm_seeder.NewSeedersStack(db)
	seedersStack.AddSeeder(&usersSeeder)

	//Apply seed
	seedersStack.Seed()
	// fmt.Println(39)
	//Apply clear
	// err = seedersStack.Clear()
	// fmt.Println(err)
	// userSeeder := seeders.NewUserSeeder(db)
	// userSeeder.SetUsers()

	return db
}
