package services

import (
	"belajargo/models"
	"belajargo/repos"
)

type PrizeServiceInterface interface {
	SemuaDataService() ([]models.Prize, error)
	CariByIdService(ID int) (models.Prize, error)
	InsertDataService(input models.PrizesInput) (models.Prize, error)
	UpdateDataService(ID int, input models.PrizesInput) (models.Prize, error)
	HapusDataService(ID int) (models.Prize, error)
}

type initPrizeService struct {
	srvc repos.PrizesInterface
}

func PrizeService(srvcRepo repos.PrizesInterface) *initPrizeService {
	return &initPrizeService{srvcRepo}
}

// SERVICE
// SemuaDataService() ([]models.Prize, error)
// CariByIdService(ID int) (models.Prize, error)
// InsertDataService(b models.PrizesInput) (models.Prize, error)

func (sv *initPrizeService) SemuaDataService() ([]models.Prize, error) {
	prize, err := sv.srvc.SemuaData()
	return prize, err
}

func (sv *initPrizeService) CariByIdService(ID int) (models.Prize, error) {
	prize, err := sv.srvc.CariById(ID)
	return prize, err
}

func (sv *initPrizeService) InsertDataService(input models.PrizesInput) (models.Prize, error) {
	indexPrizeParse, _ := input.IndexPrize.Int64()
	qoutaParse, _ := input.Qouta.Int64()
	dayLimitParse, _ := input.DayLimit.Int64()

	prizeData := models.Prize{
		IndexPrize: int(indexPrizeParse),
		PrizeName:  input.PrizeName,
		Qouta:      int(qoutaParse),
		DayLimit:   int(dayLimitParse),
	}
	prize, err := sv.srvc.InsertData(prizeData)
	return prize, err
}

func (sv *initPrizeService) UpdateDataService(ID int, input models.PrizesInput) (models.Prize, error) {
	indexPrizeParse, _ := input.IndexPrize.Int64()
	qoutaParse, _ := input.Qouta.Int64()
	dayLimitParse, _ := input.DayLimit.Int64()

	cariData, _ := sv.srvc.CariById(ID)
	cariData.IndexPrize = int(indexPrizeParse)
	cariData.PrizeName = input.PrizeName
	cariData.Qouta = int(qoutaParse)
	cariData.DayLimit = int(dayLimitParse)

	prize, err := sv.srvc.UbahData(cariData)
	return prize, err
}

func (sv *initPrizeService) HapusDataService(ID int) (models.Prize, error) {
	cariData, _ := sv.srvc.CariById(ID)
	prize, err := sv.srvc.HapusData(cariData)
	return prize, err
}
