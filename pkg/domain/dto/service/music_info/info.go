package music_info

type InfoRequest struct {
	Group string
	Song  string
}

type InfoResponse struct {
	ReleaseDate string
	Text        string
	Link        string
}
