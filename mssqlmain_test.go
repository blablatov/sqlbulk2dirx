package main

import (
	"fmt"
	"log"
	"sync"
	"testing"

	"github.com/blablatov/sqlbulk2dirx/mssqldsn"
	"github.com/blablatov/sqlbulk2dirx/sqlinsertrs"
)

const (
	insertIntegrTableSql = "INSERT DirectumRX.dbo.dBase SELECT Id, Discriminator, Name, BusinessUnit, ItemNumber, ItemName FROM DirectumRX.dbo.dBaseTest WHERE BusinessUnit = 65;"
)

func Test(t *testing.T) {
	var tests = []struct {
		user       string
		password   string
		server     string
		database   string
		connString string
		port       int
	}{
		{"admin", "qwerty", "google.com", "dbase", "mystring", 1999},
		{".", "\t", ",", "?", "!!", 10001},
		{"\t", "NaN\null\n\n", "NaN\null\n", "\ntrue\n", "\t\n", 23409},
		{"Data for test", "Number 99999 to data test", ")([]$%", "Number 1234567890000", "\n&*@#", 9876},
		{"Yes, no", "No, or, yes, _, ops", "NULL, true", "false, or, yes", "-, ops", 65535},
	}

	var prevuser string
	for _, test := range tests {
		if test.user != prevuser {
			fmt.Printf("\n%s\n", test.user)
			prevuser = test.user
		}
	}

	var prevpassword string
	for _, test := range tests {
		if test.password != prevpassword {
			fmt.Printf("\n%s\n", test.password)
			prevpassword = test.password
		}
	}

	var prevserver string
	for _, test := range tests {
		if test.server != prevserver {
			fmt.Printf("\n%s\n", test.server)
			prevserver = test.server
		}
	}

	var prevdatabase string
	for _, test := range tests {
		if test.database != prevdatabase {
			fmt.Printf("\n%s\n", test.database)
			prevdatabase = test.database
		}
	}

	var prevconnString string
	for _, test := range tests {
		if test.connString != prevconnString {
			fmt.Printf("\n%s\n", test.connString)
			prevconnString = test.connString
		}
	}

	var prevport int
	for _, test := range tests {
		if test.port != prevport {
			fmt.Printf("\n%d\n", test.port)
			prevport = test.port
		}
	}
}

func BenchmarkInterfaceDsn(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < 5; i++ {
		dd := mssqldsn.DataDsn{
			Debug:    true,
			User:     "user",
			Password: "password",
			Port:     1433,
			Server:   "rx-db-directum",
			Database: "DirectumRX",
		}
		var d mssqldsn.ConDsner = dd
		db := d.SqlConDsn()
		defer db.Close()
		fmt.Println("Result via interface method dsn", db)
	}
}

func BenchmarkGoroutineDsn(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < 5; i++ {
		dd := mssqldsn.DataDsn{
			Debug:    true,
			User:     "user",
			Password: "password",
			Port:     1433,
			Server:   "rx-db-directum",
			Database: "DirectumRX",
		}
		var d mssqldsn.ConDsner = dd
		db := d.SqlConDsn()
		defer db.Close()

		cs := make(chan string) // Channel of function mssql-request. Канал функции mssql-запроса.
		var wg sync.WaitGroup
		go sqlinsertrs.SqlInserTrs(insertIntegrTableSql, db, cs)
		// Getting data from goroutine. Получение данных из канала горутины.
		log.Println("\nResult of sql-request: ", <-cs)
		go func() {
			wg.Wait()
			close(cs)
		}()
	}
}
