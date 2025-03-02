package controllers

import (
	"gocomprobem/config"
	"gocomprobem/models"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateReport(c *gin.Context) {
	var report models.Report
	userID := c.PostForm("user_id")

	userIDUint, err := strconv.ParseUint(userID, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User ID tidak valid"})
        return
    }

	// Ambil data dari form-data
	report.Category = c.PostForm("category")
	report.Title = c.PostForm("title")
	report.Description = c.PostForm("description")
	report.UserID = userIDUint

	// Validasi input wajib
	validate := validator.New()
	if err := validate.StructPartial(report,
		"Category",
		"Title",
		"Description",
	); err != nil {
		errors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validasi gagal",
			"details": errors.Error(),
		})
		return
	}

	// Handle file upload
	file, err := c.FormFile("attachment")
	if err == nil {
		filename := time.Now().Format("20060102150405") + filepath.Ext(file.Filename)
		savePath := "uploads/" + filename
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupload file"})
			return
		}
		report.Attachment = savePath
	}

	// Simpan ke db
	result := config.DB.Create(&report)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Laporan berhasil dikirim",
		"data":    report,
	})
}

func GetAllReports(c *gin.Context) {
	var reports []models.Report
	result := config.DB.Find(&reports)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": reports,
	})
}

func GetReportByID(c *gin.Context) {
	id := c.Param("id")
	var report models.Report

	result := config.DB.First(&report, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Laporan tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": report,
	})
}

func GetReportsByUser(c *gin.Context) {
    userID := c.Param("user_id")
    
    var reports []models.Report
    result := config.DB.Where("user_id = ?", userID).Find(&reports)
    
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "data": reports,
    })
}