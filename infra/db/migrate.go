import (
    "fmt"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
    migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB(db *sqlx.DB, dir string) error {
    migrations := &migrate.FileMigrationSource{
        Dir: dir,
    }

    n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
    if err != nil {
        return err
    }

    fmt.Printf("Successfully applied %d migrations\n", n)
    return nil
}