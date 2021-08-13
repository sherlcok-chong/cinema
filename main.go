package main

import (
	"net/http"
	"web/Cinema/controller"
)

func main() {
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("view/static"))))
	http.Handle("/pages/",http.StripPrefix( "/pages/",http.FileServer(http.Dir("view/pages"))))
	http.HandleFunc("/main",controller.GetMoviePage)
	http.HandleFunc("/login",controller.Login)
	http.HandleFunc("/logout",controller.Logout)
	http.HandleFunc("/register",controller.Register)
	http.HandleFunc("/buyTicket",controller.BuyTicket)
	http.HandleFunc("/setSit",controller.SetSit)
	http.HandleFunc("/getTic",controller.GetTci)
	http.HandleFunc("/confirmTic",controller.ConfirmTci)
	http.HandleFunc("/buyOk",controller.BuyOk)
	http.HandleFunc("/myOrders",controller.MyOrder)
	http.HandleFunc("/refundTic",controller.RefundTic)
	http.HandleFunc("/addMovie",controller.AddMovie)
	http.HandleFunc("/addInfo",controller.AddInfo)
	http.HandleFunc("/addPlan",controller.AddPlan)
	http.HandleFunc("/addMyPlan",controller.AddMyPlan)
	http.HandleFunc("/deleteInfo",controller.LoadAllPlan)
	http.HandleFunc("/deletePlan",controller.DeletePlan)
	http.HandleFunc("/query",controller.SearchMovies)
	http.ListenAndServe(":1122",nil)
	//t:=time.GetNowDateTimeStr()
	//fmt.Println(t)
	//r ,_:= time.ParseDateTime(t)
	//fmt.Println(r)
	//t2 := time.GetNowTime()
	//fmt.Println(t2)
	//time2.Sleep(5*time2.Second)
	//s:=t2.After(r)
	//fmt.Println(s)
}
