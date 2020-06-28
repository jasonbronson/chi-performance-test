package main

import (
	 "github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *Database/sql

func init(){
	//dbconn := "postgresql://doadmin:l8visgg5vd5p0fpu@db-postgresql-nyc3-88586-do-user-7602630-0.a.db.ondigitalocean.com:25060/defaultdb?sslmode=require"
	DB, err = gorm.Open( "postgres", "host=db-postgresql-nyc3-88586-do-user-7602630-0.a.db.ondigitalocean.com port=25060 user=doadmin dbname=defaultdb sslmode=disable password=l8visgg5vd5p0fpu")
    if err != nil {
     panic("failed to connect database")
    }
    defer db.Close()
}