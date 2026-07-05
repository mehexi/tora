package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type eztvTorrent struct {
	Hash      string `json:"hash"`
	Filename  string `json:"filename"`
	Title     string `json:"title"`
	SizeBytes string `json:"size_bytes"`
	Seeds     int    `json:"seeds"`
	Peers     int    `json:"peers"`
}

type eztvResponse struct {
	Torrents []eztvTorrent `json:"torrents"`
}

func fetchEzTV(host, query string) ([]Torrent, error) {
	apiURL := fmt.Sprintf("https://%s/api/get-torrents", host)
	if query != "" {
		apiURL += "?imdb_id=" + url.QueryEscape(query)
	}

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result eztvResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return parseEzTV(result.Torrents), nil
}

func parseEzTV(torrents []eztvTorrent) []Torrent {
	out := make([]Torrent, 0, len(torrents))
	for _, t := range torrents {
		magnet := fmt.Sprintf("magnet:?xt=urn:btih:%s&dn=%s", t.Hash, url.QueryEscape(t.Filename))

		bytes, _ := strconv.ParseFloat(t.SizeBytes, 64)
		sizeGB := bytes / (1024 * 1024 * 1024)
		var sizeStr string
		if sizeGB >= 1 {
			sizeStr = fmt.Sprintf("%.2f GB", sizeGB)
		} else {
			sizeMB := bytes / (1024 * 1024)
			sizeStr = fmt.Sprintf("%.2f MB", sizeMB)
		}

		out = append(out, Torrent{
			Name:     t.Title,
			Size:     sizeStr,
			Seeders:  t.Seeds,
			Leechers: t.Peers,
			Magnet:   magnet,
			Category: "TV",
			Source:   "EZTV",
		})
	}
	return out
}

func SearchEzTV() ([]Torrent, error) {
	hosts := []string{"eztvx.to", "eztv.re"}
	var lastErr error
	for _, host := range hosts {
		torrents, err := fetchEzTV(host, "")
		if err != nil {
			lastErr = err
			continue
		}
		return torrents, nil
	}
	return nil, lastErr
}
