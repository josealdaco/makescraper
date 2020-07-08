package main

import (
	"fmt"

	"github.com/gocolly/colly"
	 "encoding/json"
	 "io"
   "log"
   "os"
)

type webData struct {
	Body string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {

	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("body", func(e *colly.HTMLElement) {
                //bodyData := e.Attr("body")
				fmt.Println(e.Text)
				jsonStruct := &webData{Body:e.Text}
		// Print link
                //fmt.Printf("Text found: %q -> %s\n", e.Text, link)
			result, err := json.Marshal(jsonStruct)
			   if err != nil {
				   fmt.Println(err)
				   return
			   }

			  // fmt.Println(result)
			   err2 := WriteToFile("result.txt", string(result))
		   if err2 != nil {
	   		log.Fatal(err2)
   		}
	})


	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://news.ycombinator.com/")


}

func WriteToFile(filename string, data string) error {
    file, err := os.Create(filename) ///
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = io.WriteString(file, data)
    if err != nil {
        return err
    }
    return file.Sync()
}
