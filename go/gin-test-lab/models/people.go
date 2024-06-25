package models

import (
	"context"
	"fmt"
	"log"

	"github.com/iancoleman/strcase"
	"github.com/jackc/pgx/v5/pgxpool"
)

type People struct {
	Id        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Age       int32  `db:"age"`
}

// Create a custom PeopleModel type which wraps the sql.DB connection pool.
type PeopleModel struct {
	DB *pgxpool.Pool
}

// Use a method on the custom PeopleModel type to run the SQL query.
func (m PeopleModel) All() ([]any, error) {
	ctx := context.Background()
	columnNames, err := m.getColumnNames(ctx)
	if err != nil {
		return nil, err
	}

	for _, cn := range columnNames {
		log.Printf("column name: %s", cn)
	}
	rows, err := m.DB.Query(ctx, "SELECT * FROM people")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ppls []any

	for rows.Next() {
		// var ppl map[string]interface{}

		// for i := range 

		p, err := rows.Values()
		if err != nil {
			return nil, err
		}

		fmt.Printf("Printing: %v", p)

	// 	err = rows.Scan(&ppl)
	// 	if err != nil {
	// 		return nil, err
	// 	}

		ppls = append(ppls, p)
	// }
	// if err = rows.Err(); err != nil {
	// 	return nil, err
	}

	// return ppls, nil

	return ppls,nil
}

func (m PeopleModel) getColumnNames (ctx context.Context) ( []string, error) {
	sql := `SELECT column_name 
	FROM information_schema.columns 
	WHERE table_schema = 'public' 
	AND table_name = 'people';`

	rows, err := m.DB.Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res [][]string
	for rows.Next() {
		var h []string

		val, err := rows.Values()
		if err != nil {
			return nil, err
		}

		for i, v := range val {
			h[i] = strcase.ToLowerCamel(v.(string))
		}

		res = append(res, h)
	}

	return res[0], nil
}

// Use a method on the custom PeopleModel type to run the SQL query.
func (m PeopleModel) Count() (any, error) {
	ctx := context.Background()
	rows, err := m.DB.Query(ctx, "")
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var rs []any

	for rows.Next() {
		var r RowCount

		v, err := rows.Values()
		if err != nil {
			return 0, err
		}

		fmt.Printf("Printing: %v", v)

		err = rows.Scan(&r.Count)
		if err != nil {
			return 0, err
		}

		rs = append(rs, v[0])
	}
	if err = rows.Err(); err != nil {
		return 0, err
	}

	return rs[0], nil
}
