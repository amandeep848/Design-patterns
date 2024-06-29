package main

import "fmt"

// Strategy Interface : struct can be assigned any type in runtime
type IDBconnection interface {
	Connect()
}

// Concrete Strategies
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

// Context
type DBConnection struct {
	Db IDBconnection
}

func (con DBConnection) DBConnect() { // Receiver Function for struct DBConnection
	con.Db.Connect()
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
// Context : Represents the entity that uses different strategies.
// Strategy Interface : An interface that defines the method signature that all concrete strategies must implement.
// Concrete Strategies : A set of structs/classes that implement the Strategy Interface.
