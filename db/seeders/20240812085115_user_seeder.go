package seeders

import (
	"echo-demo-project/models"
	"time"

	gorm_seeder "github.com/kachit/gorm-seeder"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UsersSeeder struct {
	gorm_seeder.SeederAbstract
}

func NewUsersSeeder(cfg gorm_seeder.SeederConfiguration) UsersSeeder {
	return UsersSeeder{gorm_seeder.NewSeederAbstract(cfg)}
}

// Implement Seed method
func (s *UsersSeeder) Seed(db *gorm.DB) error {
	versionID := 20240812085115
	seederResult := db.Find(&models.Seeder{VersionID: versionID})
	if seederResult.RowsAffected != 0 {
		return nil
	}
	users := []models.User{
		{
			Email:    "user1@example.com",
			Name:     "user1",
			Password: "password1",
		},
		{
			Email:    "user2@example.com",
			Name:     "user2",
			Password: "password2",
		},
	}

	for i, value := range users {
		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(value.Password), bcrypt.DefaultCost)

		if err != nil {
			panic(err)
		}
		users[i].Password = string(encryptedPassword)
	}

	db.Create(users)

	seed := models.Seeder{
		VersionID: versionID,
		IsApplied: true,
		Tstamp:    time.Now().UTC(),
	}

	return db.CreateInBatches(seed, s.Configuration.Rows).Error
}

// Implement Clear method
func (s *UsersSeeder) Clear(db *gorm.DB) error {
	entity := models.User{}
	return s.SeederAbstract.Delete(db, entity.TableName())
}
