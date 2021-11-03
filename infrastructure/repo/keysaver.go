package repo

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

func (k *keySaver) Get(stringChanged string) (string, error) {
	selectScript := "SELECT \"longUrl\" FROM \"urls\" WHERE \"Key\" =$1"
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
	return "", err
}

func (k *keySaver) Save(url domain.LongURL, key string) error {
	k.mutex.RLock()
	defer k.mutex.RUnlock()
	//k.repo123.Repo[key] = url.LongURLData
	//k.mutex.RUnlock()

	//psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//
	//db, err := sql.Open("postgres", psqlconn)
	//CheckError(err)
	//
	//defer db.Close()
	//
	//var urlToCut UrlToCut
	//err = ctx.Bind(&urlToCut)
	//if err != nil {
	//	log.Print("Bind returned error")
	//	return ctx.String(http.StatusBadRequest, err.Error())
	//}
	//
	//key, err := u.service.MakeKey(urlToCut.LongUrl)
	//
	//// dynamic
	//insertScript := `insert into "urls"("longUrl", "shortUrl") values($1, $2)`
	//_, err = db.Exec(insertScript, urlToCut.LongUrl, key)
	//CheckError(err)

	insertScript := `insert into "urls"("longUrl", "shortUrl") values($1, $2)`
	_, err := k.repository.Exec(insertScript, url.LongURLData, key)
	if err != nil {
		log.Println("Insert execute error:", err.Error())
	}

	return nil
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "klklkl"
	dbname   = "URLchanger"
)
