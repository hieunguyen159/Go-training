package helpers

func CheckExist(item string, arr []string) bool {
	for _, value := range arr {
		if item == value {
			return true
		}
	}
	return false
}
