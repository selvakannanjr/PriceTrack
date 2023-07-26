package collector

import (
	"crypto/tls"
	"net/http"

	"github.com/gocolly/colly"
)

func NewColWithConfig(async bool,parthreads int,tlsskip bool,ecomsite string)(*colly.Collector){
	c := colly.NewCollector(
		colly.Async(async),
	)
	c.WithTransport(&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : tlsskip},
	})
	c.Limit(&colly.LimitRule{Parallelism: parthreads})

	if ecomsite=="amz"{
		configureforamazon(c)
	}

	return c
}