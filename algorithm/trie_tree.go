package algorithm

type TrieTreeNode struct {
	isWord bool
	m      map[byte]*TrieTreeNode
}

func (t *TrieTreeNode) Insert(s string) {
	tmp := t.m
	for i := 0; i < len(s); i++ {
		if _, ok := tmp[s[i]]; !ok {
			tmp[s[i]] = &TrieTreeNode{isWord: false, m: make(map[byte]*TrieTreeNode)}
		}
		if i == len(s)-1 {
			tmp[s[i]].isWord = true
		} else {
			tmp = tmp[s[i]].m
		}
	}
}

func (t *TrieTreeNode) Search(s string) bool {
	tmp := t.m
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if n, ok := tmp[s[j]]; ok {
				if n.isWord {
					return true
				}
				tmp = n.m
			} else {
				tmp = t.m
				break
			}
		}
	}
	return false
}

// 返回第一个匹配的词
func (t *TrieTreeNode) GetFirst(s string) (bool, string) {
	tmp := t.m
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if n, ok := tmp[s[j]]; ok {
				if n.isWord {
					return true, s[i : j+1]
				}
				tmp = n.m
			} else {
				tmp = t.m
				break
			}
		}
	}
	return false, ""
}

// 返回匹配中最长的词
func (t *TrieTreeNode) GetLongest(s string) (isFind bool, result string) {
	tmp := t.m
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if n, ok := tmp[s[j]]; ok {
				if n.isWord {
					isFind = true
					if len(s[i:j+1]) > len(result) {
						result = s[i : j+1]
					}
				}
				tmp = n.m
			} else {
				tmp = t.m
				break
			}
		}
	}
	return
}

func (t *TrieTreeNode) StartsWith(s string) bool {
	// TODO
	return false
}

func NewTrieTree() *TrieTreeNode {
	return &TrieTreeNode{isWord: false, m: make(map[byte]*TrieTreeNode)}
}

func T_TrieTree() {
	t := NewTrieTree()
	t.Insert("bea")
	t.Insert("bear")
	t.Insert("easy")
	println(t.GetLongest("easy"))
	println(t.GetLongest("abeasy"))
	println(t.GetFirst("bebeasy"))
}
