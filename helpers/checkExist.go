package helpers

func CheckExist(item string, arr []string) bool {
	for _, value := range arr {
		if item == value {
			return true
		}
	}
	return false
}
func Find(val string, arr []string) (int, bool) {
	for i, item := range arr {
	    if item == val {
		   return i, true
	    }
	}
	return -1, false
 }