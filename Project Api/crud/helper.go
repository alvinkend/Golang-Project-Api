package crud

//GetLimitOffset nnnn
func GetLimitOffset(page, size int) (int, int) {
	var offset int

	if page == 0 || size == 0 {
		size = -1
		offset = -1
		return size, offset
	}
	offset = (page - 1) * size
	return size, offset

}
