package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Star struct {
	Rank, Name, Nick string
}

func main() {

	fmt.Println("Generating new Collector...")
	c := colly.NewCollector()

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	var stars []Star

	c.OnHTML(".row__top", func(e *colly.HTMLElement) {

		rank := e.ChildText(".rank")
		name := e.ChildText(".contributor__title")
		nick := e.ChildText(".contributor__name")

		star := Star{
			Rank: rank,
			Name: name,
			Nick: nick,
		}

		stars = append(stars, star)
	})

	fmt.Println("Trying to visit...")
	c.Visit("https://hypeauditor.com/top-instagram-all-russia/")

	file, err := os.Create("Stars.csv")
	if err != nil {
		log.Fatal("Failed to create CSV file", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"rank",
		"name",
		"nick",
	}

	writer.Write(headers)

	for _, star := range stars {
		record := []string{
			star.Rank,
			star.Name,
			star.Nick,
		}
		writer.Write(record)
	}

}
