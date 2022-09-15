package mssqldsn

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	mssql "github.com/denisenkom/go-mssqldb"
)

func Test(t *testing.T) {
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
	}
}
