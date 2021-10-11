package controls

import (
	"belajargo/models"
	"belajargo/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type prizeHandler struct {
	prizeSrvc services.PrizeServiceInterface
}

func PrizeController(przService services.PrizeServiceInterface) *prizeHandler {
	return &prizeHandler{przService}
}

// FUNGSI

func (handler *prizeHandler) PrizeIndex(c *gin.Context) {
	prz, err := handler.prizeSrvc.SemuaDataService()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"datas":   err,
			"message": "Load data gagal",
		})
		return
	}

	var przJson []models.PrizeToJson

	for _, p := range prz {
		jsonnya := convertPrizeResponse(p)

		przJson = append(przJson, jsonnya)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"datas":   przJson,
		"message": "Load data berhasil",
	})
}

func (handler *prizeHandler) PostPrize(c *gin.Context) {

	var bookData models.PrizesInput
	err := c.ShouldBindJSON(&bookData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"datas":   err,
			"message": "Post data gagal",
		})
		return
	}

	prz, err := handler.prizeSrvc.InsertDataService(bookData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"datas":   err,
			"message": "Save data gagal",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"datas":   prz,
		"message": "Post data berhasil",
	})
}

func (handler *prizeHandler) PrizeById(c *gin.Context) {

	// idnya := c.Query("id")
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)

	prz, err := handler.prizeSrvc.CariByIdService(idInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"datas":   err,
			"message": "Load data gagal",
		})
		return
	}

	jsonnya := convertPrizeResponse(prz)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"datas":   jsonnya,
		"message": "Load data berhasil",
	})
}

func (handler *prizeHandler) UpdatePrize(c *gin.Context) {

	// idnya := c.Query("id")
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)

	var bookData models.PrizesInput
	err := c.ShouldBindJSON(&bookData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"datas":   err,
			"message": "Post data gagal",
		})
		return
	}

	prz, err := handler.prizeSrvc.UpdateDataService(idInt, bookData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"datas":   err,
			"message": "Ubah data gagal",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"datas":   prz,
		"message": "Ubah data berhasil",
	})
}

func (handler *prizeHandler) DeleteById(c *gin.Context) {

	// idnya := c.Query("id")
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)

	prz, err := handler.prizeSrvc.HapusDataService(idInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"datas":   err,
			"message": "Hapus data gagal",
		})
		return
	}

	jsonnya := convertPrizeResponse(prz)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"datas":   jsonnya,
		"message": "Hapus data berhasil",
	})
}

func convertPrizeResponse(obj models.Prize) models.PrizeToJson {
	return models.PrizeToJson{
		ID:         obj.ID,
		IndexPrize: obj.IndexPrize,
		PrizeName:  obj.PrizeName,
		Qouta:      obj.Qouta,
		DayLimit:   obj.DayLimit,
		// CreatedAt:  obj.CreatedAt,
	}
}
