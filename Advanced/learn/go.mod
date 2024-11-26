module example.com/learn

go 1.23.3

require (
	cryptit v0.0.0-00010101000000-000000000000
	github.com/pborman/uuid v1.2.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/google/uuid v1.0.0 // indirect
)

replace cryptit => ../cryptit
