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

//menghandel pengambilan data di databases
func GetAll(db *sql.DB) (mahasiswa []models.Mahasiswa, err error) {

	// var mahasiswas []models.Mahasiswa
	sqlTxt := fmt.Sprintf("SELECT * FROM %v ORDER BY id DESC", table)

	rowQuery, err := db.Query(sqlTxt)

	if err != nil {
		log.Fatal(err)
		defer rowQuery.Close()
	}

	for rowQuery.Next() {
		mhs := models.Mahasiswa{}
		createdAt, updatedAt := "", ""

		if err = rowQuery.Scan(&mhs.ID,
			&mhs.NIM,
			&mhs.Nama,
			&mhs.Semester,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		mhs.CreatedAt, _ = time.Parse(datetime_layout, createdAt)
		mhs.UpdatedAt, _ = time.Parse(datetime_layout, updatedAt)

		mahasiswa = append(mahasiswa, mhs)
	}
	return mahasiswa, nil
}

//menghanel pembuatan data mahasiswa
func CreateMhs(db *sql.DB, mhs *models.Mahasiswa) (err error) {
	sqlTxt := fmt.Sprintf("INSERT INTO %v (nim, nama, semester, created_at, updated_at) VALUES(?,?,?,?,?)", table)
	timeNow := time.Now()
	resQuery, err := db.Exec(sqlTxt, mhs.NIM, mhs.Nama, mhs.Semester, timeNow, timeNow)

	if err != nil {
		return err
	}

	lastId, err := resQuery.LastInsertId()
	if err != nil {
		fmt.Println("error pada" + err.Error())
		return err
	}
	mhs.ID = int(lastId)
	mhs.CreatedAt = timeNow
	mhs.UpdatedAt = timeNow

	return nil

}

//menghandel untuk perubahan data
func UpdateMhs(db *sql.DB, mhs *models.Mahasiswa) (err error) {

	sqlTxt := fmt.Sprintf("UPDATE %v SET nim=?, nama=?, semester=?,created_at=?, updated_at=? WHERE id=?", table)
	timeNow := time.Now()

	_, err = db.Exec(sqlTxt, mhs.NIM, mhs.Nama, mhs.Semester, timeNow, timeNow, mhs.ID)

	if err != nil {
		fmt.Println("Error pada " + err.Error())
		return err
	}
	return nil
}

//menghandel untuk menghapus data di databases
func DeleteMhs(db *sql.DB, mhs *models.Mahasiswa) (err error) {
	sqlTxt := fmt.Sprintf("DELETE FROM %v WHERE id = %d", table, mhs.ID)
	resQuery, err := db.Exec(sqlTxt)
	if err != nil {
		fmt.Println("error pada " + err.Error())
		return err
	}
	_, err = resQuery.RowsAffected()

	return nil
}
