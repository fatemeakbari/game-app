package slice

func UintSliceContains(arr []uint, x uint) bool {

	for _, i := range arr {
		if i == x {
			return true
		}
	}
	return false
}
