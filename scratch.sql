

CREATE TABLE songs (
    id int AUTO_INCREMENT PRIMARY KEY UNIQUE NOT NULL,
    title VARCHAR(320) NOT NULL DEFAULT "song title",
    song_description VARCHAR(320) NOT NULL DEFAULT "song description",
    rating int,
    created_at DATETIME DEFAULT NOW(),
    updated_at DATETIME DEFAULT NOW(),
    path_to_file VARCHAR(320) UNIQUE NOT NULL DEFAULT "",
    original_filename VARCHAR(140) NOT NULL DEFAULT "",
    new_filename VARCHAR(140) NOT NULL DEFAULT "",
    filesize int,
    album_id int,
    user_id int ,
    FOREIGN KEY (album_id) REFERENCES albums(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE albums (
    id int AUTO_INCREMENT PRIMARY KEY UNIQUE NOT NULL,
    album_title VARCHAR(320) NOT NULL,
    album_description VARCHAR(320),
    album_rating int,
    album_created_at DATETIME DEFAULT NOW(),
    album_updated_at DATETIME DEFAULT NOW(),
    album_length int,
    user_id int,
    FOREIGN KEY (user_id) REFERENCES users(id)
);


CREATE TABLE genres (
    id int AUTO_INCREMENT PRIMARY KEY NOT NULL,
    title VARCHAR(100) NOT NULL,
    created_at DATETIME DEFAULT NOW(),
    updated_at DATETIME DEFAULT NOW(),
    song_id int,
    album_id int,
    FOREIGN KEY (song_id) REFERENCES songs(id),
    FOREIGN KEY (album_id) REFERENCES albums(id)
);


CREATE TABLE users (
id int AUTO_INCREMENT UNIQUE NOT NULL,
email VARCHAR(140) NOT NULL DEFAULT "mario@lofilounge.xyz" UNIQUE,
first_name VARCHAR(140) NOT NULL,
last_name VARCHAR(140) NOT NULL,
user_password BINARY(60) NOT NULL,
twitter_name VARCHAR(140),
tiktok_name VARCHAR(140),
super_user BOOL
);

CREATE TABLE comments (
    id int AUTO_INCREMENT UNIQUE NOT NULL,
    title VARCHAR(140),
    content VARCHAR(140),
    album_id int,
    song_id int,
    user_id int,
    ext_user_id int,
    FOREIGN KEY (album_id) REFERENCES albums(id),
    FOREIGN KEY (song_id) REFERENCES songs(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (ext_user_id) REFERENCES users(id)
);


INSERT INTO users (
                   email,
                   first_name,
                   last_name,
                   user_password,
                   twitter_name,
                   tiktok_name,
                   super_user) VALUES (
                                       "andrew@andrew-mccall.com",
                                       "Andrew",
                                       "McCall"

                                      );

