package hard

import (
	"fmt"
	"strconv"
)

// 282. 给表达式添加运算符

func addOperators(num string, target int) []string {
	var result []string
	var dfs func([]byte, int, int, int)

	dfs = func(curString []byte, curResult int, curIdx int, prev int) {
		if curIdx == len(num) {
			if curResult == target {
				result = append(result, string(curString))
			}
			return
		}
		for i := curIdx; i < len(num); i++ {
			// 除0外，不会有0开头的整数
			if num[curIdx] == '0' && i != curIdx {
				return
			}
			next := num[curIdx : i+1]
			nextVal, _ := strconv.Atoi(next)
			if curIdx == 0 {
				// 二元运算符，因此开头只能为默认加法
				dfs(append(curString, next...), nextVal, i+1, nextVal)
			} else {
				dfs(append(curString, "+"+next...), curResult+nextVal, i+1, nextVal)
				dfs(append(curString, "-"+next...), curResult-nextVal, i+1, -nextVal)
				// 把上一次得到的值撤销掉，加上上次的值乘以新的值
				dfs(append(curString, "*"+next...), curResult-prev+prev*nextVal, i+1, prev*nextVal)
			}
		}
	}

	dfs(make([]byte, 0, 2*len(num)-1), 0, 0, 0)

	return result
}

func T_LC282() {
	fmt.Println(addOperators("123", 6), 6)
	fmt.Println(addOperators("232", 8), 8)
	fmt.Println(addOperators("105", 5), 5)
	fmt.Println(addOperators("00", 0), 0)
	fmt.Println(addOperators("3456237490", 9191), 9191)
}
