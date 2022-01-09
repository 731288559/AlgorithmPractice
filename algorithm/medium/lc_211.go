package medium

// 211. 添加与搜索单词 - 数据结构设计

type WordDictionary struct {
	l      [26]*WordDictionary
	isWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	d := this
	for i := range word {
		if d.l[word[i]-'a'] == nil {
			d.l[word[i]-'a'] = &WordDictionary{}
		}
		d = d.l[word[i]-'a']
	}
	d.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(string, int, *WordDictionary) bool

	dfs = func(word string, idx int, node *WordDictionary) bool {
		if idx == len(word) {
			return node.isWord
		}
		if word[idx] != '.' {
			if node.l[word[idx]-'a'] != nil {
				return dfs(word, idx+1, node.l[word[idx]-'a'])
			}
			return false
		} else {
			for i := 0; i < 26; i++ {
				if node.l[i] != nil && dfs(word, idx+1, node.l[i]) {
					return true
				}
			}
		}

		return false
	}

	return dfs(word, 0, this)
}

func T_LC211() {
	obj := Constructor()
	obj.AddWord("at")
	obj.AddWord("and")
	obj.AddWord("an")
	obj.AddWord("add")
	println(obj.Search("a"))
	println(obj.Search(".at"))
	println(obj.Search("an."))
	println(obj.Search("a.d."))
	println(obj.Search("b."))
	println(obj.Search("a.d"))
	println(obj.Search("."))
}

// false
// false
// true
// false
// false
// true
// false
// 98
// 1
