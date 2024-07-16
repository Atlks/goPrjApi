package lib

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func ExtractHrefAttributes(html string) map[string]struct{} {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	hrefs := make(map[string]struct{})
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			hrefs[href] = struct{}{}
		}
	})

	return hrefs
}
