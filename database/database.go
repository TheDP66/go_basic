package main

import "fmt"

var connection string

func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	fmt.Println(connection)
	return connection
}
