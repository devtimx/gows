package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

/*userAgent simulate a browser*/
const (
	userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"
)

var UrlList = []string{
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
	// set num workers
	nWorkers := 3
	//Create buffer chanels
	jobs := make(chan string, len(UrlList))
	results := make(chan Product, len(UrlList))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}
	for _, value := range UrlList {
		jobs <- value
	}
	close(jobs)
	for r := 0; r < len(UrlList); r++ {
		<-results
	}
}

/* Worker, receives id worker, jobs Chanel and response a product chanel*/
func Worker(id int, jobs <-chan string, results chan<- Product) {
	var scrap Product
	for job := range jobs {
		fmt.Printf("Worker with id %d started fib with %s\n", id, job)
		scrap = ScraperUrl(job)

		fmt.Printf("Worker with id %d, job %s and result: %s\n", id, job, scrap)
		results <- scrap
	}
}

/* scraperUrl resolves URL, extract data and returns a string*/
func ScraperUrl(url string) Product {
	var products Product
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", userAgent)
	resp, _ := client.Do(request)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return Product{}
	}
	doc.Each(func(i int, s *goquery.Selection) {
		products = Product{
			Name:     s.Find("#productTitle").Text(),
			Price:    s.Find("#corePrice_feature_div .a-price span.a-offscreen").Text(),
			OldPrice: s.Find("span[data-a-color='secondary'] .a-offscreen").Text(),
			Brand:    s.Find(".po-brand .a-span9 span.a-size-base").Text(),
		}
	})
	return products
}
