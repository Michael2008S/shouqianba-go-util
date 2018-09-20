package shouqianba

import (
	"math/rand"
	"sort"
	"strings"
	"time"
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

func getClient_Sn(codeLenth int) (code string) {
	s := "1234567890"
	t := time.Now().UTC()
	r := rand.New(rand.NewSource(t.UnixNano()))
	for i := 0; i < codeLenth; i++ {
		a := s[r.Intn(len(s)-1)]
		code += string(a)
	}
	return
}
