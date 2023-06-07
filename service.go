package main

import (
	"database/sql"
	"fmt"
)

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

func allAlbums() ([]Album, error) {
	var albs []Album

	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		return albs, fmt.Errorf("failed to query albums %v", err)
	}

	for rows.Next() {
		var album Album
		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			return albs, fmt.Errorf("failed to scan album row: %v", err)
		}
		albs = append(albs, album)
	}

	if err := rows.Err(); err != nil {
		return albs, fmt.Errorf("error occured during rows interation: %v", err)
	}

	return albs, nil
}

func albumById(id int64) (Album, error) {
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)

	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}

func removeAlbum(id int64) (int64, error) {
	result, err := db.Exec("DELETE FROM album WHERE id = ?", id)
	if err != nil {
		return 0, fmt.Errorf("deleteAlbum: %d, %v", id, err)
	}

	deletedId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("deleteAlbum: %d, %v", deletedId, err)
	}

	return deletedId, nil
}

func alterAlbum(price float32, id int64) (int64, error) {
	result, err := db.Exec("UPDATE album SET price = ? where id = ?", price, id)
	if err != nil {
		return 0, fmt.Errorf("update album %d: %v", id, price)
	}

	id, err = result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("update album %d: %v", id, price)
	}

	return id, nil
}
