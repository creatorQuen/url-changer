package repository

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
	"sync"
	"url-changer/domain"
)

type keySaver struct {
	mutex      sync.RWMutex
	repository *sql.DB
}

func NewKeySaver(repo *sql.DB) *keySaver {
	return &keySaver{repository: repo}
}

func (k keySaver) GetFullString(stringChanged string) (string, error) {
	selectScript := "SELECT \"longUrl\" FROM \"ulrs\" WHERE \"shortUrl\" =$1"
	variable, err := k.repository.Query(selectScript, stringChanged)
	defer variable.Close()

	if err != nil {
		return "Select query error:", errors.New("no such url")
	}

	var urlLong string
	variable.Next()
	err = variable.Scan(&urlLong)
	if err != nil {
		log.Println("Scan error:", err.Error())
	}
	return urlLong, nil
}

func (k *keySaver) Save(url domain.LongURL, key string) error {
	k.mutex.RLock()
	defer k.mutex.RUnlock()

	insertScript := `insert into "ulrs"("longUrl", "shortUrl") values($1, $2)`
	_, err := k.repository.Exec(insertScript, url.LongURLData, key)
	if err != nil {
		log.Println("Insert execute error:", err.Error())
	}

	return nil
}
