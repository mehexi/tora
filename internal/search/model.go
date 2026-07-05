package search

type Torrent struct {
	Name     string
	Size     string
	Seeders  int
	Leechers int
	Magnet   string
	Category string
	Source   string
}

type apiMovie struct {
	TitleLong string       `json:"title_long"`
	Torrents  []apiTorrent `json:"torrents"`
}

type apiTorrent struct {
	Hash      string `json:"hash"`
	Quality   string `json:"quality"`
	Type      string `json:"type"`
	SizeBytes int64  `json:"size_bytes"`
	Seeds     int    `json:"seeds"`
	Peers     int    `json:"peers"`
}
