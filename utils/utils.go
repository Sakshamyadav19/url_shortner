package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
	"url-shortner/model"
)


var urlDB=make(map[string]model.Url)

func GetHash(url string) string {
	hasher:=sha256.New();

	hasher.Write([]byte(url))
	hashedbytes:=hasher.Sum(nil)
	hash:=hex.EncodeToString(hashedbytes)
	return hash[:9]
}

func CreateURL(url string)string{
	hash:=GetHash(url)
	urlDB[hash]=model.Url{
		Id:hash,
		OriginalUrl:url,
		ShortenUrl:hash,
		Date:time.Now(),
	}
	return hash
}

func GetUrl(id string)(string,error){
	url,ok:= urlDB[id]	

	if !ok {
		return "",errors.New("url not found")
	}
	return url.OriginalUrl,nil
}
