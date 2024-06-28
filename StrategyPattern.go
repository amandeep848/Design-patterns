package main

import "fmt"

type IDBconnection interface { // struct can be assigned any type in runtime
	Connect()
}

type DBConnection struct {
	Db IDBconnection // Compatible to accept any type in runtime
}

func (con DBConnection) DBConnect() { // Receiver Function for struct DBConnection
	con.Db.Connect()
}

type MySqlConnection struct {
	ConnectionString string
}

func (con MySqlConnection) Connect() {
	fmt.Println(("MySql " + con.ConnectionString))
}

type PostgressConnection struct {
	ConnectionString string
}

func (con PostgressConnection) Connect() {
	fmt.Println("Postgress " + con.ConnectionString)
}

func main() {

	MySqlConnection := MySqlConnection{ConnectionString: "MySql DB is connected"}
	con := DBConnection{Db: MySqlConnection}
	con.DBConnect()

	pgConnection := PostgressConnection{ConnectionString: "Postgress DB is connected"}
	con = DBConnection{Db: pgConnection}
	con.DBConnect()

}

// Behavioral design that is used when we have multiple algorithm for a specific task and client decide the actual implementation to be used at runtime
