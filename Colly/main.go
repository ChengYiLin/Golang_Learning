package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/utility/urlParser"
)

func main() {
	pageURL := "https://www.cakeresume.com/companies/StarkTech/jobs/front-end-engineer-ad1aa5"
	visitURL := urlParser.AppendQueryString(pageURL)

	c := colly.NewCollector(
		colly.AllowedDomains("www.cakeresume.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("------------------\nVisiting %s\n------------------\n", r.URL.String())
	})

	// c.OnHTML(".job-description-section", func(h *colly.HTMLElement) {

	// })

	// 職缺描述
	// c.OnHTML(".job-description", func(h *colly.HTMLElement) {
	// 	fmt.Println(h)
	// })

	// 職缺描述
	// c.OnHTML(".requirements", func(h *colly.HTMLElement) {
	// 	fmt.Println(h)
	// })

	c.OnHTML(".job-meta-section", func(h *colly.HTMLElement) {
		// 薪資範圍
		salaryText := h.DOM.Find(".job-salary").Text()
		// salaryText = salaryText
		fmt.Println(salaryText)
		fmt.Println("------------------")

		// 技能 tag
		h.DOM.Find(".labels .label").Each(func(i int, s *goquery.Selection) {
			fmt.Println(s.Text())
		})
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Request URL : ", r.Request.URL, "\nfailed with response : ", r, "\nError : ", e)
	})

	c.Visit(visitURL.String())
}
