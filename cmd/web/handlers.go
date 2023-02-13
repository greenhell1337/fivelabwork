package web

import (
	"html/template"
	"net/http"

	ce "github.com/greenhell1337/shopForDB/pkg/checkerr"
	"github.com/greenhell1337/shopForDB/pkg/models/address"
	client "github.com/greenhell1337/shopForDB/pkg/models/clients"
)

// main page.
func Main(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/html/main.html")
	ce.CheckErr(err)

	tmpl.Execute(w, nil)
}

// list tea page.
func TeaList(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/html/tealist.html")
	ce.CheckErr(err)

	tmpl.Execute(w, nil)
}

// abount page.
func About(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles()
	ce.CheckErr(err)

	tmpl.Execute(w, nil)
}

// sign up page.
func SignUp(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/html/signup.html")
	ce.CheckErr(err)

	entity := &client.Client{
		Name: r.FormValue("nameReg"),
		Surname: r.FormValue("surnameReg"),
		Date: map[string]string{
			r.FormValue("emailReg"): r.FormValue("passwordReg"),
		},
		Address: address.Address{
			City: r.FormValue("cityReg"),
			Street: r.FormValue("streetReg"),
			NumberHome: r.FormValue("numberHomeReg"),
			NumberFlat: r.FormValue("humberFlatReg"),
		},
	}

	tmpl.Execute(w, nil)
}

// log in page.
func LogIn(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles()
	ce.CheckErr(err)

	tmpl.Execute(w, nil)
}

// basket page.
func Basket(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles()
	ce.CheckErr(err)

	tmpl.Execute(w, nil)
}


// China list tea page.
func ChinaList(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/html/chinalist.html")
	ce.CheckErr(err)

	tmpl.Execute(w, nil)
}

// Bags list tea page.
func BagsList(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles()
	ce.CheckErr(err)

	tmpl.Execute(w, nil)
}

// Pyramids list tea page.
func PyramidsList(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles()
	ce.CheckErr(err)

	tmpl.Execute(w, nil)
}

// Siithing list tea page.
func SoothingList(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles()
	ce.CheckErr(err)

	tmpl.Execute(w, nil)
}