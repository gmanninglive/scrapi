package scraper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gmanninglive/scrapi/pkg/models"
	"github.com/gocolly/colly/v2"
)

const URL = "https://www.reviews.co.uk/company-reviews/store/yardlynk"

func RefreshReviews(res chan models.Review, done chan bool) {
	c := colly.NewCollector()

	c.OnHTML("div.Review", func(e *colly.HTMLElement) {
		author := e.DOM.Find(".Review__author")
		reviewPath, _ := author.Attr("href")
		stars := len(e.DOM.Find(".Rating__stars").Children().Filter(".stars__icon--100").Nodes)
		message := e.DOM.Find(".Review__body")
		verified := e.DOM.Find("div.BadgeElement__text:contains('Verified Buyer')")

		review := models.Review{
			Author:   strings.TrimSpace(author.Text()),
			Stars:    stars,
			Url:      fmt.Sprintf("%s%s", URL, reviewPath),
			Message:  strings.TrimSpace(message.Text()),
			Verified: verified.Text() != "",
		}

		res <- review
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		prefix := "https://www.reviews.co.uk/company-reviews/store/yardlynk/"
		href := e.Attr("href")
		if !strings.HasPrefix(href, prefix) {
			return
		}

		if _, err := strconv.Atoi(href[len(prefix):]); err == nil {
			e.Request.Visit(href)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.reviews.co.uk/company-reviews/store/yardlynk")

	done <- true
}
