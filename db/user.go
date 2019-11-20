package db

import (
	mydb "filestore-server/db/mysql"
	"fmt"
)

//UserSignup：通过用户名及密码完成user表的注册操作
func UserSignup(username, passwd string) bool {
	stmt, err := mydb.DBConn().Prepare("INSERT IGNORE INTO tbl_user(user_name,user_pwd) VALUES(?,?)")
	if err != nil {
		fmt.Println("Failed to insert,err:" + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, passwd)
	if err != nil {
		fmt.Println("Failed to insert,err:" + err.Error())
		return false
	}
	if rowsAffected, err := ret.RowsAffected(); nil == err && rowsAffected > 0 {
		return true
	}
	return false
}

//UserSignin:判断密码是否一致
func UserSignin(username, enc_pwd string) bool {
	stmt, err := mydb.DBConn().Prepare("SELECT * FROM tbl_user WHERE user_name=? LIMIT 1")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else if rows == nil {
		fmt.Println("username not found:" + username)
		return false
	} else {
		pRows := mydb.ParseRows(rows)
		if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == enc_pwd {
			return true
		}
	}
	return false
}

//UpdateToken:刷新用户登录token
func UpdateToken(username, token string) bool {
	stmt, err := mydb.DBConn().Prepare("REPLACE INTO tbl_user_token(user_name,user_token) VALUES (?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, token)
	if err != nil {
		fmt.Println(err.Error())
	}
	return true
}

type User struct {
	Username     string
	Email        string
	Phone        string
	SignupAt     string
	LastActiveAt string
	Status       int
}

func GetUserInfo(username string) (User, error) {
	user := User{}
	stmt, err := mydb.DBConn().Prepare("SELECT user_name,signup_at FROM tbl_user WHERE user_name = ? LIMIT 1")
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&user.Username, &user.SignupAt)
	return user, err
}
