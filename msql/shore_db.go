package msql

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func RowsToJSON(rows *sql.Rows) {

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values

	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

type VMDBQueries struct {
}
type Tenants struct {
}

type RingTone struct {
	IntFileName string `gorm:"column:IntFileName"`
	ExtFileName string `gorm:"column:ExtFileName"`
	RingToneID  int    `gorm:"type:int(11);primary_key;column:RingToneID"`
	UserGroupID int    `gorm:"type:int(11);primary_key;column:UserGroupID"`
}

func (r RingTone) TableName() string {
	return "ugringtones"
}

func populateRedis(rawQuery string, db *gorm.DB) {

	rows, err := db.DB().Query(rawQuery)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	data := Jsonify(rows)
	fmt.Println("data:= [\n", data, "]\n")

}

/* func getUsers(db *gorm.DB) {
	// mohFiles := "SELECT DISTINCT IntFileName, ExtFileName FROM ugringtones ORDER BY IntFileName"

	rows, err := db.Table("ugringtones").Select("IntFileName, ExtFileName").Rows()
	columns, _ := rows.Columns()

	if err != nil {
		log.Fatalf("%s", err)
	}
	defer rows.Close()

	ringTones := []RingTone{}
	db.Debug().Find(&ringTones)

	for _, r := range ringTones {
		fmt.Printf("\n%v\n", r)
	}


	var extFileName string

	for rows.Next() {

		rows.Scan(&initFile, &extFileName)
		log.Println("InitFIle=", initFile, ", extFile=", extFileName)
	}


 }*/

// func main() {
// 	// id:password@tcp(your-amazonaws-uri.com:3306)/dbname
// 	dbURI := constants.MySqlUser + ":" + constants.MySqlPassword + "@tcp(" +
// 		constants.MySqlDbServer + ":" + constants.MySqlPort + ")/" + constants.MySqlDbName + "?charset=utf8&parseTime=True&loc=Local"

// 	fmt.Println("dbURI:[", dbURI, "]")

// 	db, err := gorm.Open(constants.MySqlString, dbURI)
// 	if err != nil {
// 		log.Fatalf("%s", err)
// 	}
// 	defer db.Close()

// 	populateRedis("SELECT * FROM tenants", db)
// 	populateRedis("SELECT DISTINCT IntFileName, ExtFileName FROM ugringtones ORDER BY IntFileName", db)
// 	// getUsers(db)
// }
