package utilities

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
)

func ConvertMD5hash(actual_url string)(string,error){
	hasher:=md5.New()
    _,err:=io.WriteString(hasher,actual_url)
	if err!=nil{
       log.Fatal("Failed to write to string")
	   return "",err
	}
	ans:=hex.EncodeToString(hasher.Sum(nil))
	return ans[:6], nil
}