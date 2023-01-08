package main

import (
	"flag"
	"fmt"

	"github.com/mrinaald/my-gophercises/pkg/sitemapbuilder"
)

func main() {
	urlStr := flag.String("url", "https://calhoun.io", "The base url for the site from which to create the sitemap.")
	maxDepth := flag.Int("depth", -1, "The maximum depth to traverse for building sitemap.")

	flag.Parse()

	sitemap, err := sitemapbuilder.BuildSitemap(*urlStr, *maxDepth)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", sitemap)
}
