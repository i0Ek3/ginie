package ginie

import "strings"

// node denotes Trie node
type node struct {
	// pattern is the route to be matched, /p/:doc
	pattern string
	// part if part of pattern, like :doc
	part string
	// children nodes
	children []*node
	// isWild means whether to match exactly, while there
	// is a : or * contains in part, that would be true
	isWild bool
}

// matchForInsert matches a child node to insert, which
// insert the first node that matches successfully
func (n *node) matchForInsert(part string) *node {
	for _, ch := range n.children {
		if ch.part == part || ch.isWild {
			return ch
		}
	}
	return nil
}

// matchForSearch match nodes to search
func (n *node) matchForSearch(part string) []*node {
	nodes := make([]*node, 0)
	for _, ch := range n.children {
		if ch.part == part || ch.isWild {
			nodes = append(nodes, ch)
		}
	}
	return nodes
}

// insert inserts node found from each layer recursively，if
// there is no node matching the current part, create a new one
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchForInsert(part)
	if child == nil {
		// create a new node
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	// go to the next layer
	child.insert(pattern, parts, height+1)
}

// search searches node found from each layer recursively.
// When it matches the node of the X layer or matches a *,
// then failed.
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchForSearch(part)
	for _, ch := range children {
		// go to the next layer
		result := ch.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

// travel travel nodes recursively
func (n *node) travel(list *[]*node) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, ch := range n.children {
		ch.travel(list)
	}
}
