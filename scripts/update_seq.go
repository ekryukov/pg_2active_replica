package scripts

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

func main() {
	seqName := os.Args[1]
	seqShift, _ := strconv.Atoi(os.Args[2])
	const dbhost = "localhost"
	const dbuser = "svip"
	const dbpass = "svip1"
	const dbname = "svip_db"
	const dbschema = "bpc01"

	connStr := fmt.Sprintf("host=%s dbname=%s user=%s password='%s' sslmode=disable search_path=%s", dbhost, dbname, dbuser, dbpass, dbschema)
	db, err := sql.Open("postgres", connStr)
	checkErr(err)
	globSeq := fmt.Sprintf(`alter sequence %s using bdr`, seqName)
	query := fmt.Sprintf(`select nextval('%s')`, seqName)
	log.Print(query)
	_, err = db.Exec(globSeq)
	for i := 0; i <= seqShift; i++ {

		_, err := db.Exec(query)
		checkErr(err)

	}
	defer db.Close()

}
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
