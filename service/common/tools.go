package common

import (
	"os"
	"strconv"
	"strings"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func BoolToInt(a bool) int {
	if a {
		return 1
	}
	return 0
}

/* return if v1 is after v2 */
func VersionGreaterEqual(v1, v2 string) (is bool, err error) {
	if v1 == "debug" {
		return true, nil
	}
	if v2 == "debug" {
		return false, nil
	}
	v1 = strings.TrimPrefix(v1, "v")
	v2 = strings.TrimPrefix(v2, "v")
	v1 = strings.ReplaceAll(v1, "-", ".")
	v2 = strings.ReplaceAll(v2, "-", ".")
	a1 := strings.Split(v1, ".")
	a2 := strings.Split(v2, ".")
	l := Min(len(a1), len(a2))
	var vv1, vv2 int
	for i := 0; i < l; i++ {
		vv1, err = strconv.Atoi(a1[i])
		if err != nil {
			return
		}
		vv2, err = strconv.Atoi(a2[i])
		if err != nil {
			return
		}
		if vv1 < vv2 {
			return false, nil
		} else if vv1 > vv2 {
			return true, nil
		}
	}
	return len(a1) >= len(a2), nil
}

func IsInDocker() bool {
	_, err := os.Stat("/.dockerenv")
	return !os.IsNotExist(err)
}