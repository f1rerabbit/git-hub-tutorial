package main

import (
    "fmt"
    "net/http"
    "html/template"

// Подключение библиотеки gorilla
"github.com/gorilla/mux"

    "database/sql"

    _ "github.com/go-sql-driver/mysql"
)

// Вывод статей на страницу
type Article struct {
  Id uint16
  Title, Anons, Fulltext string
}

var posts = []Article{}
var showPost = Article{}

// go mod init knocker иначе не запустится SQL
// go mod tidy
func index(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  db, err := sql.Open("mysql", "root:root@(127.0.0.1:8080)/golang?parseTime=true")
  if err != nil {
  panic(err)
  }
  // // Выборка данных
  // res. err := db.Query("SELECT * FROM `articles`")
  // if err != nil {
  //   panic(err)
  // }
  // for res.Next() {
  //   var user User
  //   err = res.Scan(&user.Namem &user.Age)
  //   if err != nil {
  //     panic(err)
  //   }

  // Выборка данных
  res. err := db.Query("SELECT * FROM `articles`")
  if err != nil {
    panic(err)
  }

// Решение как не дать копии постов
posts = []Article{}
  for res.Next() {
    var post Article
    err = res.Scan(&post.Id, &post.title, &post.Anons, &post.Fulltext)
    if err != nil {
      panic(err)
    }
  posts = append(posts, post)

  // fmt.Println(fmt.Sprintf("Post: %s with id: %d", post.Title, post.Id))
}
// параметр указывающий какой блок хочу указывать на странице
  t.ExecuteTemplate(w, "index", posts)
}

func create(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
// Проверка на ошибку
  if err != nil {
    fmt.Fprintf(w, err.Error())
  }
// параметр указывающий какой блок хочу указывать на странице
  t.ExecuteTemplate(w, "create", nil)
}
//подключение страницы
func save_article(w http.ResponseWriter, r *http.Request) {
title := r.FormValue("title")
anons := r.FormValue("anons")
full_text := r.FormValue("full_text")

// Проверка
if title == "" || anons == "" || full_text == "" {
  fmt.Fprintf(w, "Не все данные заполнены")
} else  {

  // подключение к бд
  db, err := sql.Open("mysql", "root:root@(127.0.0.1:8080)/golang?parseTime=true")
  if err != nil {
  panic(err)
  }
  defer db.Close()
  // установка данных
  insert, err := db.Query(fmt.Sprintf("INSERT INTO `article` (`title`, `anons`, `full_text`) VALUES('%s', '%s', '%s')", title, anons, full_text))
  if err != nil {
    panic(err)
  }
  defer insert.Close()

  http.Redirect(w, r, "/", http.StatusSeeOther)
  }
}
// откуда брать
func show_post(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")
// Проверка на ошибку
  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  db, err := sql.Open("mysql", "root:root@(127.0.0.1:8080)/golang?parseTime=true")
  if err != nil {
  panic(err)
  }
  defer db.Close()
  // Выборка данных
  res. err := db.Query(fmt.Sprintf("SELECT * FROM `articles` WHERE `id` = '%s'", vars["id"]))
  if err != nil {
    panic(err)
  }
// Решение как не дать копии постов
showPost = Article{}
  for res.Next() {
    var post Article
    err = res.Scan(&post.Id, &post.title, &post.Anons, &post.Fulltext)
    if err != nil {
      panic(err)
    }
  showPost = post
}
// параметр указывающий какой блок хочу указывать на странице
  t.ExecuteTemplate(w, "index", showPost)
}

func HandleFunc() {
    rtr := mux.NewRouter()
    rtr.HandleFunc("/", index).Methods("GET")
    rtr.HandleFunc("/create", create).Methods("GET")
    rtr.HandleFunc("/save_article", save_article).Methods("POST")
    rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET")
    /post/

    http.Handle("/", rtr)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))) // новые адреса и убираем адрес

    http.ListenAndServe(":8080", nil)
    }




func main()  {
    HandleFunc()

}
