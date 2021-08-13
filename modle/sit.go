package modle

import (
	"fmt"
	"web/Cinema/utils"
)

type Sit struct {
	PerfID   string
	Perf     PerfPlan
	SitOne   []Seat
	SitTwo   []Seat
	SitThree []Seat
	SitFour  []Seat
	SitFive  []Seat
	SitSix   []Seat
	SitSeven []Seat
}
type Seat struct {
	Flag    string
	Row     int
	Line    int
	AllLine int
}

func GetSitByID(perfID string) Sit {
	sit := Sit{}
	var r, l int
	var f string
	row := utils.Db.QueryRow("select all_line from sit where performance_id = ?", perfID)
	err := row.Scan(&l)
	if err != nil {
		fmt.Println(err)
	}

	sit.PerfID = perfID
	sit.Perf = GetPerfPlanByID(perfID)
	sit.SitOne = make([]Seat, l)
	sit.SitTwo = make([]Seat, l)
	sit.SitThree = make([]Seat, l)
	sit.SitFour = make([]Seat, l)
	sit.SitFive = make([]Seat, l)
	sit.SitSix = make([]Seat, l)
	sit.SitSeven = make([]Seat, l)
	rows, _ := utils.Db.Query("select rows,flag from sit  where performance_id = ? order by rows asc ,line asc ", perfID)
	for i := 0; i < l && rows.Next(); i++ {
		rows.Scan(&r, &f)
		sit.SitOne[i].Flag = f
		sit.SitOne[i].Row = r
		sit.SitOne[i].Line = i
	}
	for i := 0; i < l && rows.Next(); i++ {
		rows.Scan(&r, &f)
		sit.SitTwo[i].Flag = f
		sit.SitTwo[i].Row = r
		sit.SitTwo[i].Line = i
	}
	for i := 0; i < l && rows.Next(); i++ {
		rows.Scan(&r, &f)
		sit.SitThree[i].Flag = f
		sit.SitThree[i].Row = r
		sit.SitThree[i].Line = i
	}
	for i := 0; i < l && rows.Next(); i++ {
		rows.Scan(&r, &f)
		sit.SitFour[i].Flag = f
		sit.SitFour[i].Row = r
		sit.SitFour[i].Line = i
	}
	for i := 0; i < l && rows.Next(); i++ {
		rows.Scan(&r, &f)
		sit.SitFive[i].Flag = f
		sit.SitFive[i].Row = r
		sit.SitFive[i].Line = i
	}
	for i := 0; i < l && rows.Next(); i++ {
		rows.Scan(&r, &f)
		sit.SitSix[i].Flag = f
		sit.SitSix[i].Row = r
		sit.SitSix[i].Line = i
	}
	if rows.Next() == false {
		sit.SitSeven = nil
	} else {
		for i := 0; i < l && rows.Next(); i++ {
			rows.Scan(&r, &f)
			sit.SitSeven[i].Flag = f
			sit.SitSeven[i].Row = r
			sit.SitSeven[i].Line = i
		}
	}

	return sit
}
func UpdateSit(tickets []Ticket, flag string) {
	for _, v := range tickets {
		utils.Db.Exec("update sit set flag = ? where rows = ? and line=? and performance_id = ? ", flag, v.Row, v.Line+1,v.PerfID)
	}
}

func GetHallSit(hall string) []Seat {
	rows,_:=utils.Db.Query("select * from hall"+hall)
	sits := make([]Seat,0)
	for rows.Next() {
		temp :=Seat{}
		rows.Scan(&temp.Row,&temp.Line,&temp.Flag,&temp.AllLine)
		sits=append(sits,temp)
	}

	return sits
}

func AddSit(perfID,hall string)  {
	sits := GetHallSit(hall)
	for _,v:=range sits{
		utils.Db.Exec("insert into sit (performance_id, rows, line,flag, all_line) value (?,?,?,?,?)",perfID,v.Row,v.Line,v.Flag,v.AllLine)
	}
}
func QueryHas(perfID string) bool {
	rows,_ := utils.Db.Query("select rows from sit where performance_id=? and flag = ?",perfID,"sold")
	if rows.Next(){
		return false
	}

	return true
}