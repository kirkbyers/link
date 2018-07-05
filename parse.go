package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link : a link (<a> tag) in an HTML doc
type Link struct {
	Href string
	Text string
}

/*
Parse : takes html doc and return slice of links parse from it
*/
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc, "")
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
