package shouqianba

import (
	"sort"
	"strings"
)

func sortMap(data map[string]string) string {
	var sortStrArr []string

	for k, v := range data {
		sortStrArr = append(sortStrArr, k+"="+v)
	}
	sort.Strings(sortStrArr)
	sortStr := strings.Join(sortStrArr, "&")

	return sortStr
}
