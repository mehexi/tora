package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	
	"github.com/mehezi/tora/internal/ui"
)

func SearchYTS(query string) ([]ui.Torrent, error) {
	apiURL := "https://yts.mx/api/v2/list_movies.json?limit=50"
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
			Movies []struct {
				TitleLong string `json:"title_long"`
				Torrents  []struct {
					Hash      string `json:"hash"`
					Quality   string `json:"quality"`
					Type      string `json:"type"`
					SizeBytes int64  `json:"size_bytes"`
					Seeds     int    `json:"seeds"`
					Peers     int    `json:"peers"`
				} `json:"torrents"`
			} `json:"movies"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var torrents []ui.Torrent
	for _, movie := range result.Data.Movies {
		for _, t := range movie.Torrents {
			name := fmt.Sprintf("%s [%s %s]", movie.TitleLong, t.Quality, t.Type)
			magnet := fmt.Sprintf("magnet:?xt=urn:btih:%s&dn=%s", t.Hash, url.QueryEscape(name))
			
			// Convert size to string with GB/MB format
			sizeGB := float64(t.SizeBytes) / (1024 * 1024 * 1024)
			var sizeStr string
			if sizeGB >= 1 {
				sizeStr = fmt.Sprintf("%.2f GB", sizeGB)
			} else {
				sizeMB := float64(t.SizeBytes) / (1024 * 1024)
				sizeStr = fmt.Sprintf("%.2f MB", sizeMB)
			}

			torrents = append(torrents, ui.Torrent{
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

	return torrents, nil
}
