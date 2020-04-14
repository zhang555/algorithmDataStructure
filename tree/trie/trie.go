package trie

type Trie struct {
	root Node
}

type Node struct {
	Byte     byte
	Word     bool
	Children []*Node
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	insert(word, &this.root)
}

func insert(word string, node *Node) {
	if word == `` {
		return
	}
	have := false
	for _, v := range node.Children {
		if v.Byte == word[0] {
			have = true
			if len(word) == 1 {
				v.Word = true
			}
			insert(word[1:], v)
			return
		}
	}

	if !have {
		n := Node{Byte: word[0]}
		if len(word) == 1 {
			n.Word = true
		}
		node.Children = append(node.Children, &n)
		insert(word[1:], &n)
	}
}

func (this *Trie) Search(word string) bool {
	return search(word, &this.root)
}

func search(word string, node *Node) bool {
	if word == `` {
		return false
	}
	for _, v := range node.Children {
		if v.Byte == word[0] {
			if len(word) == 1 && v.Word {
				return true
			}
			return search(word[1:], v)
		}
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	return startsWith(prefix, &this.root)
}

func startsWith(word string, node *Node) bool {
	if word == `` {
		return false
	}
	for _, v := range node.Children {
		if v.Byte == word[0] {
			if len(word) == 1 {
				return true
			}

			if len(word) == 0 {
				return false
			}

			return startsWith(word[1:], v)
		}
	}
	return false
}
