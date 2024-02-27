//go:generate npm run build
package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	// local db:
	_ "github.com/libsql/go-libsql"
	// remote db:
	// _ "github.com/tursodatabase/libsql-client-go/libsql"
)

func logError(err error, msg string) {
	fmt.Fprintf(os.Stderr, msg, err)
}

type RscyG struct {
	Name, Email, Busyness string
	Id, Dopness           int // 0-100
}

func rootDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return strings.TrimSuffix(exPath, "/tmp")
}

func connectDb() *sql.DB {
	dir := rootDir()
	dbPath := "/db/rscy-gs.db"
	dbUrl := "file://" + dir + dbPath

	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		logError(err, "failed to open db %s")
		os.Exit(1)
	}
	return db
}

var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	templateNames := [3]string{"welcome", "form", "list"}
	for _, name := range templateNames {
		t, err := template.ParseFiles("views/layout.html", "views/"+name+".html")
		if err == nil {
			templates[name] = t
		} else {
			panic(err)
		}
	}
}

func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
	templates["welcome"].Execute(writer, nil)
}

func getAllRscyGs(db *sql.DB) ([]RscyG, error) {
	var rscyGs []RscyG
	var err error

	rows, err := db.Query(`
		SELECT * FROM RscyGs
	`)
	if err != nil {
		logError(err, "failed to execute query: %v\n")
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		var rscyG RscyG

		if err := rows.Scan(&rscyG.Id, &rscyG.Name, &rscyG.Email, &rscyG.Busyness, &rscyG.Dopness); err != nil {
			fmt.Println("Error scanning row:", err)
			return rscyGs, err
		}

		rscyGs = append(rscyGs, rscyG)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}

	return rscyGs, err
}

func createListHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		rscyGs, err := getAllRscyGs(db)
		if err != nil {
			logError(err, "error getting them rscy Gs: %v\n")
		}
		templates["list"].Execute(writer, rscyGs)
	}
}

type formData struct {
	*RscyG
	Errors []string
}

func createFormHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			templates["form"].Execute(writer, formData{
				RscyG: &RscyG{}, Errors: []string{},
			})
		}
		if request.Method == http.MethodPost {
			request.ParseForm()
			rscyGData := RscyG{
				Name:     request.Form["name"][0],
				Email:    request.Form["email"][0],
				Busyness: request.Form["busyness"][0],
				Dopness:  rand.Intn(101),
			}
			_, err := db.Query(`
				INSERT INTO RscyGs (Name, Email, Busyness, Dopness) VALUES (?, ?, ?, ?)
			`, rscyGData.Name, rscyGData.Email, rscyGData.Busyness, rscyGData.Dopness)
			if err != nil {
				logError(err, "failed to execute query: %v\n")
			}
			http.Redirect(writer, request, "/rscy", http.StatusPermanentRedirect)
		}
	}
}

func main() {
	db := connectDb()
	loadTemplates()
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/rscy", createListHandler(db))
	http.HandleFunc("/rscy/new", createFormHandler(db))
	fsHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/files/", http.StripPrefix("/files", fsHandler))
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}
