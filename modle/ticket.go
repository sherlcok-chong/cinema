package modle

import (
	"errors"
	"fmt"
	"sync"
	"web/Cinema/utils"
)

type Ticket struct {
	OrderID   string
	UserID    string
	PerfID    string
	MovieName string
	Time      string
	Row       int64
	Line      int64
	Hall      string
	Flag      int
	Money     string
	ImgPath   string
}

var mutex sync.Mutex

func AddTicket(tickets []Ticket) error {
	sits := make(map[string]string, 0)
	rows, _ := utils.Db.Query("select row,line from ticket")
	for rows.Next() {
		row := ""
		line := ""
		rows.Scan(&row,&line)
		sits[row]=line
	}
	for _, v := range tickets {
		mutex.Lock()
		if string(v.Line)==sits[string(v.Row)]{
			return errors.New("0927")
		}

		_, err := utils.Db.Exec("insert into ticket (order_id,userid, per_id, movie_name, time, row, line, hall,flag,money,img_path) value (?,?,?,?,?,?,?,?,?,?,?)", v.OrderID, v.UserID, v.PerfID, v.MovieName, v.Time, v.Row, v.Line, v.Hall, v.Flag, v.Money, v.ImgPath)

		if err != nil {
			fmt.Println(err)
			return err
		}
		mutex.Unlock()
	}
	return nil
}
func GetTicketByFlag() []Ticket {
	tickets := make([]Ticket, 0)
	rows, _ := utils.Db.Query("select * from ticket where flag = ?", 0)
	for rows.Next() {
		v := Ticket{}
		rows.Scan(&v.OrderID, &v.UserID, &v.PerfID, &v.MovieName, &v.Time, &v.Row, &v.Line, &v.Hall, &v.Flag, &v.Money, &v.ImgPath)
		tickets = append(tickets, v)
	}
	return tickets
}
func GetTicketByUserID(userID string) []Ticket {
	tickets := make([]Ticket, 0)
	rows, _ := utils.Db.Query("select * from ticket where userid = ?", userID)
	for rows.Next() {
		v := Ticket{}
		rows.Scan(&v.OrderID, &v.UserID, &v.PerfID, &v.MovieName, &v.Time, &v.Row, &v.Line, &v.Hall, &v.Flag, &v.Money, &v.ImgPath)
		v.Flag = utils.TestTime(v.Time)
		tickets = append(tickets, v)
	}
	return tickets
}
func GetTicketByOrderID(orderID string) []Ticket {
	tickets := make([]Ticket, 0)
	rows, _ := utils.Db.Query("select * from ticket where order_id = ?", orderID)
	for rows.Next() {
		v := Ticket{}
		rows.Scan(&v.OrderID, &v.UserID, &v.PerfID, &v.MovieName, &v.Time, &v.Row, &v.Line, &v.Hall, &v.Flag, &v.Money, &v.ImgPath)
		tickets = append(tickets, v)
	}
	return tickets
}

func DeleteTicByFlag() {
	utils.Db.Exec("delete from ticket where flag=?", 0)
}
func DeleteTicByOrderID(orderID string) {
	utils.Db.Exec("delete from ticket where order_id=?", orderID)
}
func UpdateTicByFlag() {
	utils.Db.Exec("update  ticket set flag=?  where flag=?", 1, 0)
}
