package models

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Song struct {
	ID               int         `json:"id"`
	Title            string      `json:"title"`
	Description      string      `json:"description"`
	ReleaseDate      time.Time   `json:"release_date"`
	SongLength       int         `json:"song_length"`
	Rating           int         `json:"rating"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	SongGenre        []SongGenre `json:"song_genre"`
	OriginalFileName string      `json:"original_filename"`
	NewFileName      string      `json:"new_filename"`
	FileSize         int64       `json:"file_size"`
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

func (u *User) GenerateSecurePassword(password string, hash string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (u *User) ValidateSecurePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *Song) InsertSong(db *sql.DB, title string, description string, rating int, createdAt time.Time, updatedAt time.Time, pathToFile string, originalFileName string, newFileName string, filesize int, albumID int, userID int) (sql.Result, error) {

	stmt := `INSERT INTO songs (
                   title, 
                   song_description, 
                   rating, 
                   created_at, 
                   updated_at, 
                   path_to_file, 
                   original_filename, 
                   new_filename, 
                   filesize,
                   album_id, 
                   user_id) VALUES (?, 
                                    ?, 
                                    ?, 
                                    NOW(), 
                                    NOW(), 
                                    ?, 
                                    ?, 
                                    ?, 
                                    ?, 
                                    ?, 
                                    ?
);`
	result, err := db.Exec(stmt, title, description, rating, createdAt, updatedAt, pathToFile, originalFileName, newFileName, filesize, albumID, userID)

	if err != nil {
		return nil, err
	}

	return result, nil
}
