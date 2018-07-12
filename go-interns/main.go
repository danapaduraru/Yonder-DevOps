package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var templates = template.Must(template.ParseGlob("templates/*"))
var dbConfiguration = readConfiguration()
type configuration struct {
	DBName   string
	User     string
	Password string
	Host     string
	Port     string
}

//PageData holds the data that is inserted in the HTML template
type PageData struct {
	Name            string
	PostgresName    string
	PostgresVersion string
}

//NewPageDataAnon Create a new anonymous instance of pageData.
func NewPageDataAnon(pname string, pversion string) PageData {
	var data PageData
	data.Name = "Yonderist"
	data.PostgresName = pname
	data.PostgresVersion = pversion
	return data
}

//NewPageData Create a new instance of pageData
func NewPageData(name string, pname string, pversion string) PageData {
	var data PageData
	data.Name = name
	data.PostgresName = pname
	data.PostgresVersion = pversion
	return data
}

func main() {
	var listenPort = flag.String("port", "18080", "This is the port the application will bind to.")
	flag.Parse()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", helloAnonymous).Methods("GET")
	router.HandleFunc("/{token}", hello).Methods("GET")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *listenPort), router))
}
func readConfiguration() configuration {
	cfile, _ := os.Open("conf.json")
	defer cfile.Close()
	decoder := json.NewDecoder(cfile)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		log.Print(err)
	}
	return conf
}
func testDB(c configuration) string {
	var version string
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", c.Host, c.Port, c.DBName, c.User, c.Password)
	db, _ := sql.Open("postgres", connStr)
	rows, err := db.Query(`select version()`)
	if err != nil {
		log.Print(err)
		return ""
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&version); err != nil {
			log.Print(err)
		}

	}
	return version
}

func logRequest(remoteIP string, method string, path string) {
	log.Printf("%s call on endpoint %s received from: %s\n", method, path, remoteIP)
}
func helloAnonymous(w http.ResponseWriter, r *http.Request) {
	logRequest(r.RemoteAddr, r.Method, r.URL.RequestURI())
	dbresponse := testDB(dbConfiguration)
	if dbresponse == "" {
		templates.ExecuteTemplate(w, "error", nil)
	} else {
		dbdata := strings.Split(dbresponse, " ")
		var data = NewPageDataAnon(dbdata[0], dbdata[1])
		templates.ExecuteTemplate(w, "success", data)
	}
}
func hello(w http.ResponseWriter, r *http.Request) {
	logRequest(r.RemoteAddr, r.Method, r.URL.RequestURI())
	vars := mux.Vars(r)
	dbresponse := testDB(dbConfiguration)
	if dbresponse == "" {
		templates.ExecuteTemplate(w, "error", nil)
	} else {
		dbdata := strings.Split(dbresponse, " ")
		var data = NewPageData(vars["token"], dbdata[0], dbdata[1])
		templates.ExecuteTemplate(w, "success", data)
	}
	
}
