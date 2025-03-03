package utils

func RemoveDuplicates(slice []string) []string {
	seen := make(map[string]struct{})
	result := []string{}

	for _, str := range slice {
		if _, exists := seen[str]; !exists {
			seen[str] = struct{}{}
			result = append(result, str)
		}
	}

	return result
}
