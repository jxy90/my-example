package gee

import "strings"

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

// 单一匹配
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配
func (n *node) matchChildren(part string) (ans []*node) {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			ans = append(ans, child)
		}
	}
	return
}

// 构建路由
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			pattern:  "",
			part:     part,
			children: nil,
			isWild:   part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
