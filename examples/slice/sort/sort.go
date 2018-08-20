package sort

func Sort(in []string) []string {
	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in); j++ {
			if in[i] > in[j] {
				in[i], in[j] = in[j], in[i]
			}
		}
	}
	return in
}
