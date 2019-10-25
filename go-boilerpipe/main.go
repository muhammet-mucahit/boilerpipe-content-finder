package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/muhammet-mucahit/boilerplate"

	goose "github.com/advancedlogic/GoOse"
	"github.com/gin-contrib/pprof"

	"github.com/gin-gonic/gin"
)

// Used for replacing some characters in content
var replacer = strings.NewReplacer(" p ", "", "\t", " ", "\n", " ", "  ", " ")

func wordCount(value string) int {
	// Match non-space character sequences.
	re := regexp.MustCompile(`[\S]+`)

	// Find all matches and return count.
	results := re.FindAllString(value, -1)
	return len(results)
}

func boilerplateGoOse(g goose.Goose, url string, ch chan *boilerplate.ResultFormData) {
	article, err := g.ExtractFromURL(url)
	println("Language: " + article.MetaLang)

	ch <- &boilerplate.ResultFormData{
		URL:   url,
		Error: err,
		Result: boilerplate.Result{
			Title:       article.Title,
			Description: article.MetaDescription,
			H1:          article.H1,
			Content:     replacer.Replace(article.CleanedText),
			WordCount:   wordCount(article.CleanedText),
		},
	}
}

func findGoOse(g goose.Goose, urls []string) []*boilerplate.ResultFormData {
	ch := make(chan *boilerplate.ResultFormData)
	responses := []*boilerplate.ResultFormData{}

	for _, url := range urls {
		go boilerplateGoOse(g, url, ch)
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.URL)
			// fmt.Println(r.result)
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		default:
			// fmt.Printf(".")
			// time.Sleep(50 * time.Millisecond)
		}
	}
}

func goOse(c *gin.Context) {
	var json boilerplate.FormData
	c.ShouldBind(&json)
	g := goose.New()
	results := findGoOse(g, json.URLs)
	c.JSON(200, results)
	return
}

func boilerpipeContentFinder(c *gin.Context) {
	var json boilerplate.FormData
	c.ShouldBind(&json)

	cf := boilerplate.ContentFinder{}
	results := cf.Find(json.URLs)

	c.JSON(200, results)
	return
}

func main() {
	gin.SetMode(gin.DebugMode)

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

	// Grouping routes. Example Request: api.domain.com/v1/get_serp
	v1 := r.Group("/api/v1")
	{
		v1.POST("/boilerpipe", boilerpipeContentFinder)
		v1.POST("/goose", goOse)
	}
	pprof.Register(r)

	// Run with port
	r.Run(":5000")
}
