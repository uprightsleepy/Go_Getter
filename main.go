package main

import (
	"database/sql"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strconv"
	"strings"
)

import (
	_ "github.com/lib/pq"
)

type book struct {
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	ImgUrl string  `json:"imgurl"`
	Rating string  `json:"rating"`
}

func main() {

	db, err := ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
	)

	c.OnHTML("article.product_pod", func(e *colly.HTMLElement) {
		imgUrl := e.ChildAttr("div.image_container a img", "src")
		absoluteImgUrl := e.Request.AbsoluteURL(imgUrl)

		priceStr := e.ChildText("p.price_color")
		priceStr = strings.TrimPrefix(priceStr, "£")

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Printf("Failed to parse price for book: %s, error: %v", e.ChildText("h3 a"), err)
			return
		}

		class := e.ChildAttr("p.star-rating", "class")
		rating := strings.Split(class, " ")[1]

		entry := book{
			Title:  e.ChildText("h3 a"),
			Price:  price,
			ImgUrl: absoluteImgUrl,
			Rating: rating,
		}

		err = InsertBook(db, entry)
		if err != nil {
			log.Printf("Failed to insert book into database: %v", err)
		}
	})

	c.OnHTML("li.next a", func(e *colly.HTMLElement) {
		nextPage := e.Attr("href")
		fmt.Println("Visiting", e.Request.AbsoluteURL(nextPage))
		err := e.Request.Visit(nextPage)
		if err != nil {
			log.Fatal(err)
		}
	})

	err = c.Visit("https://books.toscrape.com/")
	if err != nil {
		log.Fatal(err)
	}
}
