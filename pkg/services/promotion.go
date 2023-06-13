package services

import (
	"encoding/csv"
	"github.com/google/uuid"
	"github.com/jedavard/gomotions/pkg/db"
	"github.com/jedavard/gomotions/pkg/models"
	"github.com/jedavard/gomotions/pkg/utils"
	"io"
	"log"
	"mime/multipart"
	"strconv"
)

func RecordPromotion(src multipart.File) {
	reader := csv.NewReader(src)
	reader.FieldsPerRecord = 3 // Assuming the CSV file has exactly 3 fields: id, price, expiresAt

	// Set the chunk size for bulk insertion
	chunkSize := 1000 // This can be a config variable

	// Initialize a slice to store promotion records
	var promotions []models.Promotion

	// Read and process each row of the CSV file
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			handleUploadError(err)
		}

		// Parse the record values
		id, _ := uuid.Parse(record[0])
		price, _ := strconv.ParseFloat(record[1], 64)
		expiresAt, err := utils.FormatCsvDate(record[2])
		if err != nil {
			handleUploadError(err)
			return
		}

		// Create a new promo record
		promotion := models.Promotion{
			ID:        id,
			Price:     price,
			ExpiresAt: expiresAt,
		}

		// Append the promo record to the slice
		promotions = append(promotions, promotion)

		// Check if the chunk size is reached
		if len(promotions) >= chunkSize {
			// Save the chunk of promo records into the database
			if err := db.DB.Create(&promotions).Error; err != nil {
				handleUploadError(err)
				return
			}

			// Clear the slice for the next chunk
			promotions = nil
		}
	}

	// Save the remaining records (less than the chunk size) into the database
	if len(promotions) > 0 {
		if err := db.DB.Create(&promotions).Error; err != nil {
			handleUploadError(err)
			return
		}
	}
}

func handleUploadError(err error) {
	log.Println(err)
}
