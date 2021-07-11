package technology_category

import "strings"

func makeCategoryCacheKey(key string) string {
	return strings.Join([]string{
		"technology",
		"category",
		key,
	}, ":")
}
