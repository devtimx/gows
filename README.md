# gows

### Simple scraper for amazon 

This is a simple example of scraping in "https://www.amazon.com.mx", to get name, price, old price (without discount) and brand of product.

Package used: github.com/PuerkitoBio/goquery

Version Go 1.18

To test the operation, clone this repository and install go mod with

```
go mod download
```

Replace the URLs in the following code block with your URLs

```
var urlList = []string{
	"https://www.amazon.com.mx/dp/B07ZQS6LPZ",
	"https://www.amazon.com.mx/dp/B07FDF9B46",
	"https://www.amazon.com.mx/dp/B0934C48VN",
	"https://www.amazon.com.mx/dp/B09F5VC4SF",
}

```

To increase the number of workers, change this variable

```
nWorkers := 3
```

Example result

```` 
Worker with id 2 started fib with https://www.amazon.com.mx/dp/B0934C48VN
Worker with id 1 started fib with https://www.amazon.com.mx/dp/B07FDF9B46
Worker with id 0 started fib with https://www.amazon.com.mx/dp/B07ZQS6LPZ
Worker with id 0, job https://www.amazon.com.mx/dp/B07ZQS6LPZ and result: {        Apple iPhone 11 Pro Max, 64GB, Totalmente Desbloqueado - Verde Medianoche (Reacondicionado)        $12,750.00 $14,299.99 Apple}
Worker with id 0 started fib with https://www.amazon.com.mx/dp/B09F5VC4SF
Worker with id 2, job https://www.amazon.com.mx/dp/B0934C48VN and result: {        HUAWEI Display- Monitor de 23.8", Resolución 1920 * 1080, Relación de Aspecto 16:9, Color Negro        $3,198.00  HUAWEI}
Worker with id 1, job https://www.amazon.com.mx/dp/B07FDF9B46 and result: {        Bose Home Speaker 500 - Altavoz con Amazon Alexa integrada, Peso de 2.15 kg, Negro (Triple Black)        $10,999.00  Bose}
Worker with id 0, job https://www.amazon.com.mx/dp/B09F5VC4SF and result: {        SAMSUNG Galaxy-A03s 4GB_64GB Black        $3,194.00 $4,139.00 SAMSUNG}

````