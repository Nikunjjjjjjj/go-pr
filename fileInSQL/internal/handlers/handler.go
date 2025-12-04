package handlers

import (
	"UploadFileInSql/internal/database"
	"UploadFileInSql/internal/models"
	"UploadFileInSql/internal/services"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

var discomService = services.DiscomService{}

func UploadCSV(c *gin.Context) {
	//return func(c *gin.Context) {

	// Read the uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "file upload failed"})
		return
	}

	// Save temporarily
	tempPath := "./temp.csv"
	if err := c.SaveUploadedFile(file, tempPath); err != nil {
		c.JSON(500, gin.H{"error": "failed to save uploaded file"})
		return
	}

	// Open CSV
	f, err := os.Open(tempPath)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to open CSV file"})
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.TrimLeadingSpace = true

	// Read header row
	_, err = reader.Read()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to read CSV header"})
		return
	}

	var offices []models.OfficeData

	// Parse each record
	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("row error:", err)
			continue
		}

		data := models.OfficeData{
			OfficeName:        row[0],
			Pincode:           row[1],
			OfficeType:        row[2],
			DeliveryStatus:    row[3],
			DivisionName:      row[4],
			RegionName:        row[5],
			CircleName:        row[6],
			Taluk:             row[7],
			DistrictName:      row[8],
			StateName:         row[9],
			Telephone:         row[10],
			RelatedSubOffice:  row[11],
			RelatedHeadOffice: row[12],
			Discom:            row[13],
		}

		offices = append(offices, data)
	}

	// Bulk insert (100 rows/batch)
	if err := database.DB.CreateInBatches(&offices, 100).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to insert data into DB"})
		return
	}

	c.JSON(200, gin.H{
		"message":       "CSV uploaded successfully",
		"rows_inserted": len(offices),
	})
	//}
}

func PincodeDetails(c *gin.Context) {
	pincode := c.Query("pin")
	//pincode empty condition
	res, err := discomService.FindByPinCode(pincode)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"response": res,
	})
}
