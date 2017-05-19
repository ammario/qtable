package sqlviz

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestGenerate(t *testing.T) {
	const query = `SELECT *
FROM (
    SELECT "hey" as name
    UNION SELECT "bob"
    UNION SELECT "job"
    UNION SELECT "cob"
    UNION SELECT "slob"
    UNION SELECT "dob"
)`

	db, err := sql.Open("sqlite3", "file::memory:?mode=memory&cache=shared")
	if err != nil {
		panic(err)
	}

	table, err := Generate(db.Query(query))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Table...\n%s", table)
}
