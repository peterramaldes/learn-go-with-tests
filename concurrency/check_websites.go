package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wb WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, u := range urls {
		results[u] = wb(u)
	}

	return results
}
