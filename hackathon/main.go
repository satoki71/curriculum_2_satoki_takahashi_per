package main

//at first we should "source .envmysql", "env | grep MYSQL_DATABASE"
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"hackathon/controller"
	"hackathon/dao"
	"log"
	"net/http"
	"os"
)

// ① GoプログラムからMySQLへ接続
var db *sql.DB

func init() {
	// DB接続のための準備
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	_db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db

	//err := godotenv.Load(".envmysql")
	//if err != nil {
	//	log.Fatalf("fail: load envfile, %v\n", err)
	//}
	//// ①-1
	//mysqlUser := os.Getenv("MYSQL_USER")
	//mysqlUserPwd := os.Getenv("MYSQL_PASSWORD")
	//mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	//
	//// ①-2
	//_db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(localhost:3306)/%s", mysqlUser, mysqlUserPwd, mysqlDatabase))
	//if err != nil {
	//	log.Fatalf("fail: sql.Open, %v\n", err)
	//}
	//// ①-3
	//if err := _db.Ping(); err != nil {
	//	log.Fatalf("fail: _db.Ping, %v\n", err)
	//}
	//db = _db

}

func affiliationHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Headers", "http://localhost:3000")
	//w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	//w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	//w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		controller.AffiliationRegister(w, r)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
func handler(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Access-Control-Allow-Headers", "http://localhost:3000")
	//w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	//w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	//w.Header().Set("Content-Type", "application/json")
	//se

	switch r.Method {
	case http.MethodGet:
		controller.UserSearch(w, r)
		// ②-1

	case http.MethodPost:
		controller.UserRegister(w, r)

	case http.MethodPut:
		controller.UserUpdate(w, r)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func handler2(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Access-Control-Allow-Headers", "http://localhost:3000")
	//w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	//w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	//w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		controller.AllUserSearch(w, r)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func handler3(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Headers", "http://localhost:3000")
	//w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	//w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	//w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		controller.MemberUserSearch(w, r)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func takeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Headers", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		controller.TakeSearch(w, r)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func giveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	//giveをみる
	case http.MethodGet:
		controller.GiveSearch(w, r)

	//giveする
	case http.MethodPost:
		controller.GiveRegister(w, r)

	//giveを編集する(Put)
	case http.MethodPut:
		controller.GiveUpdate(w, r)

	//giveを削除する
	case http.MethodDelete:
		controller.GiveDelete(w, r)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func main() {
	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	http.HandleFunc("/user", handler)

	http.HandleFunc("/users", handler2)

	http.HandleFunc("/members", handler3)

	http.HandleFunc("/take", takeHandler)

	http.HandleFunc("/give", giveHandler)

	http.HandleFunc("/affiliation", affiliationHandler)

	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	dao.CloseDBWithSysCall()

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
