package linkparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href, Text string
}

func ParseFromReader(r io.Reader) ([]Link, error) {
	rootNodePtr, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	return performBFS(rootNodePtr)
}

func performBFS(root *html.Node) ([]Link, error) {
	var links []Link

	queue := []*html.Node{root}
	for len(queue) != 0 {
		currNodePtr := queue[0]
		queue = queue[1:]

		if currNodePtr.Type == html.ElementNode && currNodePtr.Data == "a" {
			links = append(links, Link{
				Href: getLinkFromAttrs(currNodePtr.Attr),
				Text: getTextFromDFS(currNodePtr),
			})
			continue
		}

		nodePtrToEnqueue := currNodePtr.FirstChild
		for nodePtrToEnqueue != nil {
			queue = append(queue, nodePtrToEnqueue)
			nodePtrToEnqueue = nodePtrToEnqueue.NextSibling
		}
	}

	return links, nil
}

func getLinkFromAttrs(attrs []html.Attribute) string {
	for _, a := range attrs {
		if a.Key == "href" {
			return a.Val
		}
	}

	return ""
}

func getTextFromDFS(root *html.Node) string {
	textChunks := []string{}

	stack := []*html.Node{}
	nodePtrToPush := root.LastChild

	for nodePtrToPush != nil {
		stack = append(stack, nodePtrToPush)
		nodePtrToPush = nodePtrToPush.PrevSibling
	}

	for len(stack) != 0 {
		currNodePtr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if currNodePtr.Type == html.TextNode {
			textChunks = append(textChunks, strings.TrimSpace(currNodePtr.Data))
		}

		nodePtrToPush = currNodePtr.LastChild
		for nodePtrToPush != nil {
			stack = append(stack, nodePtrToPush)
			nodePtrToPush = nodePtrToPush.PrevSibling
		}
	}

	return strings.Join(textChunks, " ")
}
