package controllers

import (
	_ "errors"
	"ginlearn/database"
	"ginlearn/models"
	"net/http"
	_ "strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// tetapkan type
type BankRepo struct {
	Db *gorm.DB
}

func Baru() *BankRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Bank{})
	return &BankRepo{Db: db}
}

//create bank
func (repository *BankRepo) CreateBank(c *gin.Context) {
	var bank models.Bank
	c.BindJSON(&bank)
	err := models.CreateBank(repository.Db, &bank)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, bank)
}

// get banks
func (repository *BankRepo) GetBanks(c *gin.Context) {
	var bank []models.Bank
	err := models.GetBanks(repository.Db, &bank)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, bank)
}