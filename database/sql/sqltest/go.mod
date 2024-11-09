module github.com/pamburus/go-mod/database/sql/sqltest

go 1.23.2

require (
	github.com/lib/pq v1.10.9
	github.com/pamburus/go-mod/testing/mocks v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/pamburus/go-mod/testing/mocks => ../../../testing/mocks
