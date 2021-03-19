package mahasiswa

import (
	"database/sql"
	"fmt"
	"log"
	"pendalaman-res-api/models"
	"time"
)

const table = "mahasiswa"
const datetime_layout = "2006-01-02 15:04:05"

func GetAll(db *sql.DB) (mahasiswa []models.Mahasiswa, err error) {

	sqlTxt := fmt.Sprintf("SELECT * FROM %s ORDER BY id DESC", table)

	rowQuery, err := db.Query(sqlTxt)

	if err != nil {
		log.Fatal(err)
		defer rowQuery.Close()
	}

	var mahasiswas []models.Mahasiswa

	for rowQuery.Next() {
		mhs := models.Mahasiswa{}
		createdAt, updatedAt := "", ""

		err = rowQuery.Scan(&mhs.ID,
			&mhs.NIM,
			&mhs.Nama,
			&mhs.Semester,
			&createdAt,
			&updatedAt)

		mhs.CreatedAt, _ = time.Parse(datetime_layout, createdAt)
		mhs.UpdatedAt, _ = time.Parse(datetime_layout, updatedAt)

		if err != nil {
			return nil, err
		}
		mahasiswas = append(mahasiswas, mhs)
	}
	return
}

func CreateMhs(db *sql.DB, mhs *models.Mahasiswa) (err error) {
	sqlTxt := fmt.Sprintf("INSERT INTO %v (nim, nama, semester, created_at, updated_at) VALUES(?,?,?,?,?)", table)
	timeNow := time.Now()
	resQuery, err := db.Exec(sqlTxt, mhs.NIM, mhs.Nama, mhs.Semester, timeNow, timeNow)

	if err != nil {
		return err
	}

	lastId, err := resQuery.LastInsertId()
	if err != nil {
		return err
	}
	mhs.ID = int(lastId)
	mhs.CreatedAt = timeNow
	mhs.UpdatedAt = timeNow
	return

}
