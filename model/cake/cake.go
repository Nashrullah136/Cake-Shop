package cake

import (
	"Pretests/database"
	"database/sql"
	"fmt"
	"time"
)

const tableName = "cakes"

//TODO: add action log

type Cake struct {
	ID          int     `json:"id,omitempty"`
	Title       string  `json:"title,omitempty" validate:"required"`
	Description string  `json:"description,omitempty,omitempty" validate:"required"`
	Rating      float64 `json:"rating,omitempty" validate:"numeric,min=1,required"`
	Image       string  `json:"image,omitempty" validate:"required,url"`
	CreatedAt   string  `json:"created_at,omitempty"`
	UpdatedAt   string  `json:"updated_at,omitempty"`
}

func convertRowsToCakes(rows *sql.Rows) ([]Cake, error) {
	var cakes []Cake
	for rows.Next() {
		var cake Cake
		if err := rows.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image,
			&cake.CreatedAt, &cake.UpdatedAt); err != nil {
			return nil, fmt.Errorf("can't get all cakes: %v", err)
		}
		cakes = append(cakes, cake)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("can't get all cakes: %v", err)
	}
	return cakes, nil
}

func FindPagination(page, rowPerPage int) ([]Cake, error) {
	if page <= 0 || rowPerPage <= 0 {
		return FindAll()
	}
	var cakes []Cake
	dbConnection := database.GetDbConnection()
	rows, err := dbConnection.Query("SELECT * FROM "+tableName+" ORDER BY rating DESC, title ASC LIMIT ?, ?",
		rowPerPage*(page-1), rowPerPage)
	if err != nil {
		return nil, fmt.Errorf("unexpected error. Can't get cakes")
	}
	defer rows.Close()
	cakes, err = convertRowsToCakes(rows)
	if err != nil {
		return nil, err
	}
	return cakes, nil
}

func FindAll() ([]Cake, error) {
	var cakes []Cake
	dbConnection := database.GetDbConnection()
	rows, err := dbConnection.Query("SELECT * FROM " + tableName + " ORDER BY rating DESC, title ASC")
	if err != nil {
		return nil, fmt.Errorf("can't get all cakes: %v", err)
	}
	defer rows.Close()
	cakes, err = convertRowsToCakes(rows)
	if err != nil {
		return nil, err
	}
	return cakes, nil
}

func FindById(id int) (Cake, error) {
	var cake Cake
	dbConnection := database.GetDbConnection()
	row := dbConnection.QueryRow("SELECT * FROM "+tableName+" WHERE id = ?", id)
	if err := row.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return cake, fmt.Errorf("cake with id %d not found", id)
		}
		return cake, fmt.Errorf("cake with id %d: %v", id, err)
	}
	return cake, nil
}

func Insert(newCake Cake) (Cake, error) {
	dbConnection := database.GetDbConnection()
	newCake.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	newCake.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	result, err := dbConnection.Exec("INSERT INTO "+tableName+" (title, description, rating, image, created_at, "+
		"updated_at) VALUES (?, ?, ?, ?, ?, ?)", newCake.Title, newCake.Description, newCake.Rating,
		newCake.Image, newCake.CreatedAt, newCake.UpdatedAt)
	if err != nil {
		return newCake, fmt.Errorf("cake insertion error: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return newCake, fmt.Errorf("addAlbum: %v", err)
	}
	newCake.ID = int(id)
	return newCake, nil
}

func Update(id int, newCake Cake) error {
	dbConnection := database.GetDbConnection()
	newCake.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	_, err := dbConnection.Exec("UPDATE "+tableName+" SET title = ?, description = ?, rating = ?, "+
		"image = ?, updated_at = ? WHERE id = ?", newCake.Title, newCake.Description, newCake.Rating,
		newCake.Image, newCake.UpdatedAt, id)
	if err != nil {
		return fmt.Errorf("cake insertion error: %v", err)
	}
	return nil
}

func Delete(id int) error {
	dbConnection := database.GetDbConnection()
	_, err := dbConnection.Exec("DELETE FROM "+tableName+" WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("cake insertion error: %v", err)
	}
	return nil
}
