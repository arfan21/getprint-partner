package follower

import (
	"log"
	"os"
	"path/filepath"

	"github.com/arfan21/getprint-partner/utils"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func loadEnv() {
	rootPath, err := os.Getwd()

	err = godotenv.Load(os.ExpandEnv(filepath.Dir(rootPath) + "/.env"))

	if err != nil {
		log.Fatalf("can't load env file : %v", err)
	}
}

func Conn() (*gorm.DB, error) {
	loadEnv()

	return utils.Connect()
}

// func BenchmarkCreate(b *testing.B) {

// 	db, err := Conn()

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	repo := repository.NewFollowerRepo(db)
// 	repoPartner := repository.NewPartnerRepo(db)

// 	fS := NewFollowerService(repo, repoPartner)

// 	for i := 0; i < b.N; i++ {
// 		follower := new(models.Follower)
// 		follower.PartnerID = 1
// 		follower.UserID = 1
// 		fS.Create(follower)
// 	}
// }
