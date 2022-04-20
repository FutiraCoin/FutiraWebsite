package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"regexp"
	"strconv"

	//ghandlers "github.com/gorilla/handlers"
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	//"tawesoft.co.uk/go/dialog"
	//"gopkg.in/gomail.v2"
)

/* */
// SERVER PORT
const (
	PORT = ":8080"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "ayazaky"
	//dbname   = "postgres"
)

// Request is a struct decsribing request info
type Request struct {
	ID          int     `json:"id"`
	Email       string  `json:"email"`
	FullName    string  `json:"fullName"`
	TelegramId  string  `json:"telegramId"`
	PhoneNumber string  `json:"mobileNumber"`
	Quantity    float64 `json:"quantity"`
	WalletId    string  `json:"walletId"`
	PayMethod   string  `json:"payMethod"`
	Created_at  string  `json:"created_at"`
	Errors      map[string]string
}
type Team struct {
	ID          int    `json:"id"`
	Img         string `json:"img"`
	Name        string `json:"name"`
	JobTitle    string `json:"jobtitle"`
	Description string `json:"description"`
	LinkedIn    string `json:"linkedin"`
	Facebook    string `json:"facebook"`
	Twitter     string `json:"twiter"`
	Status      string `json:"status"`
}

// Function to check errors
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// connection to postgresql db
func CreateConnection(dbname string) *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	CheckError(err)
	// Check whether we can access the database by pinging it
	err = db.Ping()
	CheckError(err)
	return db
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
}

var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var rxTel = regexp.MustCompile("^@")
var rxwalId = regexp.MustCompile("^[0-9a-zA-Z]+$")

func (request *Request) Validate() bool {

	request.Errors = make(map[string]string)
	match := rxEmail.MatchString(request.Email)
	matcht := rxTel.MatchString(request.TelegramId)
	matchw := rxwalId.MatchString(request.WalletId)
	if match == false {
		request.Errors["Email"] = "Please enter a valid email address"
		fmt.Println(request.Errors)

	}
	if matcht == false {
		request.Errors["TelegramId"] = "Please enter valid teligtam id"
		fmt.Println(request.Errors)

	}
	if matchw == false {
		request.Errors["WalletId"] = "Please enter valid Wallet_id "
		fmt.Println(request.Errors)
	}

	return len(request.Errors) == 0
}

// Buy section
func RequesttHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		whoHandler(w, r)
		//render(w, "./template/index.html", nil)

	}
	if r.Method == "POST" {

		// Extract the field information about the request from the form info
		request := Request{}
		request.Email = r.PostFormValue("email")
		request.FullName = r.PostFormValue("fullname")
		request.TelegramId = r.PostFormValue("telegramid")
		request.PhoneNumber = r.PostFormValue("phonenumber")
		request.Quantity, _ = strconv.ParseFloat(r.PostFormValue("Quantity"), 64) // FormValue is always string so we well convert the result into float64
		request.WalletId = r.PostFormValue("WalletID")
		request.PayMethod = r.PostFormValue("PaymentMethod")
		if request.Validate() == false {

			render(w, "./template/index.html", request)
			return

		} else {

			db := CreateConnection("postgres")
			_, err1 := db.Query(
				`INSERT INTO public.request_info(email, fullname, telegram_id, mobile_num, quantity, wallet_id, payment_mathod)
					VALUES ($1,$2,$3,$4,$5,$6,$7);`, request.Email, request.FullName, request.TelegramId, request.PhoneNumber, request.Quantity, request.WalletId, request.PayMethod)
			CheckError(err1)
			//////////////////////////////////////////////////////////////////
			from := "alaaelshreif2018@gmail.com"
			toList := []string{"aya7zaky@gmail.com"}
			subj := "New Futira Buy Request"
			body := string("\r\n" + "Email:\r\n" + request.Email + "\r\n\r\n" + "Fullname:\r\n" + request.FullName + "\r\n\r\n" + "Telegram ID:\r\n" + request.TelegramId + "\r\n\r\n" + "Phonenumber:\r\n" + request.PhoneNumber + "\r\n\r\n" + "Quantity:\r\n" + strconv.FormatFloat(request.Quantity, 'f', -1, 64) + "\r\n\r\n" + "WalletID:\r\n" + request.WalletId + "\r\n\r\n" + "Payment Method:\r\n" + request.PayMethod)

			headers := make(map[string]string)
			headers["From"] = from
			headers["To"] = "aya7zaky@gmail.com"
			headers["Subject"] = subj

			message := ""
			for k, v := range headers {
				message += fmt.Sprintf("%s: %s\r\n", k, v)
			}
			message += "\r\n" + body
			host := "smtp.gmail.com"
			// Its the default port of smtp server
			port := "587"
			auth := smtp.PlainAuth("", "alaaelshreif2018@gmail.com", "NEVERgiveup2020", host)
			err := smtp.SendMail(host+":"+port, auth, from, toList, []byte(message))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)

			}

			fmt.Println("Successfully sent mail to all user in toList")
		}

	}
}

// who section
func whoHandler(w http.ResponseWriter, r *http.Request) {
	dbname := "DashBoard"
	db := CreateConnection(dbname)
	rows2, err2 := db.Query(`SELECT * FROM public.team  WHERE status='Active' ORDER BY id;`)
	CheckError(err2)

	defer rows2.Close()
	Members := []Team{}
	for rows2.Next() {
		t := Team{}
		err := rows2.Scan(&t.ID, &t.Img, &t.Name, &t.JobTitle, &t.Description, &t.LinkedIn, &t.Facebook, &t.Twitter, &t.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
		Members = append(Members, t)
	}

	render(w, "./template/index.html", Members)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/buy", RequesttHandler).Methods("GET", "POST")
	r.HandleFunc("/who", whoHandler)

	// serving static files
	r.PathPrefix("/template/").
		Handler(http.StripPrefix("/template/", http.FileServer(http.Dir("template"))))

	// work in port 8080
	http.ListenAndServe(PORT, r)

}
