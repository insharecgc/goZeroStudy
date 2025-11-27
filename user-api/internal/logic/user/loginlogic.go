// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"
	"time"

	"user-api/internal/model"
	"user-api/internal/svc"
	"user-api/internal/types"
	"user-api/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *util.RestResponse, err error) {
	// todo: add your logic here and delete this line
	userModel := model.NewUsersModel(l.svcCtx.MySql)
	user, err := userModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		l.Logger.Error("查询用户失败", err)
		return util.ErrorWithCodeMsg(500, "用户不存在"), nil
	}
	// 验证密码
	if !util.CheckPassword(user.Password, req.Password) {
		l.Logger.Error("密码错误")
		return util.Error(util.ErrInvalidPass), nil
	}
	jwtConfig := l.svcCtx.Config.JWTConfig
	expiration, err := time.ParseDuration(jwtConfig.ExpirationTime)
	if err != nil {
		expiration = time.Hour * 24 // 默认 24 小时
	}
	// 用gozero框架，需要用util.GetJwtToken方式生成token
	token, err := util.GetJwtToken(jwtConfig.SecretKey, time.Now().Unix(), int64(expiration.Seconds()), user.Id)
	// jwtUtil := &util.JWTUtil{
	// 	JWTConfig: l.svcCtx.Config.JWTConfig,
	// }
	// token, err := jwtUtil.GenerateToken(user.Id, user.Username)
	if err != nil {
		l.Logger.Error("生成token失败", err)
		return util.ErrorWithMsg("查询用户失败"), nil
	}
	return util.Success(token), nil
}
