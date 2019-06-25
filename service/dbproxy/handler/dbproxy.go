package handler

import (
	"context"
	"fmt"
	"go-ChatRom/common"
	mydb "go-ChatRom/service/dbproxy/db/mysql"
	proto "go-ChatRom/service/dbproxy/proto"
	util "go-ChatRom/util"
	"log"
)

// DBProxy : DBProxy结构体
type DBProxy struct{}

func (db *DBProxy) ExecuteUserSignup(ctx context.Context, req *proto.ReqSignup, res *proto.RespSignup) error {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_user (`user_name`,`user_pwd`) values (?,?)")
	if err != nil {
		log.Println("Failed to insert, err:" + err.Error())
		res.Code = common.StatusMysqlDBError
		res.Message = err.Error()
		return err
	}
	defer stmt.Close()

	ret, err := stmt.Exec(req.Username, req.EncPasswd)

	if err != nil {
		log.Println("Failed to insert, err:" + err.Error())
		res.Code = common.StatusMysqlDBError
		res.Message = err.Error()
		return err
	}
	if rowsAffected, err := ret.RowsAffected(); nil == err && rowsAffected > 0 {
		res.Code = common.StatusMysqlDBOK
		return nil
	}

	res.Code = common.StatusMysqlDBNoChange
	res.Message = "无记录更新"
	return nil
}

func (db *DBProxy) ExecuteUserSignin(ctx context.Context, req *proto.ReqSignin, res *proto.RespSignin) error {
	stmt, err := mydb.DBConn().Prepare("select * from tbl_user where user_name=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(req.Username)
	if err != nil {
		fmt.Println(err.Error())
		return err
	} else if rows == nil {
		fmt.Println("username not found: " + req.Username)
		return err
	}

	pRows := mydb.ParseRows(rows)
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == req.EncPasswd {
		// 成功匹配
		res.Code = common.StatusOK
		res.Token = util.GenToken(req.Username)
		res.Message = "登录成功"
		return nil
	}

	res.Code = common.StatusUserSigninFailed
	res.Message = "用户名或密码有误"
	return nil
}

func (db *DBProxy) ExecuteUpdateToken(ctx context.Context, req *proto.ReqUpdateToken, res *proto.RespUpdateToken) error {
	stmt, err := mydb.DBConn().Prepare(
		"replace into tbl_user_token (`user_name`,`user_token`) values (?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Username, req.Token)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	res.Code = common.StatusOK
	return nil
}

func (db *DBProxy) ExecuteGetToken(ctx context.Context, req *proto.ReqGetToken, res *proto.RespGetToken) error {
	stmt, err := mydb.DBConn().Prepare("select * from tbl_user_token where user_name=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(req.Username)
	if err != nil {
		fmt.Println(err.Error())
		return err
	} else if rows == nil {
		fmt.Println("username not found: " + req.Username)
		return err
	}

	pRows := mydb.ParseRows(rows)
	if len(pRows) == 1 {
		// 成功匹配
		res.Code = common.StatusOK
		res.Token = string(pRows[0]["user_token"].([]byte))
		return nil
	}

	res.Code = common.StatusGetTokenFailed
	res.Message = "token校验失败"
	return nil
}
