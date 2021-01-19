package utils

func FormatStatSlice(rawStatSlice []string) []string {
	var statSlice []string
	for _, stat := range rawStatSlice {
		if stat != "" {
			statSlice = append(statSlice, stat)
		}
	}
	return statSlice
}
