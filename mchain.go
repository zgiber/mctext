package mctext

import (
	"bufio"
	"io"
	"math/rand"
	"strings"
	"time"
)

type MChain struct {
	Nodes map[string]*Node
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (mc *MChain) Parse(in io.Reader) {
	wordScanner := bufio.NewScanner(in)
	wordScanner.Split(bufio.ScanWords)
	var pos int
	var last *Node
	for wordScanner.Scan() {
		word := strings.ToLower(wordScanner.Text())
		word = strings.Trim(word, "-()[]{}\"'")
		if strings.ContainsAny(word, "0123456789[](){}\"'-+/*%$") {
			// filter some unwanted words
			continue
		}

		if n, ok := mc.Nodes[word]; ok {

			if pos > 0 {
				last.AddNext(n)
				last = n
			}

		} else {
			newNode := &Node{
				value:    word,
				children: []*Node{},
				counts:   map[string]int64{},
				weights:  []float64{},
			}

			mc.Nodes[word] = newNode

			if pos > 0 {
				last.AddNext(newNode)
			}

			last = newNode
		}

		pos++
	}

	mc.generateWeights()
}

func (mc *MChain) Generate(start string, length int) string {
	rand.Seed(time.Now().UnixNano())
	return strings.Join(mc.GenerateSlice(start, length), " ")
}

func (mc *MChain) GenerateSlice(start string, length int) []string {
	var root *Node

	if n, ok := mc.Nodes[start]; ok {
		root = n
	} else {
		for _, node := range mc.Nodes {
			root = node
			break
		}
	}

	return root.Generate([]string{}, int(length))
}

func (mc *MChain) generateWeights() {
	for _, n := range mc.Nodes {
		for _, nn := range n.children {
			n.weights = append(n.weights, float64(n.counts[nn.value])/float64(n.totalCount))
		}
	}
}

type Node struct {
	value      string
	children   []*Node
	counts     map[string]int64
	weights    []float64
	totalCount int64
}

func (n *Node) AddNext(nn *Node) {

	if _, ok := n.counts[nn.value]; !ok {
		n.children = append(n.children, nn)
	}

	n.counts[nn.value]++
	n.totalCount++
}

func (n *Node) Generate(res []string, length int) []string {
	if len(res) == length || len(n.children) == 0 {
		return res
	}

	res = append(res, n.value)
	return n.Next().Generate(res, length)
}

func (n *Node) Next() *Node {
	num := rand.Float64()
	var wSum float64

	for i, w := range n.weights {
		wSum += w
		if num < wSum {
			return n.children[i]
		}
	}

	return n.children[len(n.children)-1]
}
