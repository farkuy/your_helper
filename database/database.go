package database

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"your_helper/internal/config"

	"github.com/jackc/pgx/v4"
)

type CheckedTable struct {
	name string
}

var tablesName = []CheckedTable{
	{name: "users"},
}

func Init(bdCfg config.BdConfig) (*pgx.Conn, error) {
	path := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", bdCfg.User, bdCfg.Password, bdCfg.Dbname, bdCfg.Sslmode)

	db, err := pgx.Connect(context.Background(), path)
	if err != nil {
		slog.Error("Ошибка подключения к бд")
		return nil, err
	}

	for _, query := range tablesName {
		var exists bool
		checkingTable := checkTable(query.name)

		err = db.QueryRow(context.Background(), checkingTable).Scan(&exists)
		if err != nil {
			log.Fatal("QueryRow failed: ", query.name, " ", err)
		}
	}

	return db, nil
}

func checkTable(tableName string) string {
	return fmt.Sprintf(`
        SELECT EXISTS (
            SELECT 1
            FROM information_schema.tables
            WHERE table_schema = 'public'
              AND table_name = '%s'
        )
    `, tableName)
}
