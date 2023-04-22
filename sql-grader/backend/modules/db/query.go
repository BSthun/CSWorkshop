package idb

func Where(segments ...any) string {
	var str string
	for _, segment := range segments {
		str += segment.(string) + " "
	}
	return str
}
