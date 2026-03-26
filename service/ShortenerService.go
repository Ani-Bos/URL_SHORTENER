package service

import (
	"context"
	"net/url"
	"strconv"
	"time"
	"url_shortener/logger"
	"url_shortener/model"
	"url_shortener/repository"
	"url_shortener/utilities"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/redis/go-redis/v9"
)

type ShortenerService struct{
 Repo repository.URL_shortener
 L *logger.Logger
 Filter *bloom.BloomFilter
 Rdb *redis.Client
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
	vl, err := srvc.Rdb.Get(context.Background(), short_url).Result()
	if err == nil {
		return vl, nil
	}
	shrtn_mdl, err := srvc.Repo.FindByShortUrlToActualUrl(short_url)
	if err != nil {
		return "", err
	}
	return shrtn_mdl.Actual_url, nil
}

func (srvc *ShortenerService)GeneratehexShorturl(shrtnrmdl *model.Shortener_Model){
	//validate long url
	if !validate_long_url(shrtnrmdl.Actual_url){
		srvc.L.LogFatalMessage("Failed to validate actual long url")
		return
	}
	//;genrate 6 digit md5 hash 
	//if found during bloom filtert check do a retry with predefie string
	//retry logic of hash
    var hashr string
    var err error
	retryinpt:=shrtnrmdl.Actual_url
	for i:=0;;i++{
       if i>0{
           retryinpt=shrtnrmdl.Actual_url+strconv.Itoa(i)
	   }
	   hashr,err=utilities.ConvertMD5hash(retryinpt)
	   if err!=nil{
		 return
	   }
	   //check in bloom filter if not add it
	   if !srvc.Filter.Test([]byte(hashr)){
		  break
	   }
	   //false positive in bloom filter check redis
	   _,err:=srvc.Rdb.Get(context.Background(),hashr).Result()
	   if err!=redis.Nil{
             break
	   }
	   //if redis also not found
	   //check in db to confirm
	   exist,err:=srvc.Repo.FindByShortUrl(hashr)
	   if err!=nil{
		return
	   }
	   if !exist{
		break;
	   }

	}
	///save in db
	srvc.Filter.AddString(hashr)

	shrtnrmdl.Short_url =  "http://localhost:8080/get_url/"+hashr
    err=srvc.Repo.Create_code(shrtnrmdl)
	if err!=nil{
		srvc.L.LogFatalMessage("Failed to insert to database")
		return
	}
	//save in redis and bloom filter
	srvc.Rdb.Set(context.Background(), hashr, shrtnrmdl.Actual_url, 24*time.Hour)
}

