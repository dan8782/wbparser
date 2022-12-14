package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Catalogs []struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	Shard   string `json:"shard,omitempty"`
	Query   string `json:"query,omitempty"`
	Landing bool   `json:"landing,omitempty"`
	Childs  []struct {
		ID     int    `json:"id"`
		Parent int    `json:"parent"`
		Name   string `json:"name"`
		Seo    string `json:"seo,omitempty"`
		URL    string `json:"url"`
		Shard  string `json:"shard"`
		Query  string `json:"query"`
		Childs []struct {
			ID     int    `json:"id"`
			Parent int    `json:"parent"`
			Name   string `json:"name"`
			URL    string `json:"url"`
			Shard  string `json:"shard"`
			Query  string `json:"query"`
			Seo    string `json:"seo,omitempty"`
		} `json:"childs,omitempty"`
	} `json:"childs,omitempty"`
	Seo        string `json:"seo,omitempty"`
	IsDenyLink bool   `json:"isDenyLink,omitempty"`
	Dest       []int  `json:"dest,omitempty"`
}

type Response struct {
	State   int `json:"state"`
	Version int `json:"version"`
	Data    struct {
		Products []struct {
			Sort            int    `json:"__sort"`
			Ksort           int    `json:"ksort"`
			Time1           int    `json:"time1"`
			Time2           int    `json:"time2"`
			ID              int    `json:"id"`
			Root            int    `json:"root"`
			KindID          int    `json:"kindId"`
			SubjectID       int    `json:"subjectId"`
			SubjectParentID int    `json:"subjectParentId"`
			Name            string `json:"name"`
			Brand           string `json:"brand"`
			BrandID         int    `json:"brandId"`
			SiteBrandID     int    `json:"siteBrandId"`
			Sale            int    `json:"sale"`
			PriceU          int    `json:"priceU"`
			SalePriceU      int    `json:"salePriceU"`
			AveragePrice    int    `json:"averagePrice"`
			Benefit         int    `json:"benefit"`
			Pics            int    `json:"pics"`
			Rating          int    `json:"rating"`
			Feedbacks       int    `json:"feedbacks"`
			Colors          []struct {
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"colors"`
			Sizes []struct {
				Name     string `json:"name"`
				OrigName string `json:"origName"`
				Rank     int    `json:"rank"`
				OptionID int    `json:"optionId"`
			} `json:"sizes"`
			DiffPrice    bool   `json:"diffPrice"`
			PanelPromoID int    `json:"panelPromoId,omitempty"`
			PromoTextCat string `json:"promoTextCat,omitempty"`
			IsNew        bool   `json:"isNew,omitempty"`
		} `json:"products"`
	} `json:"data"`
}

func main() {
	var adress string
	fmt.Scanf("%s",&adress)
	fmt.Println(get_catalog(adress))
}

func wf(text string) int{
	f, err := os.OpenFile("catalog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text+"\n"); err != nil {
		panic(err)
	}
	return 0;
}

func get_catalog(adress string) string {
	url := "https://static.wbstatic.net/data/main-menu-ru-ru.json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result Catalogs
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	for i := 0; i < len(result); i++ {
		fmt.Println(len(result[i].Childs))
		//_ = ioutil.WriteFile("catalogs.json", []byte(PrettyPrint(result)), 0644)
		for j := 0; j < len(result[i].Childs); j++ {
			if result[i].Childs[j].URL == adress {
				var catalogUrl string
				Query := result[i].Childs[j].Query
				Shard := result[i].Childs[j].Shard
				for page := 1; page < 101; page++ {
					pagei:= strconv.Itoa(page)
					catalogUrl = "https://catalog.wb.ru/catalog/" + Shard + "/catalog?appType=1&couponsGeo=12,3,18,15,21,101&curr=rub&dest=-1029256,-51490,-184106,123585599&emp=0&lang=ru&locale=ru&page="+pagei+"&pricemarginCoeff=1.0&reg=0&regions=68,64,83,4,38,80,33,70,82,86,75,30,69,1,48,22,66,31,40,71&sort=popular&spp=0&" + Query
					wf(catalogUrl)
					pageint:=strtoint(pagei)
					parse(catalogUrl, pageint)
				}
				return catalogUrl
			}
		}
	}
	return "I cant find this adress"
}

func strtoint(nr string) int{
	r,err:=strconv.Atoi(nr)
	if err != nil {
		fmt.Println("strconv error")
		panic(err)
	}
	return r;
}

func parse(url string,page int){
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	for i := 0; i < 99; i++ {
		wf(result.Data.Products[i].Brand)
	}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
