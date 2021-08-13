package modle

import (
	"fmt"
	"log"
	"web/Cinema/utils"
)

type Movie struct {
	MovieName   string
	MovieGrade  float64
	BoxOffice   float64
	MovieFlag   int
	TimeLong    string
	ImgPath     string
	WaitNum     int
	ReleaseTime string
	OffTime string
	PerfPlans []PerfPlan
}

func GetTheFifthMovieByGrade() []Movie {
	movies := make([]Movie, 0)
	rows, _ := utils.Db.Query("select * from movie order by grade desc ")
	for i := 0; i < 8 && rows.Next(); i++ {
		temp := Movie{}
		rows.Scan(&temp.MovieName, &temp.MovieGrade, &temp.ImgPath, &temp.BoxOffice, &temp.MovieFlag, &temp.WaitNum, &temp.TimeLong, &temp.ReleaseTime,&temp.OffTime)
		temp.PerfPlans = GetPerfPlanByName(temp.MovieName)
		movies = append(movies, temp)
	}
	return movies
}
func GetTheFifthMovieByWait() []Movie {
	movies := make([]Movie, 0)
	rows, _ := utils.Db.Query("select * from movie order by wait_num desc ")
	for i := 0; i < 10 && rows.Next(); i++ {
		temp := Movie{}
		rows.Scan(&temp.MovieName, &temp.MovieGrade, &temp.ImgPath, &temp.BoxOffice, &temp.MovieFlag, &temp.WaitNum, &temp.TimeLong, &temp.ReleaseTime, &temp.OffTime)
		temp.PerfPlans = GetPerfPlanByName(temp.MovieName)
		movies = append(movies, temp)
	}
	return movies
}
func GetTheFifthMovieByBox() []Movie {
	movies := make([]Movie, 0)
	rows, _ := utils.Db.Query("select * from movie order by box_office desc ")
	for i := 0; i < 5 && rows.Next(); i++ {
		temp := Movie{}
		rows.Scan(&temp.MovieName, &temp.MovieGrade, &temp.ImgPath, &temp.BoxOffice, &temp.MovieFlag, &temp.WaitNum, &temp.TimeLong, &temp.ReleaseTime,  &temp.OffTime)
		temp.PerfPlans = GetPerfPlanByName(temp.MovieName)
		movies = append(movies, temp)
	}
	return movies
}
func GetTheFifthMovieByFlag() []Movie {
	movies := make([]Movie, 0)
	rows, _ := utils.Db.Query("select * from movie order by movie_flag desc,release_time asc ")
	for i := 0; i < 8 && rows.Next(); i++ {
		temp := Movie{}
		rows.Scan(&temp.MovieName, &temp.MovieGrade, &temp.ImgPath, &temp.BoxOffice, &temp.MovieFlag, &temp.WaitNum, &temp.TimeLong, &temp.ReleaseTime, &temp.OffTime)
		temp.PerfPlans = GetPerfPlanByName(temp.MovieName)
		movies = append(movies, temp)
	}
	return movies
}

func GetMovieByName(movieName string) Movie {
	rows := utils.Db.QueryRow("select * from movie where movie_name = ?", movieName)
	temp := Movie{}
	err := rows.Scan(&temp.MovieName, &temp.MovieGrade, &temp.ImgPath, &temp.BoxOffice, &temp.MovieFlag, &temp.WaitNum, &temp.TimeLong, &temp.ReleaseTime, &temp.OffTime)
	if err != nil {
		fmt.Println(err)
	}
	temp.PerfPlans = GetPerfPlanByName(temp.MovieName)
	return temp
}

func AddMovie(movie Movie)  {
	_,err := utils.Db.Exec("insert into movie (movie_name, grade, img_path, box_office, movie_flag, wait_num, time_long, release_time, off_time) value (?,?,?,?,?,?,?,?,?)",movie.MovieName,movie.MovieGrade,movie.ImgPath,movie.BoxOffice,movie.MovieFlag,movie.WaitNum,movie.TimeLong,movie.ReleaseTime,movie.OffTime)
	if err != nil {
		log.Println(err)
	}
}
func GetAllMovies() []string {
	rows,err:=utils.Db.Query("select movie_name from movie")
	if err != nil {
		log.Println(err)
	}
	movies := make([]string,0)
	for rows.Next() {
		temp := ""
		rows.Scan(&temp)
		movies=append(movies,temp)
	}
	return movies
}
func GetTimeLongAndImgPath(movieName string) (long string) {
	row := utils.Db.QueryRow("select time_long from movie where movie_name = ?;",movieName)
	row.Scan(&long)
	return
}

func GetMovies(movieName string) []Movie {
	rows ,_:=utils.Db.Query("select * from movie where movie_name like '%"+movieName+"%'")
	movies := make([]Movie,0)
	for rows.Next() {
		temp := Movie{}
		rows.Scan(&temp.MovieName, &temp.MovieGrade, &temp.ImgPath, &temp.BoxOffice, &temp.MovieFlag, &temp.WaitNum, &temp.TimeLong, &temp.ReleaseTime, &temp.OffTime)
		movies=append(movies,temp)
	}
	return movies
}

