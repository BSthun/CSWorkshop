package value

func BoolCount[v int8 | uint8](bool ...bool) v {
	var n v
	for _, v := range bool {
		if v {
			n++
		}
	}
	return n
}
