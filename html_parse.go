package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PageData struct {
	URL            string   `json:"url"`
	Heading        string   `json:"heading"`
	FirstParagraph string   `json:"first_paragraph"`
	OutgoingLinks  []string `json:"outgoing_links"`
	ImageURLs      []string `json:"image_urls"`
}

func normalizeURL(urlStr string) (string, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	path := parsedURL.EscapedPath()
	normURL, err := url.JoinPath(parsedURL.Host, path)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(normURL, "/"), nil
}

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	if h := doc.Find("h1").First().Text(); h != "" {
		return h
	} else if h := doc.Find("h2").First().Text(); h != "" {
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

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}
	urls := []string{}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			hrefUrl, err := url.Parse(href)
			if err != nil {
				return
			}
			mergedUrl := baseURL.ResolveReference(hrefUrl)
			urls = append(urls, mergedUrl.String())
		}
	})

	return urls, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}
	urls := []string{}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, ok := s.Attr("src")
		if ok {
			srcUrl, err := url.Parse(src)
			if err != nil {
				return
			}
			mergedUrl := baseURL.ResolveReference(srcUrl)
			urls = append(urls, mergedUrl.String())
		}
	})

	return urls, nil
}

func extractPageData(html, pageURL string) PageData {
	parsedURL, err := url.Parse(pageURL)
	if err != nil {
		return PageData{}
	}
	links, err := getURLsFromHTML(html, parsedURL)
	if err != nil {
		return PageData{}
	}
	images, err := getImagesFromHTML(html, parsedURL)
	if err != nil {
		return PageData{}
	}
	pageDat := PageData{
		URL:            pageURL,
		Heading:        getHeadingFromHTML(html),
		FirstParagraph: getFirstParagraphFromHTML(html),
		OutgoingLinks:  links,
		ImageURLs:      images,
	}

	return pageDat
}
