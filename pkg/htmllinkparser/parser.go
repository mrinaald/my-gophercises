package htmllinkparser

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Represents a link (<a href="..."/>) in a HTML document
type Link struct {
	Href string
	Text string
}

// ParseLinks will take in a HTML document and will return a
// slice of links parsed from it.
func ParseLinks(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	links := make([]Link, 0)

	traverseDocumentTree(doc, &links)

	return links, nil
}

func traverseDocumentTree(n *html.Node, links *[]Link) {
	if n.Type == html.ErrorNode || n.Type == html.TextNode || n.Type == html.CommentNode || n.Type == html.RawNode {
		return
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		link := buildLink(n)
		*links = append(*links, link)
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseDocumentTree(c, links)
	}
}

func buildLink(n *html.Node) Link {
	var link Link

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
			break
		}
	}

	link.Text = extractAllText(n)
	return link
}

func extractAllText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}

	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractAllText(c)
	}

	return strings.Join(strings.Fields(text), " ")
}

func typeCodeToType(t html.NodeType) string {
	if t == html.ErrorNode {
		return fmt.Sprintf("%v (ErrorNode)", t)
	}
	if t == html.TextNode {
		return fmt.Sprintf("%v (TextNode)", t)
	}
	if t == html.DocumentNode {
		return fmt.Sprintf("%v (DocumentNode)", t)
	}
	if t == html.ElementNode {
		return fmt.Sprintf("%v (ElementNode)", t)
	}
	if t == html.CommentNode {
		return fmt.Sprintf("%v (CommentNode)", t)
	}
	if t == html.DoctypeNode {
		return fmt.Sprintf("%v (DoctypeNode)", t)
	}
	if t == html.RawNode {
		return fmt.Sprintf("%v (RawNode)", t)
	}
	return "Unknown"
}
