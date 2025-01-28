package mysql

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	SQL_URL     = "SQL_URL"     
	SQL_USER_DB = "SQL_USER_DB" 
)

func NewMysqlConnection(ctx context.Context) (*sql.DB, error) {
	sqlURI := os.Getenv(SQL_URL) 
	if sqlURI == "" {
		log.Println("Variável de ambiente SQL_URL não definida")
		return nil, sql.ErrNoRows
	}

	db, err := sql.Open("mysql", sqlURI)
	if err != nil {
		return nil, err
	}

	if err = db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}

	log.Println("Conexão com o MySQL estabelecida com sucesso")
	return db, nil
}
