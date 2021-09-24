package oots

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/Eldius/webcomics-fetcher2-go/comics"
	"github.com/Eldius/webcomics-fetcher2-oots/helper"
	"github.com/gocolly/colly"
)

const (
	basePath = "https://www.giantitp.com/comics/oots.html"
)

/*
Fetch fetchs strips data
*/
func Fetch() []comics.ComicStrip {
	return scrap()
}

func scrap() []comics.ComicStrip {
	results := make([]comics.ComicStrip, 0)
	requestMap := make(map[string][]string)

	c := colly.NewCollector(
		colly.AllowedDomains("giantitp.com", "www.giantitp.com"),
	)

	// Find and visit all links
	c.OnHTML(".ComicList", func(e *colly.HTMLElement) {
		textSlice := strings.SplitN(e.Text, "-", 2)

		link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		requestMap[link] = append(textSlice, e.Text)

		fmt.Printf("Adding '%s' to stack...\n", link)
		_ = e.Request.Visit(link)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Requesting '%s'...\n", r.URL.String())
		variables := requestMap[r.URL.String()]
		if len(variables) > 0 {
			r.Ctx.Put("number", variables[0])
			r.Ctx.Put("title", variables[1])
			r.Ctx.Put("label", variables[2])
		}
	})

	c.OnHTML("body > table > tbody > tr:nth-child(2) > td > table > tbody > tr > td:nth-child(2) > table > tbody > tr > td > table > tbody > tr:nth-child(2) > td > img", func(e *colly.HTMLElement) {
		label := e.Request.Ctx.Get("label")
		fmt.Printf("Found strip '%s'...\n", label)
		imageURL := e.Request.AbsoluteURL(e.Attr("src"))
		results = append(results, comics.ComicStrip{
			ID:           imageURL,
			URL:          imageURL,
			Name:         label,
			WebcomicName: helper.PluginName,
			Order:        findStripOrder(label),
			FileName:     createFilename(label, imageURL),
		})
	})
	_ = c.Visit(basePath)

	c.Wait()

	fmt.Println("Finishing scrap...")
	return results
}

func findStripOrder(name string) int {
	re := regexp.MustCompile("[0-9]+")
	order, _ := strconv.Atoi(re.FindString(name))
	return order
}

func createFilename(name string, url string) string {
	return filepath.Join(helper.PluginName, comics.SanitizeFilename(name)) + "." + FindFileExtensionFromURL(url)
}

/*
FindFileExtensionFromURL finds image extension from URL
*/
func FindFileExtensionFromURL(url string) string {
	tmp := strings.Split(url, ".")
	return tmp[len(tmp)-1]
}
