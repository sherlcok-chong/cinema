package controller

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"web/Cinema/modle"
	"web/Cinema/utils"
)

func GetMoviePage(w http.ResponseWriter, r *http.Request) {
	flag, session := IsLogin(r)
	movieFace := modle.GetMovieFace()
	if flag {
		movieFace.UserName = session.UserName
	}
	re:=modle.CheckUserState(session.UserName)
	fmt.Println(re)
	if re {
		t := template.Must(template.ParseFiles("view/pages/manager/home.html"))
		t.Execute(w, movieFace)
	} else {
		t := template.Must(template.ParseFiles("view/index.html"))
		t.Execute(w, movieFace)
	}

}
func BuyTicket(w http.ResponseWriter, r *http.Request) {
	flag, _ := IsLogin(r)
	if !flag {
		t := template.Must(template.ParseFiles("view/pages/users/login.html"))
		t.Execute(w, r)
	} else {
		movieName := r.FormValue("movieName")
		fmt.Println(movieName)
		movie := modle.GetMovieByName(movieName)
		t := template.Must(template.ParseFiles("view/pages/users/buy_ticket.html"))
		fmt.Println(movie)
		t.Execute(w, movie)
	}

}
func SetSit(w http.ResponseWriter, r *http.Request) {

	perfID := r.FormValue("perfID")
	fmt.Println(perfID)
	sit := modle.GetSitByID(perfID)
	fmt.Println(sit)
	t := template.Must(template.ParseFiles("view/pages/users/sit.html"))
	t.Execute(w, sit)
}
func GetTci(w http.ResponseWriter, r *http.Request) {
	modle.DeleteTicByFlag()
	tickets := make([]modle.Ticket, 0)
	if r.Method == "POST" { //监听是否为POST方法
		b, err := ioutil.ReadAll(r.Body)
		_, sess := IsLogin(r)
		perID := r.FormValue("perfID")

		perf := modle.GetPerfPlanByID(perID)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(b)
		for i := 10; i < len(b); i += 11 {
			r := int64(b[i] - '0')
			l := int64(b[i+1] - '0')
			orderid := utils.CreateUUID()
			ticket := modle.Ticket{
				OrderID:   orderid,
				UserID:    sess.UserID,
				PerfID:    perID,
				MovieName: perf.MovieName,
				Time:      perf.StartTime,
				Row:       r,
				Line:      l,
				Hall:      perf.Halls,
				Money:     perf.Money,
				ImgPath:   perf.ImgPath,
			}
			tickets = append(tickets, ticket)
		}

	}
	err:=modle.AddTicket(tickets)
	if err != nil {
		SetSit(w,r)
	}
}

func ConfirmTci(w http.ResponseWriter, r *http.Request) {
	tickets := modle.GetTicketByFlag()
	fmt.Println(tickets)
	t := template.Must(template.ParseFiles("view/pages/users/confirmTic.html"))
	t.Execute(w, tickets)

}
func BuyOk(w http.ResponseWriter, r *http.Request) {
	tickets := modle.GetTicketByFlag()
	modle.UpdateTicByFlag()
	modle.UpdateSit(tickets,"sold")
	t := template.Must(template.ParseFiles("view/pages/users/BuyOk.html"))
	t.Execute(w, r)
}
func SearchMovies(w http.ResponseWriter, r *http.Request){
	movieName := r.PostFormValue("kw")
	movies := modle.GetMovies(movieName)
	t := template.Must(template.ParseFiles("view/pages/users/search.html"))
	t.Execute(w,movies)
}
