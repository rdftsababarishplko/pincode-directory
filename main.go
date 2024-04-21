package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var templates = template.Must(template.ParseGlob("/home/eye/go/pincode-directory/templates/*.html"))

type Pincode struct {
	Pincode       string
	Office_Type   string
	Circle_Name   string
	Region_Name   string
	Division_Name string
	Office_Name   string
	Delivery      string
	District      string
	StateName     string
}

func main() {
	var err error
	db, err = sql.Open("sqlite3", "/home/eye/go/pincode-directory/data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/read-pincode", readPincodeHandler)
	http.HandleFunc("/create-pincode", createPincodeHandler)
	http.HandleFunc("/update-pincode", updatePincodeHandler)
	http.HandleFunc("/delete-pincode", deletePincodeHandler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}

func readPincodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		pincode := r.Form.Get("pincode")

		row := db.QueryRow("SELECT * FROM pincode WHERE Pincode = ?", pincode)
		var pincodeDetails Pincode
		err := row.Scan(&pincodeDetails.Pincode, &pincodeDetails.Office_Type, &pincodeDetails.Circle_Name,
			&pincodeDetails.Region_Name, &pincodeDetails.Division_Name, &pincodeDetails.Office_Name,
			&pincodeDetails.Delivery, &pincodeDetails.District, &pincodeDetails.StateName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		renderTemplate(w, "read_pincode", pincodeDetails)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func createPincodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse form data for pin code details
		pincode := r.FormValue("pincode")
		officeType := r.FormValue("officeType")
		circleName := r.FormValue("circleName")
		regionName := r.FormValue("regionName")
		division := r.FormValue("division")
		officeName := r.FormValue("officeName")
		delivery := r.FormValue("delivery")
		district := r.FormValue("district")
		stateName := r.FormValue("stateName")

		// Insert new pin code into database
		_, err := db.Exec("INSERT INTO pincode VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
			pincode, officeType, circleName, regionName, division, officeName, delivery, district, stateName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		renderTemplate(w, "create_pincode", nil)
	}
}

func updatePincodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		// Retrieve pincode from form
		pincode := r.Form.Get("pincode")

		// Query database for pin code details to update
		row := db.QueryRow("SELECT * FROM pincode WHERE Pincode = ?", pincode)
		var pincodeDetails Pincode
		err := row.Scan(&pincodeDetails.Pincode, &pincodeDetails.Office_Type, &pincodeDetails.Circle_Name,
			&pincodeDetails.Region_Name, &pincodeDetails.Division_Name, &pincodeDetails.Office_Name,
			&pincodeDetails.Delivery, &pincodeDetails.District, &pincodeDetails.StateName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Render the update form with pin code details
		renderTemplate(w, "update_pincode_form", pincodeDetails)
	}
}

func deletePincodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		// Retrieve pin code from form
		pincode := r.Form.Get("pincode")

		// Query database to check if pin code exists
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM pincode WHERE Pincode = ?", pincode).Scan(&count)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check if pin code exists in the database
		if count == 0 {
			http.Error(w, "Pin code not found", http.StatusNotFound)
			return
		}

		// Perform delete operation
		_, err = db.Exec("DELETE FROM pincode WHERE Pincode = ?", pincode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to home page or display success message
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
