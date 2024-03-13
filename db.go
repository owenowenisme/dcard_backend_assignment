package main

import (
	"database/sql"
	"os"
	"fmt"
	"log"
	_ "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)
func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}
func connectDB() *sql.DB {
	host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname))
	if err != nil {
		panic(err)
	}
	return db
}
func createAds(ad Ad) {
	db := connectDB()
	defer db.Close()

	sqlStatement := `
	INSERT INTO ads (title, start_at, end_at, age_start, age_end, country, platform, gender)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, ad.Title, ad.StartAt, ad.EndAt, ad.Conditions.AgeStart, ad.Conditions.AgeEnd, pq.Array(ad.Conditions.Country), pq.Array(ad.Conditions.Platform),ad.Conditions.Gender).Scan(&id)
	if err != nil {
		panic(err)
	}
}
func retrieveAds(q QueryCondition) ([]map[string]interface{}, error){
	db := connectDB()
	defer db.Close()

	sqlStatement := "SELECT * FROM ads WHERE NOW() BETWEEN start_at AND end_at"
    var args []interface{}

	i := 1
	if q.Age != 0 {
		sqlStatement += fmt.Sprintf(" AND (age_start <= $%d AND age_end >= $%d OR age IS NULL)", i, i)
		args = append(args, q.Age)
		i++
	}
	if q.Gender != "" {
		sqlStatement += fmt.Sprintf(" AND (gender = $%d OR gender IS NULL)", i)
		args = append(args, q.Gender)
		i++
	}
	if q.Country != "" {
		sqlStatement += fmt.Sprintf(" AND ($%d = ANY(country) OR country IS NULL)", i)
		args = append(args, q.Country)
		i++
	}
	if q.Platform != "" {
		sqlStatement += fmt.Sprintf(" AND ($%d = ANY(platform) OR platform IS NULL)", i)
		args = append(args, q.Platform)
		i++
	}

	sqlStatement +=(" ORDER BY end_at ASC")
	if q.Offset != 0 {
		sqlStatement += fmt.Sprintf(" OFFSET $%d", i)
		args = append(args, q.Offset)
		i++
	}
	if q.Limit == 0 {
		q.Limit = 5
	}
	sqlStatement += fmt.Sprintf(" LIMIT $%d",i)
	args = append(args, q.Limit)
    rows, err := db.Query(sqlStatement, args...)
	log.Println(sqlStatement,args)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	columns,_ := rows.Columns()
	
	values := make([]interface{}, len(columns)) 
	pointers := make([]interface{}, len(columns)) 
	var result []map[string]interface{}
	for rows.Next() {
		for i := range values {
			pointers[i] = &values[i]
		}

		if err := rows.Scan(pointers...); err != nil {
			log.Fatal(err)
		}

		row := make(map[string]interface{})
		for i, colName := range columns {
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				row[colName] = string(b)
			} else {
				row[colName] = val
			}
		}
		result = append(result, row)
	}
	return result,nil
}
func getNOW() string{
	db := connectDB()
	defer db.Close()
	var now string
	err := db.QueryRow("SELECT NOW()").Scan(&now)
	if err != nil {
		panic(err)
	}
	return now
}
