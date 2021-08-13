package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"web/Cinema/modle"
	"web/Cinema/utils"
)

func Login(w http.ResponseWriter,r *http.Request)  {
	flag,_:=IsLogin(r)
	if flag {
		GetMoviePage(w,r)
	} else {
		//获取用户名和密码
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		//调用测验函数
		fmt.Println(username, password)
		user, re := modle.CheckPassword(username, password)
		fmt.Println(user)
		if re != false {
			//正确
			fmt.Println("查询成功")
			uuid := utils.CreateUUID()
			sess := modle.Session{
				SessionID: uuid,
				UserName:  user.UserName,
				UserID:    user.UserID,
			}

			err := modle.AddSession(sess)
			if err != nil {
				fmt.Println(re)
			}
			//创建cookie
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "http://localhost:1122/main", http.StatusMovedPermanently)
		} else {
			//false
			fmt.Println("查询错误")
			t := template.Must(template.ParseFiles("view/pages/users/login.html"))
			t.Execute(w, "")
		}
	}
	//t := template.Must(template.ParseFiles("view/pages/users/login.html"))
	//t.Execute(w,r)
}
func Logout(w http.ResponseWriter, r *http.Request)  {
	cookie,_:=r.Cookie("user")
	if cookie != nil {
		cookieVal:=cookie.Value
		modle.DeleteSession(cookieVal)
		cookie.MaxAge=-1
	}
	GetMoviePage(w,r)
}
func IsLogin(r *http.Request) (bool,modle.Session)  {
	cookie , _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		//fmt.Println("cookieval",cookieValue)
		//查session
		session, err := modle.GetSessionById(cookieValue)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(session)
		if session.UserName != "" {
			return true,session
		}
		return false,modle.Session{}
	}
	return false,modle.Session{}
}
func Register(w http.ResponseWriter, r *http.Request){
	//t := template.Must(template.ParseFiles("view/pages/users/register.html"))
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	_, re := modle.CheckPassword(username, password)
	if re != false {
		//正确
		fmt.Println("增加失败")
		t := template.Must(template.ParseFiles("view/pages/user/register.html"))
		t.Execute(w, false)
	} else {
		//false
		fmt.Println("增加成功")
		userId := utils.CreateUUID()
		modle.InsertUser(modle.User{UserID: userId, UserName: username, Password: password, Email: email})
		GetMoviePage(w,r)
	}
}
func MyOrder(w http.ResponseWriter , r *http.Request)  {
	_,session:= IsLogin(r)
	tickets := modle.GetTicketByUserID(session.UserID)
	t := template.Must(template.ParseFiles("view/pages/users/myOrder.html"))
	t.Execute(w,tickets)
}
func RefundTic(w http.ResponseWriter, r *http.Request)  {
	orderID := r.FormValue("orderID")
	ticket:=modle.GetTicketByOrderID(orderID)
	modle.UpdateSit(ticket,"selectable")
	modle.DeleteTicByOrderID(orderID)
	MyOrder(w,r)
}