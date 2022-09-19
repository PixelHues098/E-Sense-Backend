package main

// Album holds of few important details about it.
type Album struct {
    ID     int     `json:"id,omitempty"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// NewAlbum is Album constructor.
func NewAlbum(id int, title, artist string, price float64) Album {
    return Album{id, title, artist, price}
}
