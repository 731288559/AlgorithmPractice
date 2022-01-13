package medium

// 677. 键值映射
// TODO: 目前方法是将val存储在每个单词的末尾字母处，计算Sum需要dfs深度遍历，由于递归消耗内存相对大，
// 		希望改进为在每个树的节点存储当前前缀的Sum，通过新的map存储当前已存储的单词，通过delta处理修改同一单词的情况
// DONE

type TrieTree struct {
	val int
	m   []*TrieTree
}

type MapSum struct {
	val   int
	root  *TrieTree
	words map[string]int
}

func MapSumConstructor() MapSum {
	return MapSum{words: make(map[string]int), root: &TrieTree{m: make([]*TrieTree, 26)}}
}

func (this *MapSum) Insert(key string, val int) {
	delta := val
	if v, ok := this.words[key]; ok {
		delta = val - v
	}
	this.words[key] = val

	tmp := this.root.m
	for _, n := range key {
		if tmp[n-'a'] == nil {
			tmp[n-'a'] = &TrieTree{m: make([]*TrieTree, 26)}
		}
		tmp[n-'a'].val += delta
		tmp = tmp[n-'a'].m
	}
}

func (this *MapSum) Sum(prefix string) int {
	tmp := this.root
	for _, n := range prefix {
		if tmp.m[n-'a'] == nil {
			return 0
		}
		tmp = tmp.m[n-'a']
	}
	return tmp.val
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */

func T_LC677() {
	obj := MapSumConstructor()
	obj.Insert("apple", 3)
	println(obj.Sum("ap"))
	obj.Insert("app", 2)
	obj.Insert("apple", 2)
	println(obj.Sum("ap"))
}
