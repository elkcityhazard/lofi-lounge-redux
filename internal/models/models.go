package models

import "time"

type Song struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Year        int         `json:"year"`
	ReleaseDate time.Time   `json:"release_date"`
	SongLength  int         `json:"song_length"`
	Rating      int         `json:"rating"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	SongGenre   []SongGenre `json:"song_genre"`
}

type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SongGenre struct {
	ID        int       `json:"id"`
	SongID    int       `json:"song_id"`
	GenreID   int       `json:"genre_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Genre     Genre     `json:"genre"`
}

type User struct {
	ID          int     `json:"id"`
	Email       string  `json:"email_address"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Password    string  `json:"-"`
	Albums      []Album `json:"albums"`
	Songs       []Song  `json:"songs"`
	Genre       Genre   `json:"genre"`
	TwitterName string  `json:"twitter_name"`
	TwitterURL  string  `json:"twitter_url"`
	TikTokName  string  `json:"tiktok_name"`
	TikTokURL   string  `json:"tiktok_url"`
	SuperUser   bool    `json:"-"`
	IsLoggedIn  bool    `json:"-"`
}

type Album struct {
	ID        int       `json:"id"`
	Title     string    `json:"album_title"`
	Length    time.Time `json:"album_length""`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tracks    []Song    `json:"track_list"`
}

type Error struct {
	Status  int
	Message error
}
