package main

import (
    "fmt"
    "time"
    "net/http"
    "html/template"
    "database/sql"

    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    Name string
    Age uint16
    Money int16
    Avg_grades, Hapiness float64
    Hobbies []string
}

type User1 struct {
    Username string
    Email string
}

func (u User) getAllInfo() string {
    return fmt.Sprintf("User name is: %s. He is %d and he has money "+
    "equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
    u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request){
    //bob := User{Name: "Bob", Age: 25, Money: -50, Avg_grades: 4.2, Hapiness: 0.8}
    bob := User{"Bob", 25, -50, 4.2, 0.8, []string {"Beer","Eat","Sleep"}}
    //fmt.Println("Go рулит!")
    //fmt.Fprintf(w, "Go is super easy")
    //fmt.Fprintf(w, bob.Name)
    //bob.setNewName("Ihor")
    //fmt.Fprintf(w, bob.getAllInfo())
    tmpl, _ := template.ParseFiles("templates/home_page.html")
    tmpl.Execute(w, bob)
}

func about_page(w http.ResponseWriter, r *http.Request){
    fmt.Println("Go рулит!")
    fmt.Fprintf(w, "About page ")
}

func handleRequest(){
    http.HandleFunc("/", home_page)
    http.HandleFunc("/about/", about_page)
    http.ListenAndServe(":8090", nil)
}

func main (){
    //var bob user =
    // handleRequest()
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/rcmoney")
    if err != nil { panic(err) }

    defer db.Close()
    //insert, err := db.Query("INSERT INTO `tbl_user` (`ord`, `user_create`, `username`, `password`, `email`, `hash`) VALUES(0, 6, 'kika', 'pass', 'kika@kika.test', 'test33')")

    //if err != nil { panic(err) }

    //defer insert.Close()

    res, err := db.Query("SELECT `username`, `email` FROM `tbl_user`")

    if err != nil {panic(err)}

    for res.Next(){
        var user User1
        err = res.Scan(&user.Username, &user.Email)
        if err != nil {panic(err)}
        fmt.Println("User: "+user.Username+"/ email: "+user.Email)

    }

        defer res.Close()
        db.SetConnMaxLifetime(time.Minute * 3)
        db.SetMaxOpenConns(10)
        db.SetMaxIdleConns(10)

        fmt.Println("mysql")
}

