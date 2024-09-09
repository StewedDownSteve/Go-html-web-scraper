# Web Scraper for Books(Go)
This project is a web scraper made with Golang that pulls book info from Books to Scrape. It uses the Colly library to grab stuff like titles, prices, and availability from the HTML. The cool part is it scrapes things asynchronously for speed and saves all the data in a JSON file so you can check it out easily.



![img of scraper results](https://github.com/StewedDownSteve/Go-html-web-scraper/blob/main/go-scraper-demo/Go-web-scraper-data-SH.png)

## How It's Made:

**Tech used:** Golang, Colly (Web Scraping Library), JSON

This project was one of my first tries at learning Golang and creating a web scraper. I figured out the basics of how Go is set up, learned to import and use packages like Colly, and got the hang of parsing and scraping HTML content. The scraper uses asynchronous requests to minimize scraping time, and I made sure there's error handling in place so any failed requests are handled smoothly. 

I also added features to switch between book categories on the site easily and save the scraped data into a JSON file for quick access and viewing.

## Optimizations
One big upgrade I made in this project was switching to async scraping. At first, the scraper was going through pages one at a time, which was pretty slow and clunky. Once I made it async, the time it took to scrape the data dropped a lot. I also threw in a simple way for users to pick and scrape specific book categories instead of just using fixed URLs. Adding the feature of having the scraped data being printed as a JSON file instead of just a string in the terminal is pretty helpful as well.

Looking ahead, I think could also think about adding some rate-limiting so we don’t overwhelm the server or using a headless browser for sites with dynamic content.

## Lessons Learned:

Diving into this project really helped me get the hang of Golang, especially its concurrency model, which is super useful for web scraping. I also got my first taste of working with JSON data in Go. Plus, it was a solid intro to web scraping overall—figuring out how to parse HTML, deal with pagination, and how to scrape for different data.

## Examples:
Take a look at these couple of examples that I have in my own portfolio:



