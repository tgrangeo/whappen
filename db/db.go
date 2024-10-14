package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tgrangeo/whappen/rss"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./databse.db")
	if err != nil {
		return nil, err
	}
	query := `CREATE TABLE IF NOT EXISTS article (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    title TEXT NOT NULL,
	    link TEXT NOT NULL,
	    date TEXT,
	    to_read BOOLEAN DEFAULT TRUE
	);`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InsertArticle(db *sql.DB, article rss.Article) error {
	query := `INSERT INTO article (title, link, date, to_read) 
          VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, article.Title, article.Link, article.Date, article.ToRead)
	return err
}

func InsertToRead(db *sql.DB, article rss.Article) error {
	query := `INSERT INTO article (title, link, date, to_read) 
	          VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, article.Title, article.Link, article.Date, true)
	return err
}

func RemoveArticle(db *sql.DB, articleLink string) error {
	query := `DELETE FROM article WHERE link = ?`
	_, err := db.Exec(query, articleLink)
	return err
}

func FetchToReadArticle(db *sql.DB) ([]rss.Article, error) {
	query := `SELECT title, link, date, to_read FROM article WHERE to_read = TRUE`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []rss.Article
	for rows.Next() {
		var article rss.Article
		if err := rows.Scan(&article.Title, &article.Link, &article.Date, &article.ToRead); err != nil {
			return nil, err
		}
		fmt.Println("here",article)
		articles = append(articles, article)
	}

	return articles, nil
}

func MarkAsRead(db *sql.DB, articleLink string) error {
	query := `UPDATE article SET to_read = FALSE WHERE link = ?`
	_, err := db.Exec(query, articleLink)
	return err
}
