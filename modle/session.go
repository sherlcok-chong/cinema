package modle

import "web/Cinema/utils"

type Session struct {
	SessionID string
	UserID    string
	UserName  string
}

func AddSession(session Session) error {
	_,err := utils.Db.Exec("insert into session (session_id, user_id, user_name) value (?,?,?)",session.SessionID,session.UserID,session.UserName)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSession(sessionID string) error {
	_,err:= utils.Db.Exec("delete from session where session_id=?",sessionID)
	if err != nil {
		return err
	}
	return nil
}

func GetSessionById(SessId string) (Session,error){
	row :=utils.Db.QueryRow("select session_id,user_name,user_id from session where session_id = ?",SessId)
	var sess Session
	err:=row.Scan(&sess.SessionID,&sess.UserName,&sess.UserID)
	return sess, err
}