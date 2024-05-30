package handler

import (
	"encoding/json"
	"net/http"
	"url-shortner/utils"
)


func RedirectHandler(res http.ResponseWriter,req *http.Request){
	id:=req.URL.Path[len("/redirect/"):]
	url,err:=utils.GetUrl(id)

	if err!=nil{
		http.Error(res,"Invalid Request",http.StatusNotFound)
	}

	http.Redirect(res,req,url,http.StatusFound)
}

func ShortenUrl(res http.ResponseWriter,req *http.Request){
	decoder := json.NewDecoder(req.Body)
	var data struct{
		Url string
	}
	
	decoder.Decode(&data)

	hash:=utils.CreateURL(data.Url)

	response:= struct{
		Short_url string `json:"short_url"`
		}{Short_url:hash}

	res.Header().Set("Content-Type", "application/json")

	json.NewEncoder(res).Encode(response)

}