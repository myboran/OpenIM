package utils

import (
	"math/rand"
	"strconv"
)

//judge a string whether in the  string list
func IsContain(target string, List []string) bool {
	for _, element := range List {

		if target == element {
			return true
		}
	}
	return false
}

func GetMsgID(sendID string) string {
	t := int64ToString(GetCurrentTimestampByNano())
	return Md5(t + sendID + int64ToString(rand.Int63n(GetCurrentTimestampByNano())))
}

func int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
