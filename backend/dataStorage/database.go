package dataStorage

import "database/sql"

func GetDB() *sql.DB {
	dbInstance, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/knowledge_checkup")

	if err != nil {
		panic(err.Error())
	}

	return dbInstance
}
