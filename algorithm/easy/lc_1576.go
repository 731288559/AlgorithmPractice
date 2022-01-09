package easy

func modifyString(s string) string {
	bts := []byte(s)

	for idx, v := range bts {
		if v == '?' {
			for i := 'a'; i < 'a'+26; i++ {
				if idx == 0 {
					if len(bts) == 1 {
						return string(i)
					}
					if byte(i) != bts[idx+1] {
						bts[idx] = byte(i)
						break
					}
				} else if idx == len(bts)-1 {
					if byte(i) != bts[idx-1] {
						bts[idx] = byte(i)
						break
					}
				} else {
					if byte(i) != bts[idx+1] && byte(i) != bts[idx-1] {
						bts[idx] = byte(i)
						break
					}
				}
			}
		}
	}
	return string(bts)
}

func T_LC1576() {
	println(modifyString("?zs"))
	println(modifyString("ubv?w"))
	println(modifyString("j?qg??b"))
	println(modifyString("??yw?ipkj?"))
	println(modifyString("?"))
}
