package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func fetchYTS(host, query string) ([]Torrent, error) {

	apiURL := fmt.Sprintf("https://%s/api/v2/list_movies.json?limit=50", host)

	if query != "" {
		apiURL += "&query_term=" + url.QueryEscape(query)
	}

	resp, err := http.Get(apiURL)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data struct {
			Movies []apiMovie `json:"movies"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return parseTorrents(result.Data.Movies), nil
}

func parseTorrents(movies []apiMovie) []Torrent {
	var torrents []Torrent
	for _, movie := range movies {
		for _, t := range movie.Torrents {
			name := fmt.Sprintf("%s [%s %s]", movie.TitleLong, t.Quality, t.Type)
			magnet := fmt.Sprintf("magnet:?xt=urn:btih:%s&dn=%s", t.Hash, url.QueryEscape(name))

			sizeGB := float64(t.SizeBytes) / (1024 * 1024 * 1024)
			var sizeStr string
			if sizeGB >= 1 {
				sizeStr = fmt.Sprintf("%.2f GB", sizeGB)
			} else {
				sizeMB := float64(t.SizeBytes) / (1024 * 1024)
				sizeStr = fmt.Sprintf("%.2f MB", sizeMB)
			}

			torrents = append(torrents, Torrent{
				Name:     name,
				Size:     sizeStr,
				Seeders:  t.Seeds,
				Leechers: t.Peers,
				Magnet:   magnet,
				Category: "Movies",
				Source:   "YTS",
			})
		}
	}
	return torrents
}

func SearchYTS(query string) ([]Torrent, error) {
	hosts := []string{"yts.mx", "yts.am", "yts.rs"}
	var lastErr error
	for _, host := range hosts {
		torrents, err := fetchYTS(host, query)
		if err != nil {
			lastErr = err
			continue
		}
		return torrents, nil
	}
	return nil, lastErr
}
