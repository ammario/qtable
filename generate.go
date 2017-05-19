package qtable

import (
	"bytes"
	"database/sql"
	"fmt"

	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
)

// Generate accepts rows and generates a table.
// It accepts an error as it's second parameter so it can be composed with sql.DB.Query
// It closes rows.
func Generate(rows *sql.Rows, e error) (table []byte, err error) {
	if e != nil {
		return nil, e
	}
	buf := &bytes.Buffer{}
	err = GenerateCustom(rows, tablewriter.NewWriter(buf))
	return buf.Bytes(), err
}

// GenerateCustom accepts rows and an instantiated table writer
// wr hould not already have it's header set.
// It closes rows.
func GenerateCustom(rows *sql.Rows, wr *tablewriter.Table) error {
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return errors.Wrap(err, "failed to get columns")
	}
	wr.SetHeader(cols)

	//rows.Scan only wants pointers
	row := make([]interface{}, len(cols))
	rowPts := make([]interface{}, len(cols))
	for i := range rowPts {
		rowPts[i] = &row[i]
	}

	strs := make([]string, len(cols))

	for rows.Next() {
		if err := rows.Scan(rowPts...); err != nil {
			return errors.Wrap(err, "failed to scan row")
		}

		//get string representation of row values
		for i := range row {
			switch row[i].(type) { //TODO.. more conditions
			case []byte:
				strs[i] = string(row[i].([]byte))
			default:
				strs[i] = fmt.Sprintf("%v", row[i])
			}
		}

		wr.Append(strs)
	}
	wr.Render()
	return nil
}
