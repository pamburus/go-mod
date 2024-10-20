module github.com/pamburus/go-mod/database/sql/qxpgx

go 1.23.1

replace github.com/pamburus/go-mod/database/sql/qx => ../qx

replace github.com/pamburus/go-mod/database/sql/qb => ../qb

require (
	github.com/jackc/pgx/v5 v5.7.1
	github.com/pamburus/go-mod/database/sql/qb v0.0.0-00010101000000-000000000000
	github.com/pamburus/go-mod/database/sql/qx v0.0.0-00010101000000-000000000000
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/pamburus/go-mod/gi v0.4.0 // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/text v0.18.0 // indirect
)
