package database

import (
	"context"
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DbPool *pgxpool.Pool

// func init() {

// }

func GetCount(query string, conn pgx.Conn) (int, error) {
	ctx := context.Background()
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var rs []int

	for rows.Next() {
		v, err := rows.Values()
		if err != nil {
			return 0, err
		}

		rs = append(rs, v[0].(int))
	}
	if err = rows.Err(); err != nil {
		return 0, err
	}

	return rs[0], nil
}

func GetColumnNames (ctx context.Context, conn *pgx.Conn, schema string, table string) ( []string, error) {
	sql := fmt.Sprintf(`SELECT column_name 
	FROM information_schema.columns 
	WHERE table_schema = '%s' 
	AND table_name = '%s';`, schema, table)

	rows, err := conn.Query(ctx, sql)
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