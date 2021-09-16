package main
import (
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
type Tag struct {
    ID   int    `json:"admin_id"`
    Name string `json:"admin_name"`
}
func main() {
    // Open up our database connection.
    db, err := sql.Open("mysql", "yourusername:yourpass@tcp(yourip:yourport)/yourdb")

    // if there is an error opening the connection, handle it
    if err != nil {
        log.Print(err.Error())
    }
    defer db.Close()

    // Execute the query
    results, err := db.Query("SELECT admin_id, admin_name FROM tbl_admin")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    for results.Next() {
        var tag Tag
        // for each row, scan the result into our tag composite object
        err = results.Scan(&tag.ID, &tag.Name)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                // and then print out the tag's Name attribute
        log.Printf(tag.Name)
    }

}