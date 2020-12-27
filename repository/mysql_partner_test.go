package repository

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/arfan21/getprint-partner/models"
	"github.com/arfan21/getprint-partner/utils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
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

func ClearTable(db *gorm.DB) {
	db.Migrator().DropTable(&models.Address{}, &models.Price{}, &models.Partner{})

	db.AutoMigrate(&models.Partner{}, &models.Price{}, &models.Address{})
}

func TestCreate(t *testing.T) {
	db, err := Conn()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// ClearTable(db)

	repo := NewPartnerRepo(db)

	price := new(models.Price)
	price.Print.Scan(1000)
	price.Fotocopy.Scan(1000)
	price.Scan.Scan(1000)

	address := new(models.Address)
	address.Address = "karanganyar, jawa tengah"
	address.Lat = "1.2131231"
	address.Lng = "-12,1200012"

	payload := &models.Partner{
		UserID:      1,
		Name:        "enter komp",
		Email:       "enter@enter.com",
		PhoneNumber: "6281222212",
		Price:       *price,
		Address:     *address,
	}

	err = repo.Create(payload)

	assert.NoError(t, err, "should be no error")
	assert.NotZero(t, payload.ID, "should be no zero")
	assert.Equal(t, *&payload.ID, *&payload.Price.PartnerID, "should be equal")
	assert.Equal(t, *&payload.ID, *&payload.Address.PartnerID, "should be equal")
}

func TestGet(t *testing.T) {
	db, err := Conn()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repo := NewPartnerRepo(db)

	data, err := repo.Fetch("status=?", "active")

	d, _ := json.MarshalIndent(data, "", "\t")

	log.Println(string(d))

	assert.NoError(t, err, "should be no error")
	assert.NotNil(t, data, "should be not nil")
}

func TestGetById(t *testing.T) {
	db, err := Conn()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repo := NewPartnerRepo(db)

	data, err := repo.GetByID(2)

	if err != nil {
		if err.Error() == "partner not found" {
			assert.Error(t, err)
			return
		}
	}

	d, _ := json.MarshalIndent(data, "", "\t")

	log.Println(string(d))

	assert.NoError(t, err, "should be no error")
	assert.NotNil(t, data, "should be not nil")
}

func TestUpdate(t *testing.T) {
	db, err := Conn()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	price := new(models.Price)
	price.Print.Scan(1250)
	price.Fotocopy.Scan(2000)
	price.Scan.Scan(500)

	address := new(models.Address)
	address.Address = "malang, jawa timur"
	address.Lat = "1.2131"
	address.Lng = "-12,2221"

	payload := &models.Partner{
		Name:        "kios komp",
		Email:       "kios@enter.com",
		PhoneNumber: "628",
		Price:       *price,
		Address:     *address,
		Status:      "active",
	}

	repo := NewPartnerRepo(db)

	err = repo.Update(1, payload)

	assert.NoError(t, err, "should be no error")
}
