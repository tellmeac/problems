package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://a.com",
		"http://b.com",
		"http://c.com",
		"http://d.com",
	}

	c := make(chan string)
	go func() {
		for _, url := range urls {
			c <- url
		}
		close(c)
	}()

	s := Service{cli: http.DefaultClient}

	for url := range c {
		go func(url string) {
			ans, ok := s.Get(context.Background(), url)
			if !ok {
				fmt.Println(ans)
			}
		}(url)
	}
}

type Service struct {
	cli *http.Client
}

func (s *Service) Get(ctx context.Context, url string) (ans string, ok bool) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)

	resp, err := s.cli.Do(req)
	if err != nil {
		return fmt.Sprintf("адрес %q: not ok: %v", url, err), false
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("адрес %q: not 200, but %d", url, resp.StatusCode), false
	}

	return fmt.Sprintf("адрес %q: ok", url), true
}
