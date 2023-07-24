package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
)

func newcolwithconfig(async bool,parthreads int,tlsskip bool)(*colly.Collector){
	c := colly.NewCollector(
		colly.Async(async),
	)
	c.WithTransport(&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : tlsskip},
	})
	c.Limit(&colly.LimitRule{Parallelism: parthreads})

	return c
}

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

func main(){

	urls := []string{
		"https://www.amazon.in/Ikea-Reinforced-Polypropylene-Multipurpose-Foldable/dp/B07MFD92ZH/ref=pd_rhf_d_gw_s_pd_crcd_sccl_1_2/257-4378282-9046557?pd_rd_w=Fx6Wx&content-id=amzn1.sym.785b16db-ca40-46a3-ae75-2b38bb48d1aa&pf_rd_p=785b16db-ca40-46a3-ae75-2b38bb48d1aa&pf_rd_r=0NF985AB8KYTA2G24YPB&pd_rd_wg=zLHk6&pd_rd_r=1529b101-33a2-4bc7-8f6c-c00b600f2b40&pd_rd_i=B07MFD92ZH&psc=1",
		"https://www.amazon.in/UrbanBotanicsTM-Cold-Pressed-Moisturizing-Eyelashes-Hexane-Free/dp/B07GWY28FZ/?_encoding=UTF8&pd_rd_w=kIiWo&content-id=amzn1.sym.b7cb997f-9c87-4a35-b7f5-06020612a7f4&pf_rd_p=b7cb997f-9c87-4a35-b7f5-06020612a7f4&pf_rd_r=F6KSVMMBRTTN516NBAVV&pd_rd_wg=sKTVq&pd_rd_r=75006a11-5079-47fe-b25d-0a92111072ae&ref_=pd_gw_ref_b68b93d5-3891-4115-990c-1eec06090807",
	}

	c := newcolwithconfig(true,2,true)
	configureforamazon(c)
	for _,url := range urls{
		c.Visit(url)
	}
	c.Wait()
}