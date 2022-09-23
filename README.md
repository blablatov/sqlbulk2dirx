### sqlbulk2dirx
### Ru

Демо пакеты модуля на Go, для копирования данных (sql bulk copy) между таблицами СУБД MSSQL.  
Дополнительный код на C# для подключения Go-модуля как сторонней библиотеки в Directum RX. 

Для выполнения, запустить модуль:  
 
	mssqlmain  

Если параметры подключения DSN верны, проверить появление данных в реципиент-таблице.

***Схема обмена данными между Go-пакетами (scheme exchange of data between Go-packets):***
			
```mermaid
graph TB

  SubGraph1 --> SubGraph1Flow
  subgraph "Packet DSN"
  SubGraph1Flow(Create DSN)
  end
  
  SubGraph3 --> SubGraph2Flow
  subgraph "DBMS MSSQL"
  SubGraph2Flow(Tables of data MSSQL)
  end

  subgraph "Module MSSQL"
  Node1[Packet write to MSSQL `mssqlmain`] --> SubGraph1[Request to create DSN `mssqldsn`]
  SubGraph1Flow -- Response with DSN data --> Node1 
  Node1 --> SubGraph3[BulkCopy method of goroutine packet `sqlinsertrs`]
end
```	  

Демо код DS, для вызова сторонней либры из Directum RX. Demo code to call outside library from Directum RX:    

```C#
public vittual void ProcessSqlStart()
{
  var role = Roles.GetAll(r => Equals(r.Name, "Администраторы справочников")).FirstOrDefault();
  if (Users.Current.IncludedIn(role))
  {
    GoProcessSql.MyProcess.Main(); 
    Dialogs.ShowMessage("Копирование справочника SAP выполнено");
  }
  else
  {
    Dialogs.ShowMessage("У пользователя нет прав на копирование справочника");
  }
}
```
В СУБД Directum RX должна быть таблица-справочник: "Администраторы справочников".


### En

It's demo packages on Go for copying data (sql bulk copy) between DBMS tables of MSSQL.   
Additional C# code for connecting the Go-module as outside library in Directum RX.   

For execute, run module:  

	mssqlmain

If DSN connection parameters are correct, check if copyed the data in to recipient-table.    
The Directum RX DBMS should have a reference table: "Directory administrators".

***Общая схема обмена (general exchange scheme):***

```mermaid
graph TB

  SubGraph1 --> SubGraph1Flow
  subgraph "MSSQL"
  SubGraph1Flow(DBMS)
  SubGraph1Flow -- Select --> Donor
  SubGraph1Flow -- Insert --> Recipient
  end

  subgraph "Directum RX"
  Node1[Function Directum RX] --> Node2[C# library `GoProcessSql.dll`]
  Node2 --> SubGraph1[Go module `mssqlmain`]
  SubGraph1 --> FinalThing[Show Message]
end
```

