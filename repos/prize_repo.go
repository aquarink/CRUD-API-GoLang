package repos

import (
	"belajargo/models"

	"gorm.io/gorm"
)

type PrizesInterface interface {
	SemuaData() ([]models.Prize, error)
	CariById(ID int) (models.Prize, error)
	InsertData(b models.Prize) (models.Prize, error)
	UbahData(b models.Prize) (models.Prize, error)
	HapusData(b models.Prize) (models.Prize, error)
}

type prizesDB struct {
	database *gorm.DB
}

func PrizeRepository(db *gorm.DB) *prizesDB {
	return &prizesDB{db}
}

// REPO

func (repo *prizesDB) SemuaData() ([]models.Prize, error) {
	var prizes []models.Prize

	err := repo.database.Find(&prizes).Error

	return prizes, err
}

func (repo *prizesDB) CariById(ID int) (models.Prize, error) {
	var book models.Prize

	// err := repo.database.Where("id = ?", ID).Find(&book).Error
	err := repo.database.Find(&book, ID).Error

	return book, err
}

func (repo *prizesDB) UbahData(bk models.Prize) (models.Prize, error) {
	err := repo.database.Save(&bk).Error

	return bk, err
}

func (repo *prizesDB) InsertData(bk models.Prize) (models.Prize, error) {
	err := repo.database.Create(&bk).Error

	return bk, err
}

func (repo *prizesDB) HapusData(bk models.Prize) (models.Prize, error) {
	err := repo.database.Delete(&bk).Error

	return bk, err
}
