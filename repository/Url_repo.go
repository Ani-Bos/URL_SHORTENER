package repository

import (

	"url_shortener/model"

	"gorm.io/gorm"
)
type URL_shortener interface{
	Create_code( shortner *model.Shortener_Model)error
	Get_code( id uint)(*model.Shortener_Model,error)
	Update_code(shortner *model.Shortener_Model) error
	FindByShortUrl(short_url string)(bool,error)
	FindByShortUrlToActualUrl(short_url string)(*model.Shortener_Model,error)
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

func (repo *ShortenerRepository)FindByShortUrl(shorturl string)(bool,error){
	var shortner model.Shortener_Model
	err:=repo.DB.First(&shortner.Short_url).Error
	if err==gorm.ErrRecordNotFound{
		return false,err
	}
	if err!=nil{
		return false,err
	}
	return true,nil
}

func (repo *ShortenerRepository)FindByShortUrlToActualUrl(shorturl string) (*model.Shortener_Model, error) {
	var shortner model.Shortener_Model
	err := repo.DB.
		Where("short_url = ?", shorturl).
		First(&shortner).Error
	if err != nil {
		return nil, err
	}
	return &shortner, nil
}