package text

func FormattedText(list []string) string {
	var msg string
	last := len(list) - 1
	for i, el := range list {
		if i == last {
			msg += el
		} else {
			msg += el + ", "
		}
	}

	return msg
}
