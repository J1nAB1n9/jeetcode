package example1410

//双引号：字符实体为 &quot; ，对应的字符是 " 。
//单引号：字符实体为 &apos; ，对应的字符是 ' 。
//与符号：字符实体为 &amp; ，对应对的字符是 & 。
//大于号：字符实体为 &gt; ，对应的字符是 > 。
//小于号：字符实体为 &lt; ，对应的字符是 < 。
//斜线号：字符实体为 &frasl; ，对应的字符是 / 。
var frasl = []rune{'f', 'r', 'a', 's', 'l', ';'}
var quot = []rune{'q', 'u', 'o', 't', ';'}
var apos = []rune{'a', 'p', 'o', 's', ';'}
var amp = []rune{'a', 'm', 'p', ';'}
var gt = []rune{'g', 't', ';'}
var lt = []rune{'l', 't', ';'}

func equipRune(ar []rune, br []rune) bool {
	if len(ar) != len(br) {
		return false
	}

	for i := range ar {
		if ar[i] != br[i] {
			return false
		}
	}
	return true
}

func entityParser(text string) string {
	strSize := len(text)
	ans := make([]rune, 0, strSize)

	// l := 0
	rs := []rune(text)
	for i := 0; i < strSize; i++ {
		r := rs[i]
		if r == '&' {
			if i+3 < strSize {
				if equipRune(rs[i+1:i+4], gt) {
					// find
					i += 3
					r = '>'
				} else if equipRune(rs[i+1:i+4], lt) {
					// find
					i += 3
					r = '<'
				}
			}
			if i+4 < strSize {
				if equipRune(rs[i+1:i+5], amp) {
					// find
					i += 4
					r = '&'
				}
			}
			if i+5 < strSize {
				if equipRune(rs[i+1:i+6], quot) {
					// find
					i += 5
					r = '"'
				} else if equipRune(rs[i+1:i+6], apos) {
					// find
					i += 5
					r = '\''
				}
			}
			if i+6 < strSize {
				if equipRune(rs[i+1:i+7], frasl) {
					// find
					i += 6
					r = '/'
				}
			}
		}
		ans = append(ans, r)
	}
	return string(ans)
}
