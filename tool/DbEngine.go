package tool

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DBTool struct{
	*sql.DB
}

func DbEngine()(db *sql.DB,err error){
	cfg,_ := ParseConfig("./config/app.json")
	database := cfg.Db
	connStr := database.User+":"+database.Password+"@tcp("+database.Host+":"+database.Port+")/"+database.Db_name
	db,err = sql.Open("mysql",connStr)
	if err != nil{
		log.Fatalf(err.Error())
		return nil,err
	}
	return db,nil
}


