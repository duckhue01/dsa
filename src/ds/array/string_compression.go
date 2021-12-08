package array

// String Compression: Implement a method to perform basic string compression using the counts
// of repeated characters. For example, the string aabcccccaaa would become a2b1c5a3. If the
// "compressed" string would not become smaller than the original string, your method should return
// the original string. You can assume the string has only uppercase and lowercase letters (a - z).

func compress(str string) string {
	
	count := 1
	cur := str[0]
	lens := len(str)
	news := make([]rune, 0, lens)
	if lens < 2 {
		return str
	}
	for i := 1; i < lens; i++ {
		if len(news)+2 >= lens {
			return str
		}
		if str[i] != cur {
			news = append(news, rune(cur), rune(count+48))
			cur = str[i]
			count = 1
		} else {
			count++
		}
	}

	if len(news)+2 >= lens {
		return str
	}
	news = append(news, rune(cur), rune(count+48))
	return string(news)
}
