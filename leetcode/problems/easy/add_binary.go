package easy

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	var res string
	var c int
	for i := range a {
		if i < len(b) && b[len(b)-1-i] == '1' {
			if a[len(a)-1-i] == '0' {
				c++
			} else {
				c += 2
			}
		} else if a[len(a)-1-i] == '1' {
			c++
		}

		switch c {
		case 0:
			res = "0" + res
		case 1:
			res = "1" + res
			c--
		case 2:
			res = "0" + res
			c--
		default:
			res = "1" + res
			c -= 2
		}
	}
	if c > 0 {
		return "1" + res
	}
	return res
}
