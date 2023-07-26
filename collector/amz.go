package collector

import (
	"fmt"

	"github.com/gocolly/colly"
)
func configureforamazon(c *colly.Collector){
	c.OnHTML("div.a-section.a-spacing-none.aok-align-center", func(e *colly.HTMLElement) {
		price := e.ChildText("span.a-price span.a-offscreen")
		fmt.Println("Price->",price)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...",r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
	})
}
