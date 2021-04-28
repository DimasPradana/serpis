package config

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
	log "github.com/sirupsen/logrus"
)

func KonekOracle(username, pwd, host, port, sn string) (*sql.DB, error) {
	connString := fmt.Sprintf("%v/%v@%v:%v/%v", username, pwd, host, port, sn)
	db, err := sql.Open("godror", connString)
	// if err != nil {
	if err = db.Ping(); err != nil {
		log.Fatalf("koneksi ✗, dengan error : %v", err)
		db.Close()
	} else {
		log.Tracef("koneksi ✓")
	}
	return db, nil
}

func KonekMysql(){}
func KonekPostgre(){}
