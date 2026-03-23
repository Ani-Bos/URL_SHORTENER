package repository

import (

	"url_shortener/model"

	"gorm.io/gorm"
)
type URL_shortener interface{
	Create_code( shortner *model.Shortener_Model)error
	Get_code( id uint)(*model.Shortener_Model,error)
	Update_code(shortner *model.Shortener_Model) error
}

type ShortenerRepository struct{
  DB *gorm.DB
}

func (repo *ShortenerRepository)Create_code(shortener *model.Shortener_Model)error{
	 return repo.DB.Create(shortener).Error
}

func (repo *ShortenerRepository)Update_code(shortener *model.Shortener_Model)error{
	return repo.DB.Save(shortener).Error
}

func (repo *ShortenerRepository)Get_code(id uint)(*model.Shortener_Model,error){
	var shortener model.Shortener_Model
	err:=repo.DB.First(&shortener,id).Error
	return &shortener,err
}