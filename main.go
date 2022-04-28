package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

/*userAgent simulate a browser*/
const (
	userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"
)

var urlList = []string{
	//"https://www.amazon.com.mx/dp/B07Z4681LQ",
	"https://www.amazon.com.mx/dp/B07ZQS6LPZ",
	"https://www.amazon.com.mx/dp/B07FDF9B46",
	"https://www.amazon.com.mx/dp/B0934C48VN",
	"https://www.amazon.com.mx/dp/B09F5VC4SF",
}

type Product struct {
	Name     string
	Price    string
	OldPrice string
	Brand    string
}

func main() {
	var products []Product
	var wg sync.WaitGroup
	chanel := make(chan Product)

	wg.Add(len(urlList))

	for _, url := range urlList {
		go scraperUrl(url, chanel)
	}

	for range urlList {
		go func() {
			defer wg.Done()
			product := <-chanel
			products = append(products, product)
		}()
	}
	wg.Wait()
	for _, product := range products {
		fmt.Printf("Name: %s Price: %s OldProce: %s Brand: %s\n", product.Name, product.Price, product.OldPrice, product.Brand)
	}
	close(chanel)
}

/* scraperUrl resolves URL, extract data and returns a string*/
func scraperUrl(url string, chanel chan<- Product) {
	var products Product
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", userAgent)
	resp, _ := client.Do(request)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}
	doc.Each(func(i int, s *goquery.Selection) {
		products = Product{
			Name:     s.Find("#productTitle").Text(),
			Price:    s.Find("#corePrice_feature_div .a-price span.a-offscreen").Text(),
			OldPrice: s.Find("span[data-a-color='secondary'] .a-offscreen").Text(),
			Brand:    s.Find(".po-brand .a-span9 span.a-size-base").Text(),
		}
	})

	defer func() {
		chanel <- products
	}()
}
