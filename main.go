package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const (
	userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"
)

var urlList = []string{
	"https://www.amazon.com.mx/dp/B07Z4681LQ",
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
		fmt.Println(product.Name)
	}
	close(chanel)
}

func scraperUrl(url string, chanel chan<- Product) {
	var product Product
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", userAgent)
	resp, _ := client.Do(request)

	name, err := inspectHtml(resp, "#productTitle")
	if err != nil {
		println(err)
		return
	}

	/*price, err := inspectHtml(resp, "spam .a-price spam .a-offscreen ")
	if err != nil {
		println(err)
		return
	}*/

	/*oldPrice, err := inspectHtml(resp, "#")
	if err != nil {
		println(err)
		return
	}*/

	/*brand, err := inspectHtml(resp, ".po-brand .a-size-base")
	if err != nil {
		println(err)
		return
	}*/

	product = Product{
		Name: name,
	}

	defer func() {
		chanel <- product
	}()
}

func inspectHtml(req *http.Response, tags string) (string, error) {
	item := ""
	doc, err := goquery.NewDocumentFromReader(req.Body)
	if err != nil {
		return "", err
	}
	doc.Find(tags).Each(func(i int, s *goquery.Selection) {
		item = s.Text()
	})
	return item, nil
}
