package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

type extractedJob struct {
	id       string
	title    string
	company  string
	location string
	summary  string
}

func main() {
	totalPages := getPageNum()

	for i := 0; i < totalPages; i++ {
		getPageInfo(i)
	}

}

func getPageNum() int {
	pages := 0

	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length() //find all the links
	})

	return pages
}

func getPageInfo(page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jcs-JobTitle")
	searchCards.Each(func(i int, card *goquery.Selection) {
		extractJob(card)
	})

}

func extractJob(card *goquery.Selection) {
	//id, _ := card.Attr("data-jk")
	//title := cleanString(card.Find(".title").Text())
	//location := cleanString(card.Find(".companyLocation").Text())
	//fmt.Println(id, title, location)

	id, _ := card.Attr("data-jk")
	title := card.Find(".title").Text()
	location := cleanString(card.Find(".companyLocation").Text())

	fmt.Println(id, title, location)
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")

}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("status code error: ", res.StatusCode, res.Status)
	}
}
