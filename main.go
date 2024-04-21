package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// Define a global database connection
var db *sql.DB

func main() {
	// Open the SQLite database
	var err error
	db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	// Create tables if they don't exist
	if err := createTables(); err != nil {
		log.Fatal("Error creating database tables:", err)
	}

	// Define HTTP routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/read-pincode", readPincodeHandler)
	http.HandleFunc("/create-pincode", createPincodeHandler)
	http.HandleFunc("/update-pincode", updatePincodeHandler)
	http.HandleFunc("/delete-pincode", deletePincodeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the HTTP server
	log.Println("Server listening on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

// indexHandler serves the index.html page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

// readPincodeHandler serves the read_pincode.html page
func readPincodeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "read_pincode.html", nil)
}

// createPincodeHandler serves the create_pincode.html page
func createPincodeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "create_pincode.html", nil)
}

// updatePincodeHandler serves the update_pincode.html page
func updatePincodeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "update_pincode.html", nil)
}

// deletePincodeHandler serves the delete_pincode.html page
func deletePincodeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "delete_pincode.html", nil)
}

// renderTemplate renders the specified HTML template
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpl = filepath.Join("templates", tmpl)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
	}
}

// createTables initializes the database schema
func createTables() error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS pincode (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			area TEXT NOT NULL,
			pincode INTEGER NOT NULL
		);
	`)
	return err
}
