package modle

import (
	"fmt"
	"github.com/go-lib-utils-master/time"
	"log"
	time2 "time"
	"web/Cinema/utils"
)

type PerfPlan struct {
	PerfID    string
	MovieName string
	ImgPath   string
	StartTime string
	EndTime   string
	Halls     string
	ElseSpe   string
	Money     string
}

func GetPerfPlanByName(movieName string) []PerfPlan {
	perfPlan := make([]PerfPlan, 0)
	rows, _ := utils.Db.Query("select * from performance_plan where movie_name = ?", movieName)
	for rows.Next() {
		temp := PerfPlan{}
		rows.Scan(&temp.PerfID, &temp.MovieName, &temp.StartTime, &temp.EndTime, &temp.Halls, &temp.Money, &temp.ElseSpe)
		if !DeleteExpiredPlan(temp.PerfID,temp.StartTime) {
			perfPlan = append(perfPlan, temp)
		}
	}
	return perfPlan
}
func DeleteExpiredPlan(perfID,startTime string) bool {
	now := time2.Now()
	start,_ := time.ParseDateTime(startTime)
	if now.After(start){
		utils.Db.Exec("delete from performance_plan where performance_id = ?",perfID)
	}
	return now.After(start)
}
func GetPerfPlansByID(perfID string) []PerfPlan {
	perfPlan := make([]PerfPlan, 0)
	rows, _ := utils.Db.Query("select * from performance_plan where performance_id = ?", perfID)
	for rows.Next() {
		temp := PerfPlan{}
		rows.Scan(&temp.PerfID, &temp.MovieName, &temp.StartTime, &temp.EndTime, &temp.Halls, &temp.Money, &temp.ElseSpe)
		perfPlan = append(perfPlan, temp)
	}
	return perfPlan
}
func GetPerfPlanByID(perfID string) PerfPlan {
	row := utils.Db.QueryRow("select * from performance_plan where performance_id = ?", perfID)
	temp := PerfPlan{}
	row.Scan(&temp.PerfID, &temp.MovieName, &temp.StartTime, &temp.EndTime, &temp.Halls, &temp.Money, &temp.ElseSpe)
	rows := utils.Db.QueryRow("select img_path from movie where movie_name = ?", temp.MovieName)
	rows.Scan(&temp.ImgPath)
	return temp
}
func GetPlanByHall(hall string) []string {
	rows,_:=utils.Db.Query("select start_time,end_time from performance_plan where cinema_halls = ?",hall)
	plans := make([]string,0)
	for rows.Next() {
		temp1 := ""
		temp2 := ""
		rows.Scan(&temp1,&temp2)
		plans=append(plans,temp1,temp2)
	}
	return plans
}
func AddPlan(plan PerfPlan) error {
	_,err := utils.Db.Exec("insert into performance_plan (performance_id, movie_name, start_time, end_time, cinema_halls, money) value (?,?,?,?,?,?)", plan.PerfID,plan.MovieName,plan.StartTime,plan.EndTime,plan.Halls,plan.Money)
	if err != nil {
		return err
	}
	return nil
}
func GetAllPlans() []PerfPlan {
	rows,err:=utils.Db.Query("select performance_id,movie_name,start_time,cinema_halls  from performance_plan")
	if err != nil {
		log.Println(err)
	}
	plans := make([]PerfPlan,0)
	for rows.Next() {
		temp:= PerfPlan{}

		err=rows.Scan(&temp.PerfID,&temp.MovieName,&temp.StartTime,&temp.Halls)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(temp)
		plans = append(plans,temp)
	}
	return plans
}
func DeletePlan(perfID string)  {
	_,err:=utils.Db.Exec("delete from performance_plan where performance_id = ?",perfID)
	if err != nil {
		log.Println(err)
	}
	return
}