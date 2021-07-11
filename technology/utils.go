package technology

import "strings"

func makeTechnologyCacheKey(key string) string {
	return strings.Join([]string{
		"technology",
		"technology",
		key,
	}, ":")
}
