package fetcher

import (
	"io"
	"net/http"
)

func Fetch(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
