package main

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	type res struct {
		url string
		ok  bool
	}
	resChan := make(chan res, 100)
	defer close(resChan)

	for _, url := range urls {
		go func() {
			ok := wc(url)
			resChan <- res{url: url, ok: ok}
		}()
	}

	for i := 0; i < len(urls); i++ {
		res := <-resChan
		results[res.url] = res.ok
	}

	return results
}
