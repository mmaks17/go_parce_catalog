package main

import (
	"fmt"
	"log"
	"strings"
	"encoding/csv"
	"os"
	"github.com/PuerkitoBio/goquery"
)

const MainURL = "https://somesite"
func parcetovar(s string){
	doc, err := goquery.NewDocument(s)
	if err != nil {
		log.Fatal(err)
	}
	var id_tov string
	var name_tov string
	var pri_tov string 

	id_tov = strings.Replace( doc.Find(".product-id").First().Text(), "ID товара: ", "",1)
	name_tov = doc.Find("h1").Contents().Text()
	pri_tov = strings.TrimSpace(doc.Find(".card__price .card__priceval .price").First().Text())

	file, err := os.OpenFile("test.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
          defer file.Close()
        if err != nil {
          os.Exit(1)
	}
	csvWriter := csv.NewWriter(file)
	csvWriter.Write([]string {id_tov,pri_tov, name_tov})
        csvWriter.Flush()
}
func parcekateg(s string){
	doc, err := goquery.NewDocument(s)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".prod__link").Each(func(index int, item *goquery.Selection) {
		tov_url,_ :=item.Attr("href")
		parcetovar(MainURL+tov_url)
    })
}
func parcemenu(s string){
	doc, err := goquery.NewDocument(s)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".fgrid__item-i").Each(func(index int, item *goquery.Selection) {
		kat_url,_ :=item.Attr("href")
		parcekateg(MainURL+kat_url + "?a=1")
    })
}

func main() {

    parcemenu(MainURL+"/catalog/")

}
