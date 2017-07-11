package ydata

import (
	"database/sql"
	"fmt"

	ydlconf "github.com/teo-mateo/ydl/config"

	//I do this because I want this. fuck your motherfucking warnings.
	_ "github.com/lib/pq"
)

// YQueueAdd adds a new video to the queue (db layer)
func YQueueAdd(url string) (int, error) {
	psqlInfo := ydlconf.PgConnectionString()
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	defer db.Close()
	var id int
	err = db.QueryRow("SELECT yqueue_ins($1)", url).Scan(&id)

	//no error
	return id, nil
}

// YQueueUpdate ...
func YQueueUpdate(id int, status int, file ...string) error {
	psqlInfo := ydlconf.PgConnectionString()
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	var fn interface{}
	if len(file) == 0 {
		fn = nil
	} else {
		fn = file[0]
	}

	_, err = db.Exec("select yqueue_upd($1, $2, $3)", id, status, fn)
	if err != nil {
		return err
	}

	//no error
	return nil
}
