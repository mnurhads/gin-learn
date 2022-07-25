package controllers

import (
	"errors"
	"ginlearn/database"
	"ginlearn/models"
	"net/http"
	"strconv"
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

// get bank by id
func (repository *BankRepo) GetBankById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var bankId models.Bank
	err := models.GetBankById(repository.Db, &bankId, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, bankId)
}

func (repository *BankRepo) UpdateBank(c *gin.Context) {
	var bank models.Bank
	id, _ := strconv.Atoi(c.Param("id"))
	err   := models.GetBankById(repository.Db, &bank, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&bank)
	err = models.UpdateBank(repository.Db, &bank)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, bank)
}

func (repository *BankRepo) DeleteBank(c *gin.Context) {
	var bankDel models.Bank
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteBank(repository.Db, &bankDel, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message" : "Bank Deleted Succesfully.!"})
}