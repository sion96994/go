package see

import (
	"strings"
	"log"
)

type node struct {
	// 待匹配路由
	pattern string
	// 路由组成
	part string
	// 子节点
	children []*node
	// 是否精确匹配
	isWild bool
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		log.Printf("matchChild 匹配节点 -> %v; 路由节点 -> %v", child.part, part)
		if child.part == part || child.isWild {
			log.Printf("matchChild 匹配成功 -> %#v", child)
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
			log.Printf("matchChildren 查找到子节点 -> %#v", child)
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	log.Printf("insert 插入节点 parts -> %v", parts)
	if len(parts) == height {
		n.pattern = pattern
		log.Printf("insert 插入完成 node.Pattern -> %v", n.pattern)
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		log.Printf("插入新节点 part -> %v", part)
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	log.Printf("search 开始节点查询 -> %v", parts)
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			log.Printf("查询失败")
			return nil
		}
		log.Printf("查询成功 pattern -> %v", n.pattern)
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		log.Printf("search 递归查找节点 -> %#v", child)
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}