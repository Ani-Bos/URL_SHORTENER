package service
import (
	"url_shortener/logger"
	"net/url"
	"url_shortener/model"
	"url_shortener/utilities"
	"url_shortener/repository"
)

type ShortenerService struct{
 Repo repository.URL_shortener
 L *logger.Logger
}

func (srvc *ShortenerService)CreateShortUrl(shortner *model.Shortener_Model){
	//validate long url
	if !validate_long_url(shortner.Actual_url){
		srvc.L.LogFatalMessage("Failed to validate actual long url")
		return
	}
	//insert into db to fetch id
	//url swhorturl is base62id
	err:=srvc.Repo.Create_code(shortner)
	if err!=nil{
		srvc.L.LogFatalMessage("Failed to insert to database")
		return
	}
	id:=utilities.EncodeBase62(int(shortner.ID))
	shortner.Short_url="http://localhost:8080/get_url/"+id
	err=srvc.Repo.Update_code(shortner)
	if err!=nil{
		srvc.L.LogFatalMessage("Failed to update the table with new shorturl")
		return
	}
	
}

func validate_long_url(actual_url string)bool{
	u,err:=url.Parse(actual_url)
	return err == nil && u.Scheme != "" && u.Host != ""
}


func (srvc *ShortenerService)GetUrl(short_url string)(string,error){
	srvc.L.LogMessage("Entering into fetching long url from short url")
	short_url_id:=utilities.DecodeBase62(short_url)
	shrtn_mdl,err:=srvc.Repo.Get_code(uint(short_url_id))
	if err!=nil{
		srvc.L.LogFatalMessage("Failed to retrieve actual url from short url")
		return "", err
	}
	return shrtn_mdl.Actual_url,nil
}