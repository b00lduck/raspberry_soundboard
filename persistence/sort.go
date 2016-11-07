package persistence

type ByNumPlayed []Sound

func (s ByNumPlayed) Len() int {
	return len(s)
}
func (s ByNumPlayed) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByNumPlayed) Less(i, j int) bool {
	return s[i].Count > s[j].Count
}
