package IsAnagram

import (
	"reflect"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	splittedS := strings.Split(s, "")
	splittedT := strings.Split(t, "")
	sort.Strings(splittedS)
	sort.Strings(splittedT)
	return reflect.DeepEqual(splittedS, splittedT)
}
