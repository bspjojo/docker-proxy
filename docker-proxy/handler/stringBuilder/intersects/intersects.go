package intersects

func CheckIfIntersects(a1 []string, a2 []string) bool {
	for _, i := range a1 {
		for _, x := range a2 {
			if i == x {
				return true
			}
		}
	}

	return false
}
