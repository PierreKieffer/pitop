package utils

func FormatStatSlice(rawStatSlice []string) []string {
	var statSlice = make([]string, 0, len(rawStatSlice))
	for i := range rawStatSlice {
		if rawStatSlice[i] != "" {
			statSlice = append(statSlice, rawStatSlice[i])
		}
	}
	return statSlice
}
