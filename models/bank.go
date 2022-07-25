package models

import (
	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	Id  	 int 	`json:”id”`
	NamaBank string `json:”nama_bank”`
	Norek    string `json:”norek”`
	AtasNama string `json:”atas_nama”`
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