package sqlinsertrs

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	mssql "github.com/denisenkom/go-mssqldb"
)

const (
	insertIntegrTableSql = "INSERT DirectumRX.dbo.dBase SELECT Id, Discriminator, Name, BusinessUnit, ItemNumber, ItemName FROM DirectumRX.dbo.dBaseTest WHERE BusinessUnit = 65;"
)

func Test(t *testing.T) {
	var tested = []struct {
		insertIntegrTableSql string
	}{
		{"    :, , :"},
		{"wertyuittt_90/,_=="},
		{"__+_8_0000//,_+=;;;.?/-@"},
		{"QQVNMBVFRERTYuiop,asdfghU56789543211IJNM<>L><M???_"},
	}

	var previnsertIntegrTableSql string
	for _, test := range tested {
		if test.insertIntegrTableSql != previnsertIntegrTableSql {
			fmt.Printf("\n%s\n", test.insertIntegrTableSql)
			previnsertIntegrTableSql = test.insertIntegrTableSql
		}
	}

	var tests = []struct {
		User     string
		Password string
		Server   string
		Database string
		Debug    bool
		Port     int
	}{
		{"admin", "qwerty", "google.com", "dbase", true, 1999},
		{".", "\t", ",", "?", false, 10001},
		{"\t", "NaN\null\n\n", "NaN\null\n", "\ntrue\n", false, 23409},
		{"Data for test", "\n&*@#", ")([]$%", "Number 1234567890000", true, 9876},
		{"Yes, no", "No, or, yes, _, ops", "NULL, true", "false, or, yes", false, 60005},
	}

	var prevUser string
	for _, test := range tests {
		if test.User != prevUser {
			fmt.Printf("\n%s\n", test.User)
			prevUser = test.User
		}
	}

	var prevPassword string
	for _, test := range tests {
		if test.Password != prevPassword {
			fmt.Printf("\n%s\n", test.Password)
			prevPassword = test.Password
		}
	}

	var prevServer string
	for _, test := range tests {
		if test.Server != prevServer {
			fmt.Printf("\n%s\n", test.Server)
			prevServer = test.Server
		}
	}

	var prevDatabase string
	for _, test := range tests {
		if test.Database != prevDatabase {
			fmt.Printf("\n%s\n", test.Database)
			prevDatabase = test.Database
		}
	}

	var prevDebug bool
	for _, test := range tests {
		if test.Debug != prevDebug {
			fmt.Printf("\n%t\n", test.Debug)
			prevDebug = test.Debug
		}
	}

	var prevPort int
	for _, test := range tests {
		if test.Port != prevPort {
			fmt.Printf("\n%d\n", test.Port)
			prevPort = test.Port
		}
	}

	connString := fmt.Sprintf("server=%s; user id=%s; password=%s; port=%d; database=%s", prevServer, prevUser, prevPassword, prevPort, prevDatabase)
	if prevDebug {
		fmt.Printf("MSSQL connString:%s\n", connString)

		// Creating a new connector object. Создание объекта подключения.
		connector, err := mssql.NewConnector(connString)
		if err != nil {
			log.Println(err)
		}
		log.Println(connector)

		// Pass connector to sql.OpenDB to get a sql.DB object. Получение объекта sql.DB
		db := sql.OpenDB(connector)
		log.Println(db)

		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}
		defer tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.
		// Откат действий будет проигнорирован, если позже выполнение транзакции будет зафиксировано.

		stmt, err := tx.Prepare(insertIntegrTableSql)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
		// Выполненые операторы занимают ресурсы сервера и должны быть закрыты, после выполнения.

		if _, err := stmt.Exec("open source"); err != nil {
			log.Fatal(err)
		}

		if err := tx.Commit(); err != nil {
			log.Fatal(err)
		}
	}
}
