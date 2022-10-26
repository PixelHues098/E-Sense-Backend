package main

// Album holds of few important details about it.
type Album struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  string `json:"price"`
}

type User struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"passwordHash"`
}

// NewAlbum is Album constructor.
func NewAlbum(id int, title, artist string, price string) Album {
	return Album{id, title, artist, price}
}

func NewUser(id int, firstName string, lastName string, username string, email string, passwordHash string) User {
	println("test")
	return User{id, firstName, lastName, username, email, passwordHash}
}
