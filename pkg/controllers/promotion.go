package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jedavard/gomotions/pkg/db"
	"github.com/jedavard/gomotions/pkg/models"
	"github.com/jedavard/gomotions/pkg/services"
	"github.com/jedavard/gomotions/pkg/utils"
	"log"
	"net/http"
)

// GetPromotion godoc
// @Summary      Get promotion
// @Description  Get promotion by promotion id
// @Tags         Promotions
// @Param        id path string true "Promotion ID"
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Promotion
// @Failure      400  {object}  utils.HTTPError
// @Failure      404  {object}  utils.HTTPError
// @Router       /promotions/{id} [get]
func GetPromotion(c *gin.Context) {
	id, err := utils.QueryParamId(c)
	if err != nil {
		log.Println(err)
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	log.Println("GetPromotion id:", id)

	var p models.Promotion
	db.DB.Where(&models.Promotion{ID: id}).First(&p)

	c.JSON(200, p)
}

// UploadPromotion godoc
// @Summary      Upload promotions
// @Description  Upload promotions from CSV file
// @Tags         Promotions
// @Produce      json
// @Accept       multipart/form-data
// @Param        file formData file true "CSV File"
// @Success      200  {object}  models.Promotion
// @Failure      400  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router       /promotions/bulk [post]
func UploadPromotion(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	go services.RecordPromotion(src)

	c.JSON(http.StatusOK, gin.H{"message": "CSV file uploaded successfully, records are being saved to DB, you can track the progress, and we will notify when it's finished."}) // progress and notification not implemented
}
