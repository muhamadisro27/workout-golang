package helper

import "strconv"

func ConvertStringToInt(categoryId string) (id int) {
	id, err := strconv.Atoi(categoryId)
	PanicIfError(err)

	return id
}
