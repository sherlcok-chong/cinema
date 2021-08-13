package modle

import (
	"fmt"
	"web/Cinema/utils"
)

type User struct {
	UserID    string
	UserName  string
	Password  string
	IsManager int
	Email     string
}

// CheckUserName 查询用户名和密码
func CheckUserName(username string) (User ,error){
	fmt.Println("数据库了",username)
	row := utils.Db.QueryRow("select id,username,password,email from users where username=? ",username)

	user := User{}
	err:=row.Scan(&user.UserID,&user.UserName,&user.Password,&user.Email)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("数据库了",user)
	return user ,nil
}

func CheckPassword(username string, password string)(User, bool ){
	user,err:= CheckUserName(username)
	if err != nil {
		panic(err)
	}
	if user.Password!=password||password=="" {
		return user,false
	}
	return user,true
}
func InsertUser(user User) bool {
	_,err := utils.Db.Exec("insert into users (id,username,password,email,is_manager) value (?,?,?,?,?) ",user.UserID,user.UserName,user.Password,user.Email,0)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func CheckUserState(username string) bool {
	row := utils.Db.QueryRow("select is_manager from users where username=? ",username)
	state := 0
	row.Scan(&state)
	if state==0{
		return false
	}
	return true
}