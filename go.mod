module github.com/blablatov/mssql2sql

go 1.16

require (
	github.com/blablatov/sqlbulk2dirx/mssqldsn v0.0.0-00010101000000-000000000000 // indirect
	github.com/blablatov/sqlbulk2dirx/sqlinsertrs v0.0.0-00010101000000-000000000000 // indirect
	github.com/denisenkom/go-mssqldb v0.12.2
)

replace github.com/blablatov/sqlbulk2dirx/sqlinsertrs => ./sqlinsertrs

replace github.com/blablatov/sqlbulk2dirx/mssqldsn => ./mssqldsn
