package models

import (
	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	Id  	 int 	`json:”id”`
	NamaBank string `json:”namabank”`
	Norek    string `json:”norek”`
	AtasNama string `json:”atasnama”`
	Saldo    int    `json:”saldo”`
}

func CreateBank(db *gorm.DB, Bank *Bank) (err error) {
	err = db.Create(Bank).Error
	if err != nil {
		return err
	}
	return nil
}

func GetBanks(db *gorm.DB, Bank *[]Bank) (err error) {
	err = db.Find(Bank).Error
	if err != nil {
		return err
	}
	return nil
}

func GetBankById(db *gorm.DB, Bank *Bank, id int) (err error) {
	err = db.Where("id = ?", id).First(Bank).Error
	if err != nil {
		return err
	}

	return err
}

func UpdateBank(db *gorm.DB, Bank *Bank) (err error) {
	db.Save(Bank)
	return nil
}

func DeleteBank(db *gorm.DB, Bank *Bank, id int) (err error) {
	db.Where("id = ?", id).Delete(Bank)
	return nil
}