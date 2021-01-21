package main

import (
	"fmt"
	"strconv"

	"./utils"
)

var csvUtils = utils.NewCsvUtils()
var httpUtils = utils.NewHttpUtils()
var releaseUtils = utils.NewReleaseUtils()
var stringUtils = utils.NewStringUtils()

func main() {
	pageCount := extractPageCount()
	
	allReleases := make(map[string]string)
	for i:= 1; i <= pageCount; i++ {
		fmt.Printf("Processing page %d of %d\n", i, pageCount)
		releases := extractReleases(i)

		for k, v := range releases {
			allReleases[k] = v
		}
	}

	csvUtils.WriteFile("results.csv", allReleases)

	fmt.Println("Done")
}

func extractPageCount() int {
	html, httpError := httpUtils.Get("https://funxd.site")

	if httpError != nil {
		panic("Failed to fetch web page for page count extraction")
	}

	pageCount := stringUtils.SubstringAfter(html, "class=\"page-numbers\"")
	pageCount = stringUtils.SubstringBefore(pageCount, "</a>")
	pageCount = stringUtils.SubstringAfter(pageCount, "</span>")

	ret, conversionError := strconv.Atoi(pageCount)

	if conversionError != nil {
		panic("Failed to extract page count")
	}

	return ret
}

func extractReleases(pageCount int) map[string]string {
	link := "https://funxd.site/page/" + strconv.Itoa(pageCount) + "/"
	html, httpError := httpUtils.Get(link)

	if httpError != nil {
		panic("Failed to fetch web page for release extraction from " + link)
	}

	html = stringUtils.ReplaceAll(html, "\n", "")
	return releaseUtils.ExtractAllReleases(html)
}
