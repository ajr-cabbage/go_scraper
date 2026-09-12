package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	if h := doc.Find("h1.title").Text(); h != "" {
		return h
	} else if h := doc.Find("h2.title").Text(); h != "" {
		return h
	}

	return ""
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	if p := doc.Find("main").Find("p").First().Text(); p != "" {
		return p
	} else if p := doc.Find("p").First().Text(); p != "" {
		return p
	}

	return ""
}
