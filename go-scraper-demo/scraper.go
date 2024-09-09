package main

// imports the packages: encoding data into JSON format, formatted I/O, logging errors, time-related functions, and Colly for web scraping
import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
)

// Struct to define the fields of each scraped item. each field stores said data
type Item struct {
	Link    string `json:"link"`
	Name    string `json:"name"`
	Price   string `json:"price"`
	Instock string `json:"instock"`
}

// Function to measure and print the time taken by any process
func timer(name string) func() {
	// Capture the current time when the function starts
	start := time.Now()
	return func() {
		// Print the time taken when the process finishes
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	// Defer timer start function until the end of the this main function
	defer timer("main")()

	// ========= Change this URL for the different book categories=========
	// Change this line to switch between categories
	categoryURL := "https://books.toscrape.com/catalogue/category/books/travel_2/index.html"

	// This creates a new collector(scrapper) with async
	c := colly.NewCollector(colly.Async(true))

	// Slice to store scrapped items
	items := []Item{}

	// Set a callback to be executed when matching HTML elements are found
	// extract link from HTML
	// visit the extracted link
	c.OnHTML("div.side_categories li ul li", func(h *colly.HTMLElement) {
		link := h.ChildAttr("a", "href")
		c.Visit(h.Request.AbsoluteURL(link))

	})
	// Callback to handle pagination, visiting the next page of results.
	// Visit the link to the next page
	c.OnHTML("li.next a", func(h *colly.HTMLElement) {
		c.Visit(h.Request.AbsoluteURL(h.Attr("href")))
	})
	// Callback to scrape data from each product element
	// Create a new Item struct and populate it with scraped data
	// extracts product link, producnt name, and down the list
	// Add the item to the list of scraped items
	c.OnHTML("article.product_pod", func(h *colly.HTMLElement) {
		i := Item{
			Link:    h.ChildAttr("a", "href"),
			Name:    h.ChildAttr("h3 a", "title"),
			Price:   h.ChildText("p.price_color"),
			Instock: h.ChildText("p.instock"),
		}
		items = append(items, i)
	})
	// Callback to log the URL of each visited page
	// Print the visited URL
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	// Start scraping by visiting the first page in the category
	c.Visit(categoryURL)

	// Wait for asynchronous requests to complete
	c.Wait()

	// Convert the scraped data to JSON
	// Log and exit if there is an error during JSON converssion
	data, err := json.MarshalIndent(items, " ", "")
	if err != nil {
		log.Fatal()
	}

	// ==== CHOOSE EITHER JSON as String or JSON save as a file

	// Print the JSON data as a string
	// fmt.Println(string(data))

	// I added this to save results to a JSON file
	err = os.WriteFile("scrapped_data.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//  Print confirmation that json file was created
	fmt.Println("Scraped data saved to scraped_data.json")

}

// The tutorial was using this website for scrapping:
// https://books.toscrape.com/catalogue/page-1.html
// The above link was used before we used async and category

// orginal tutorial https://www.youtube.com/watch?v=wUSgA8WEy4Q
