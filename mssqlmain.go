// Main demo package for create DSN and copy data between tables of MSSQL DBMS.
// Основной демо пакет для создания DSN и копирования данных между таблицами СУБД MSSQL.

package main

import (
	"fmt"
	"log"

	"time"

	"github.com/blablatov/sqlbulk2dirx/mssqldsn"
	"github.com/blablatov/sqlbulk2dirx/sqlinsertrs"
)

// Demo constant request of data insert. One can forming dynamic request.
// Демо константа insert запроса. Запрос можно формировать динамически.
const (
	insertIntegrTableSql = "INSERT DirectumRX.dbo.dBase SELECT Id, Discriminator, Name, BusinessUnit, ItemNumber, ItemName FROM DirectumRX.dbo.dBaseTest WHERE BusinessUnit = 65;"
)

func main() {
	start := time.Now()
	// Structure of DSN. Структура DSN.
	dd := mssqldsn.DataDsn{
		Debug:    true,
		User:     "user",
		Password: "password",
		Port:     1433,
		Server:   "rx-db-directum",
		Database: "DirectumRX",
	}
	// Calling interface method to form DSN and create connect to DBMS.
	// Вызов метода интерфейса, для формирования DSN и создания подключения к СУБД.
	var d mssqldsn.ConDsner = dd
	db := d.SqlConDsn()
	defer db.Close()

	cs := make(chan string) // Channel of function mssql-request. Канал функции mssql-запроса.
	go sqlinsertrs.SqlInserTrs(insertIntegrTableSql, db, cs)
	// Getting data from channel of goroutine. Получение данных запроса из канала горутины.
	log.Println("\nResult of sql-request: ", <-cs)
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs Request execution time\n", secs)
}
