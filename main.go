//go:generate npm run build
package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	env "github.com/gofor-little/env"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func logError(err error, msg string) {
	fmt.Fprintf(os.Stderr, msg, err)
}

type RscyG struct {
	Name, Email, Busyness string
	Id, Dopness           int // 0-100
	CreatedAt             time.Time
}

func connectDb() *sql.DB {
	dbPath := env.Get("DB_URL", "")
	dbToken := env.Get("DB_TOKEN", "")
	dbUrl := dbPath + "?authToken=" + dbToken
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

func getAllRscyGs(db *sql.DB, limit int, page int) ([]RscyG, error) {
	var rows *sql.Rows
	var rscyGs []RscyG
	var err error
	isNoLimit := limit < 0
	if isNoLimit {
		rows, err = db.Query(`
			SELECT * FROM RscyGs ORDER BY CreatedAt DESC;
		`)
	} else {
		offset := (page - 1) * limit
		rows, err = db.Query(`
		SELECT * FROM RscyGs ORDER BY CreatedAt DESC LIMIT ? OFFSET ?;
	`, limit, offset)
	}
	if err != nil {
		logError(err, "failed to execute query: %v\n")
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		var rscyG RscyG

		if err := rows.Scan(&rscyG.Id, &rscyG.Name, &rscyG.Email, &rscyG.Busyness, &rscyG.Dopness, &rscyG.CreatedAt); err != nil {
			fmt.Println("Error scanning row:", err)
			return rscyGs, err
		}
		fmt.Println("sup bro", rscyG.CreatedAt)
		rscyGs = append(rscyGs, rscyG)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}

	return rscyGs, err
}

func createListHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryLimit := request.URL.Query().Get("limit")
		queryPage := request.URL.Query().Get("page")
		limit, err := strconv.ParseInt(queryLimit, 0, 64)
		if err != nil {
			limit = -1
		}
		page, err := strconv.ParseInt(queryPage, 0, 64)
		if err != nil {
			page = 1
		}
		rscyGs, err := getAllRscyGs(db, int(limit), int(page))
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
				INSERT INTO RscyGs (Name, Email, Busyness, Dopness) VALUES (?, ?, ?, ?);
			`, rscyGData.Name, rscyGData.Email, rscyGData.Busyness, rscyGData.Dopness)
			if err != nil {
				logError(err, "failed to execute query: %v\n")
			}
			http.Redirect(writer, request, "/rscy?limit=10&page=1", http.StatusPermanentRedirect)
		}
	}
}

func main() {
	env.Load("./.env")
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
