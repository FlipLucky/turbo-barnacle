package db

import (
	"database/sql"
	"os"
	"strings"

	"github.com/FlipLucky/turbo-barnacle/internal/elements"
	"github.com/go-sql-driver/mysql"
)

type DBClient struct {
	*sql.DB
}

func createConfig() mysql.Config {
	return mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		DBName: os.Getenv("MYSQL_DATABASE"),
		Addr:   "db:3306",
		Net:    "tcp",
	}
}

func OpenDb() (*DBClient, error) {
	cfg := createConfig()
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	return &DBClient{db}, nil
}

func (db *DBClient) TestConnection() error {
	return db.Ping()
}

func (db *DBClient) RunMigrations() {

}

type getQuery struct {
	Table   string
	Columns []string
}

func (q getQuery) createQuery() string {
	columns := "*"
	if len(q.Columns) > 0 {
		columns = strings.Join(q.Columns, ", ")
	}
	querySections := []string{
		"SELECT",
		columns,
		"FROM",
		q.Table,
	}

	return strings.Join(querySections, " ")
}

func (db *DBClient) Get(query getQuery) (*sql.Rows, error) {

	results, err := db.Query(query.createQuery())
	if err != nil {
		return nil, err
	}
	return results, nil
}

func GetPages(db *DBClient) ([]elements.Page, error) {
	pageCollection := []elements.Page{}
	getPageQuery := getQuery{Table: "pages", Columns: []string{}}
	results, err := db.Get(getPageQuery)
	if err != nil {
		return nil, err
	}
	for results.Next() {
		var page elements.Page
		err = results.Scan(&page.Id, &page.Title, &page.Slug, &page.Template)
		if err != nil {
			return nil, err
		}
		pageCollection = append(pageCollection, page)
	}
	return pageCollection, nil
}

// if err != nil {
// 	log.Fatal(err)
// }
// pageCollection := []elements.Page{}
//
// for results.Next() {
// 	var page elements.Page
//
// 	err = results.Scan(&page.Id, &page.Title, &page.Slug, &page.Template)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	pageCollection = append(pageCollection, page)
// }
