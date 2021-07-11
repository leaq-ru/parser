package dns

import "strings"

func makeDNSCacheKey(key string) string {
	return strings.Join([]string{
		"technology",
		"dns",
		key,
	}, ":")
}
