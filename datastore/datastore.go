package datastore

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

// Database type
type Type string

const (
	TypeMySQL Type = "MySQL"
)

// Contains db wrappers and Info struct
var (
	// SQL wrapper
	SQL *sqlx.DB
	// Database info
	databases Info
)

// Contains database type and specific info structs
type Info struct {
	Type  Type
	MySQL MySQLInfo
}

// Contains mysql info
type MySQLInfo struct {
	Hostname  string `json:"Hostname"`
	Name      string `json:"Name"`
	Username  string `json:"Username"`
	Password  string `json:"Password"`
	Port      string `json:"Port"`
	Parameter string `json:"Parameter"`
	Flow      *sql.DB
}

// Returns Data Source Name
func getDSN(ci MySQLInfo) string {
	// Example: root:@tcp(localhost:3306)/test
	return ci.Username +
		":" +
		ci.Password +
		"@tcp(" +
		ci.Hostname +
		":" +
		ci.Port +
		")/" +
		ci.Name + ci.Parameter
}

// Connects to the database
func Connect(d Info) {
	var err error

	databases = d

	switch d.Type {
	case TypeMySQL:
		if SQL, err = sqlx.Connect("mysql", getDSN(d.MySQL)); err != nil {
			log.Fatalf("sqlx.Open failed : %v", err)
		}

		if err = SQL.Ping(); err != nil {
			log.Fatalf("sqlx.Ping failed : %v", err)
		}
	default:
		log.Fatalf("No registred database in config")
	}
}

func ReadConfig() Info {
	return databases
}
