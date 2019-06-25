package handler

import (
	"context"
	"go-ChatRom/common"
	cfg "go-ChatRom/config"
	accountproto "go-ChatRom/service/account/proto"
	dbProto "go-ChatRom/service/dbproxy/proto"
	"go-ChatRom/util"
	"log"

	"github.com/micro/go-micro"
)

// User: 用于实现AccountService的对象
type User struct{}

var (
	dbCli dbProto.DBProxyService
)

func init() {
	service := micro.NewService()
	// 初始化， 解析命令行参数等
	service.Init()
	// 初始化一个dbproxy服务的客户端
	dbCli = dbProto.NewDBProxyService("go.micro.service.dbproxy", service.Client())
}

// Signup : 处理用户注册请求
func (u *User) Signup(ctx context.Context, req *accountproto.ReqSignup, res *accountproto.RespSignup) error {
	username := req.Username
	passwd := req.Password

	// 参数简单校验
	usernameLen := len(username)
	passedLen := len(passwd)

	if usernameLen < 5 || usernameLen > 12 || passedLen < 8 || passedLen > 20 {
		res.Code = common.StatusParamInvalid
		res.Message = "注册参数无效"
		return nil
	}

	// 对密码进行加盐及取Sha1值加密
	encPasswd := util.Sha1([]byte(passwd + cfg.PasswordSalt))
	// 将用户信息注册到用户表中
	dbResp, err := dbCli.ExecuteUserSignup(context.TODO(), &dbProto.ReqSignup{
		Username:  username,
		EncPasswd: encPasswd,
	})

	if err == nil && dbResp.Code == common.StatusMysqlDBOK {
		res.Code = common.StatusOK
		res.Message = "注册成功"
	} else if err == nil && dbResp.Code == common.StatusMysqlDBNoChange {
		res.Code = common.StatusRegisterFailed
		res.Message = "该用户已存在"
	} else {
		res.Code = common.StatusRegisterFailed
		res.Message = "未知原因，注册失败"
	}

	return err
}

// Signup : 处理用户注册请求
func (u *User) Signin(ctx context.Context, req *accountproto.ReqSignin, res *accountproto.RespSignin) error {
	username := req.Username
	passwd := req.Password

	// 参数简单校验
	usernameLen := len(username)
	passedLen := len(passwd)

	if usernameLen < 5 || usernameLen > 12 || passedLen < 8 || passedLen > 20 {
		res.Code = common.StatusParamInvalid
		res.Message = "登录参数无效"
		return nil
	}

	// 对密码进行加盐及取Sha1值加密
	encPasswd := util.Sha1([]byte(passwd + cfg.PasswordSalt))
	// 从用户表中查询用户的信息
	dbResp, err := dbCli.ExecuteUserSignin(context.TODO(), &dbProto.ReqSignin{
		Username:  username,
		EncPasswd: encPasswd,
	})

	res.Code = dbResp.Code
	res.Message = dbResp.Message
	if err != nil {
		return err
	}

	// 查找成功，更新token
	updateResp, err := dbCli.ExecuteUpdateToken(context.TODO(), &dbProto.ReqUpdateToken{
		Username: username,
		Token:    dbResp.Token,
	})

	res.Code = updateResp.Code
	res.Message = dbResp.Message
	if err != nil {
		log.Println(err.Error())
		return err
	}

	res.Token = dbResp.Token
	return nil
}

func (u *User) UpdateToken(ctx context.Context, req *accountproto.ReqUpdateToken, res *accountproto.RespUpdateToken) error {
	return nil
}

func (u *User) GetToken(ctx context.Context, req *accountproto.ReqGetToken, res *accountproto.RespGetToken) error {
	return nil
}
