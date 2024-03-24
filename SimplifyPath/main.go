package SimplifyPath

import "strings"

func simplifyPath(path string) string {
	var splitted = strings.Split(path, "/")
	var res []string
	for _, block := range splitted {
		if block == ".." {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else if block != "." && block != "" {
			res = append(res, block)
		}
	}
	var joined = strings.Join(append(res), "/")
	return "/" + joined
}
