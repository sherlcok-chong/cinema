package controller

import (
	"fmt"
	"github.com/go-lib-utils-master/time"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	time2 "time"
	"web/Cinema/modle"
	"web/Cinema/utils"
)

const PATH = "D:\\gogo\\src\\web\\Cinema\\view\\static\\img\\"

func UpLoadImg(f multipart.File, h *multipart.FileHeader, err error) string {
	if err != nil {
		log.Println(err)
		return ""
	}
	if err != nil {
		fmt.Println("one", err)
	}
	fileName := h.Filename
	fmt.Println(fileName)
	t, err := os.Create(PATH + fileName)
	if err != nil {
		log.Println(err)
	}
	if _, err := io.Copy(t, f); err != nil {
		fmt.Println(err)
	}

	return "static/img/" + fileName
}
func checkIsManager(w http.ResponseWriter, r *http.Request) bool {
	is,see:=IsLogin(r)
	if !is {
		t := template.Must(template.ParseFiles("view/pages/users/login.html"))
		t.Execute(w, r)
		return false
	}
	is=modle.CheckUserState(see.UserName)
	if !is {
		GetMoviePage(w,r)
		return false
	}
	return true
}
func AddInfo(w http.ResponseWriter, r *http.Request) {

	if !checkIsManager(w,r){
		return
	}

	t := template.Must(template.ParseFiles("view/pages/manager/addMovie.html"))
	flag := r.FormValue("flag")
	t.Execute(w, flag)
}
func AddMovie(w http.ResponseWriter, r *http.Request) {
	if !checkIsManager(w,r){
		return
	}
	movieName := r.PostFormValue("movieName")
	ge, _ := strconv.ParseFloat(r.PostFormValue("ge"), 0)
	point, _ := strconv.ParseFloat(r.PostFormValue("point"), 0)
	f, h, err := r.FormFile("test")
	ImgPath := UpLoadImg(f, h, err)
	flag, _ := strconv.ParseInt(r.PostFormValue("flag"), 10, 0)
	release := r.PostFormValue("release")
	timeLong := r.PostFormValue("timelong")
	offTime := r.PostFormValue("offtime")

	fmt.Println(movieName, ge, point, ImgPath, err, flag, release, timeLong, offTime)
	movie := modle.Movie{
		MovieName:   movieName,
		MovieGrade:  ge + point/10,
		BoxOffice:   0,
		MovieFlag:   int(flag),
		TimeLong:    timeLong,
		ImgPath:     ImgPath,
		ReleaseTime: release,
		OffTime:     offTime,
	}
	modle.AddMovie(movie)
	http.Redirect(w, r, "http://localhost:1122/addInfo?flag="+"1", http.StatusMovedPermanently)
}
func AddPlan(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("view/pages/manager/addPlan.html"))
	movies := modle.GetAllMovies()
	flag := r.FormValue("flag")
	type temp struct {
		Movies []string
		Flag   string
	}
	t.Execute(w, temp{movies, flag})
}
func AddMyPlan(w http.ResponseWriter, r *http.Request) {
	if !checkIsManager(w,r){
		return
	}

	movieName := r.PostFormValue("movieName")
	dateStart := r.PostFormValue("date")
	dateStart = strings.Replace(dateStart, "T", " ", -1)
	dateStart = dateStart + ":00"

	end := GetEndDate(movieName, dateStart)
	dateEnd := time.ParseDataTimeToStr(end)
	start, _ := time.ParseDateTime(dateStart)
	hall := r.PostFormValue("hall")
	money := r.PostFormValue("money")
	elseSpe := r.PostFormValue("elseSpe")
	re := CheckDate(start, end, hall)
	if re {
		plan := modle.PerfPlan{
			PerfID:    utils.CreateUUID(),
			MovieName: movieName,
			StartTime: dateStart,
			EndTime:   dateEnd,
			Halls:     hall,
			Money:     money,
			ElseSpe:   elseSpe,
		}
		modle.AddPlan(plan)
		modle.AddSit(plan.PerfID, hall)
		http.Redirect(w, r, "http://localhost:1122/addPlan?flag="+"增加成功", http.StatusMovedPermanently)
	} else {
		http.Redirect(w, r, "http://localhost:1122/addPlan?flag="+"请检查时间是否冲突", http.StatusMovedPermanently)
	}
}
func GetEndDate(movieName, dateStart string) time2.Time {
	timeLong := modle.GetTimeLongAndImgPath(movieName)
	fmt.Println(dateStart)
	start, err := time.ParseDateTime(dateStart)
	if err != nil {
		log.Println(err)
	}
	mm, err := time2.ParseDuration(timeLong + "m")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(mm)
	end := start.Add(mm)
	return end
}
func CheckDate(dateStart, dateEnd time2.Time, hall string) bool {

	dates := modle.GetPlanByHall(hall)
	for _, v := range dates {
		t, err := time.ParseDateTime(v)
		if err != nil {
			log.Println("??", err)
		}
		if t.After(dateStart) && t.Before(dateEnd) {
			return false
		}
	}
	return true
}
func DeletePlan(w http.ResponseWriter, r *http.Request) {
	if !checkIsManager(w,r){
		return
	}
	perfID := r.FormValue("perfID")
	flag := modle.QueryHas(perfID)
	if flag {
		modle.DeletePlan(perfID)
		http.Redirect(w, r, "http://localhost:1122/deleteInfo", http.StatusMovedPermanently)
	} else {
		http.Redirect(w, r, "http://localhost:1122/deleteInfo?flag="+"该影片存在售出，不可删除", http.StatusMovedPermanently)
	}

}

func LoadAllPlan(w http.ResponseWriter, r *http.Request) {
	if !checkIsManager(w,r){
		return
	}
	t := template.Must(template.ParseFiles("view/pages/manager/deletePlan.html"))
	plans := modle.GetAllPlans()
	flag := r.FormValue("flag")
	type temp struct {
		Plans []modle.PerfPlan
		Flag  string
	}
	t.Execute(w, temp{plans,flag})
}
