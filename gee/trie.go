package gee

import "strings"

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang;即完整路由
	part     string  // 路由中的一部分，例如 :lang；即完整路由的最后一块
	children []*node // 子节点，例如 [doc, tutorial, intro]；即下一个可能的路由
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 节点插入
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		//已到所需parttern
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	//如果没有匹配到当前part的节点，则新建一个
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// 节点查找
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			//如果匹配到:lang或*，则该nodepattern为空，返回nil
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
