package main

import(
	"github.com/gocolly/colly"
	"fmt"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)

	c.OnRequest(func(r *colly.Request){
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
		fmt.Println("Visitind",r.URL)
	})

	c.OnResponse(func(r *colly.Response){
		fmt.Println("Response Code",r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error){
		fmt.Println("Error Code",err.Error())
	})

	c.Visit("https://www.quotes.toscrape.com/random")
}