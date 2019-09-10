// package main

// import (
// 	"fmt"
// 	_"github.com/weigj/go-odbc"

// 	_ "database/sql"
// 	// "github.com/weigj/odbc"
// )

// func main() {

// 	conn, err := sql.Open("GODBC", "driver={SQL Server};SERVER=localhost;UID=sa;PWD=t7710301988;")
// 	// conn, err := odbc.Connect("DSN=GODBC", "driver={SQL Server};SERVER=localhost;UID=sa;PWD=t7710301988;")
// 	// conn, err := odbc.Open("DSN=GODBC")
// 	if err != nil {
// 		fmt.Println("Connecting Error")
// 		return
// 	}
// 	fmt.Println("Connecting OK")
// 	defer conn.Close()
// 	// stmt, err := conn.Prepare("SELECT [ID],[Name],[School],[Sex] FROM [dbo].[Student]");
// 	// if(err!=nil){
// 	// fmt.Println("Query Error",err);
// 	// return;
// 	// }
// 	// defer stmt.Close();
// 	// row,err := stmt.FetchAll();
// 	// if err!=nil {
// 	// fmt.Println("Query Error",err);
// 	// return;
// 	// }
// 	// defer row.Close();
// 	// for row.Next() {
// 	// var id string;
// 	// if err := row.Scan(&id);err==nil {
// 	// fmt.Println(id);
// 	// }
// 	// }
// 	// fmt.Printf("%s\n","finish");
// 	// return;
// }

package main

import (
	"fmt"

	"github.com/appleboy/com/random"
)


func main() {
	fmt.Println("Hello world A b c")
	fmt.Println(random.String(10))
}
