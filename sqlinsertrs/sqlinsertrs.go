// Demo module for copying data between tables per transaction.
// Демонстрационный модуль для копирования данных между таблицами за транзакцию.

package sqlinsertrs

import (
	"database/sql"
	"log"
)

func SqlInserTrs(insertIntegrTableSql string, db *sql.DB, cs chan string) {
	// Transaction start. Начало транзакции.
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
	cs <- "Data written success to DB." // Sending result to main module. Передача результата в main-программу.
}
