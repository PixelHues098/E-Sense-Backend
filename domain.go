package main

// Album holds of few important details about it.
type Album struct {
    ID     int     `json:"id,omitempty"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  string `json:"price"`
}

// NewAlbum is Album constructor.
func NewAlbum(id int, title, artist string, price string) Album {
    return Album{id, title, artist, price}
}
