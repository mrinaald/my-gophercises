package sitemapbuilder

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/mrinaald/my-gophercises/pkg/htmllinkparser"
)

const (
	// default Sitemap namespace protocol
	xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"
)

type urlset struct {
	Xmlns string `xml:"xmlns,attr"`
	Urls  []loc  `xml:"url"`
}

type loc struct {
	Location string `xml:"loc"`
}

func BuildSitemap(urlStr string, maxDepth int) (string, error) {
	depth := 0

	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	// Queue containing urls to be processed
	var urlQueue []string
	urlQueue = append(urlQueue, urlStr)

	var links []string
	isLinkVisited := make(map[string]struct{})

	for len(urlQueue) > 0 && (maxDepth == -1 || depth <= maxDepth) {
		var nextQueue []string

		for _, u := range urlQueue {
			if _, ok := isLinkVisited[u]; ok {
				continue
			}

			links = append(links, u)
			isLinkVisited[u] = struct{}{}

			pageLinks, err := getLinksFromPage(u)
			if err != nil {
				return "", err
			}

			for _, l := range pageLinks {
				if _, ok := isLinkVisited[l]; !ok {
					nextQueue = append(nextQueue, l)
				}
			}
		}

		depth++
		urlQueue = nextQueue
	}

	sitemap := urlset{
		Xmlns: xmlns,
		Urls:  make([]loc, len(links)),
	}

	for i, l := range links {
		sitemap.Urls[i] = loc{Location: l}
	}

	data, err := xml.MarshalIndent(sitemap, "", "  ")
	if err != nil {
		return "", err
	}

	return xml.Header + string(data), nil
}

func getLinksFromPage(urlStr string) ([]string, error) {
	resp, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}

	baseUrlStr := baseUrl.String()

	links := parseLinks(resp.Body, baseUrlStr)

	return filterLinks(links, withPrefixFilterFn(baseUrlStr)), nil
}

func parseLinks(r io.Reader, baseUrl string) []string {
	links, err := htmllinkparser.ParseLinks(r)
	if err != nil {
		return nil
	}

	var ret []string
	for _, l := range links {
		switch {
		case strings.Contains(l.Href, "#"):
			continue
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, baseUrl+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}

	return ret
}

func filterLinks(links []string, filterFn func(string) bool) []string {
	var filteredLinks []string
	for _, l := range links {
		if filterFn(l) {
			filteredLinks = append(filteredLinks, l)
		}
	}

	return filteredLinks
}

func withPrefixFilterFn(baseUrl string) func(string) bool {
	return func(l string) bool {
		return l != baseUrl && strings.HasPrefix(l, baseUrl)
	}
}
