package utils

import (
	"github.com/jinzhu/copier"
)

func CopyStructFields(a interface{}, b interface{}, fields ...string) (err error) {
	return copier.Copy(a, b)
}
