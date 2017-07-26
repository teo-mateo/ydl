package ydata

import (
	"database/sql"

	ydlconf "github.com/teo-mateo/ydl/config"

	"fmt"
	//I do this because I want this. fuck your motherfucking warnings.
	_ "github.com/lib/pq"
)

const (
	STATUSQueued1      = 1
	STATUSDownloading2 = 2
	STATUSDownloaded3  = 3
	STATUSError4       = 4
	STATUSSkipped5     = 5
)

func connect() (db *sql.DB, err error) {
	psqlInfo := ydlconf.PgConnectionString()
	db, err = sql.Open("postgres", psqlInfo)
	return
}

// YQueueAdd adds a new video to the queue (db layer)
func YQueueAdd(url string, who string) (int, error) {
	db, err := connect()
	if err != nil {
		return -1, err
	}
	defer db.Close()
	var id int
	err = db.QueryRow("SELECT yqueue_ins($1, $2)", url, who).Scan(&id)
	if err != nil {
		return -1, err
	}

	//no error
	return id, nil
}

// YQueueUpdate ...
func YQueueUpdate(id int, status int, file ...string) error {
	db, err := connect()
	if err != nil {
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

// YQueueGet ...
func YQueueGet(id int) (YQueue, error) {
	db, err := connect()
	if err != nil {
		return YQueue{}, err
	}
	defer db.Close()

	var result = YQueue{}
	err = db.QueryRow("select * from yqueue_get($1)", id).
		Scan(&result.ID, &result.YTUrl, &result.Status, &result.File, &result.Lastupdate, &result.Who)

	if err != nil {
		return YQueue{}, err
	}

	return result, nil
}

//YQueueGetAll ...
func YQueueGetAll(status int, who string) ([]YQueue, error) {
	db, err := connect()
	if err != nil {
		return []YQueue{}, nil
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("select * from yqueue_get(null, %d, '%s')", status, who))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]YQueue, 0)
	for rows.Next() {
		yq := YQueue{}
		err := rows.Scan(&yq.ID, &yq.YTUrl, &yq.Status, &yq.File, &yq.Lastupdate, &yq.Who)
		if err != nil {
			return nil, err
		}
		result = append(result, yq)
	}

	return result, nil
}

//YQueueDelete ...
func YQueueDelete(id int) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("select yqueue_del(%d)", id))
	if err != nil {
		return err
	}

	return nil
}

// YQueue ...
type YQueue struct {
	ID         int
	YTUrl      string
	Status     int
	File       sql.NullString
	Lastupdate sql.NullString
	Who        sql.NullString
}
