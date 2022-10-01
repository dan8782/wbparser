// package main

// import (
// 	//"fmt"
// 	"log"
// 	"github.com/PuerkitoBio/goquery"
// 	"github.com/geziyor/geziyor"
// 	"github.com/geziyor/geziyor/client"
// 	//"github.com/geziyor/geziyor/export"
// )

// func main() {
//     geziyor.NewGeziyor(&geziyor.Options{
//         RobotsTxtDisabled: true,
//         StartRequestsFunc: func(g *geziyor.Geziyor) {
//             g.GetRendered("https://catalog.wb.ru/catalog/bags2/catalog?appType=1&couponsGeo=12,3,18,15,21,101&curr=rub&dest=-1029256,-51490,-184106,123585599&emp=0&lang=ru&locale=ru&pricemarginCoeff=1.0&reg=0&regions=68,64,83,4,38,80,33,70,82,86,75,30,69,1,48,22,66,31,40,71&sort=popular&spp=0&subject=50", g.Opt.ParseFunc)
//         },
//         ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
//             r.HTMLDoc.Find("body").Each(func(_ int, s *goquery.Selection) {
//                 log.Println(s.Find("div.wrapper").Text())
//             })
//             //fmt.Println(string(r.Body))
//         },
//         //BrowserEndpoint: "ws://localhost:50",
//     }).Start()
//     // geziyor.NewGeziyor(&geziyor.Options{
//     //     RobotsTxtDisabled: true,
//     //     StartURLs: []string{"https://www.wildberries.ru/catalog/aksessuary/sumki-i-ryukzaki/sumki"},
//     //     ParseFunc: quotesParse,
//     //     Exporters: []export.Exporter{&export.JSON{}},
//     // }).Start()
// }

// // func quotesParse(g *geziyor.Geziyor, r *client.Response) {
// //     fmt.Println(r.HTMLDoc.Find("body"))
// //     r.HTMLDoc.Find("header__container").Each(func(i int, s *goquery.Selection) {
// //         g.Exports <- map[string]interface{}{
// //             "text":   s.Find("strong.brand-name").Text(),
// //         }
// //     })
// //     if href, ok := r.HTMLDoc.Find("div.product-card-list").Attr("href"); ok {
// //         g.Get(r.JoinURL(href), quotesParse)
// //     }
// // }

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type catalogs []struct {
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

func main(){


    page:="1"
    priceRange:="priceU=9700;100000&"
    url :="https://catalog.wb.ru/catalog/bags2/catalog?appType=1&couponsGeo=12,3,18,15,21,101&curr=rub&dest=-1029256,-51490,-184106,123585599&emp=0&lang=ru&locale=ru&page="+page+"&"+priceRange+"pricemarginCoeff=1.0&reg=0&regions=68,64,83,4,38,80,33,70,82,86,75,30,69,1,48,22,66,31,40,71&sort=popular&spp=0&subject=50"
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
    //fmt.Println(PrettyPrint(result.Data.Products))
    _ = ioutil.WriteFile("test.json", []byte(PrettyPrint(result.Data.Products[0])), 0644)
    get_catalog()
}

func get_catalog(){
    url := "https://static.wbstatic.net/data/main-menu-ru-ru.json"
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("No response from request")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    var result catalogs
    if err := json.Unmarshal(body, &result); err != nil { 
        fmt.Println("Can not unmarshal JSON")
    }
    _ = ioutil.WriteFile("catalogs.json", []byte(PrettyPrint(result[0])), 0644)
}

func PrettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}