package search

func Search(data []string, q string) bool {
	for i := 0; i < len(data); i++ {
		if data[i] == q {
			return true
		}
	}
	return false
}
