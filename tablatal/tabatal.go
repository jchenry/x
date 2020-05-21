package tabatal

import (
	"encoding/csv"
	"io"
	"strings"
)

func Parse(data string) (a []map[string]interface{}, err error) {
	// a := []map[string]interface{}{}
	r := csv.NewReader(strings.NewReader(data))
	r.Comma = '\t'
	r.Comment = ';'
	var header []string
	for i := 0; ; i++ {
		row, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return a, err
		}
		if i == 0 {
			header = row
			continue
		}

		record := map[string]interface{}{}
		for i, h := range header {
			record[h] = row[i]
		}
		a = append(a, record)
	}
	return a, nil
}
