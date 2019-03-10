package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "abc123"
	dbName := "hortaConecta"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func ConsultAdvertisement(id int64) Advertisements {
	db := dbConn()
	qryConsult, err := db.Query("SELECT "+
		"id, "+
		"title, "+
		"description, "+
		"idAdvertiser, "+
		"advertisementState "+
		"FROM "+
		"advertisement "+
		"WHERE "+
		"id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	advertisement := Advertisement{}
	advertisements := []Advertisement{}
	for qryConsult.Next() {
		err = qryConsult.Scan(&advertisement.Id, &advertisement.Title, &advertisement.Description, &advertisement.IdAdvertiser, &advertisement.State)

		if err != nil {
			panic(err.Error())
		}
		advertisements = append(advertisements, advertisement)
	}
	defer db.Close()
	return advertisements

}

func ConsultAllAdvertisements() Advertisements {
	db := dbConn()
	qryConsult, err := db.Query("SELECT " +
		"id, " +
		"title, " +
		"description, " +
		"idAdvertiser, " +
		"advertisementState " +
		"FROM " +
		"advertisement")
	if err != nil {
		panic(err.Error())
	}

	advertisement := Advertisement{}
	advertisements := []Advertisement{}
	for qryConsult.Next() {
		err = qryConsult.Scan(&advertisement.Id, &advertisement.Title, &advertisement.Description, &advertisement.IdAdvertiser, &advertisement.State)

		if err != nil {
			panic(err.Error())
		}
		advertisements = append(advertisements, advertisement)
	}
	defer db.Close()
	return advertisements

}

func InsertAdvertisement(ad Advertisement) Advertisement {
	//TODO: Do error
	db := dbConn()
	qryInsert, err := db.Exec(
		"INSERT INTO advertisement( "+
			"	title, "+
			"	description, "+
			"	idAdvertiser, "+
			"	advertisementState "+
			") VALUES( "+
			"	?, "+
			"	?, "+
			"	?, "+
			"	? "+
			") ", ad.Title, ad.Description, ad.IdAdvertiser, ad.State)
	if err != nil {
		panic(err.Error())
	}

	id, err := qryInsert.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	adReturn := Advertisement{}
	adReturn.Id = id
	adReturn.State = ad.State
	adReturn.Title = ad.Title
	adReturn.Description = ad.Description
	adReturn.IdAdvertiser = ad.IdAdvertiser

	defer db.Close()
	return adReturn

}
