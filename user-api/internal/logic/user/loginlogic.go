// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"

	"user-api/internal/svc"
	"user-api/internal/types"
	"user-api/internal/model"
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

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	userModel := model.NewUsersModel(l.svcCtx.MySql)
	user, err := userModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		l.Logger.Error("查询用户失败", err)
		return &types.LoginResp{
			Code: 500,
			Msg: "用户不存在",
		}, nil
	}
	// 验证密码
	if !util.CheckPassword(user.Password, req.Password) {
		l.Logger.Error("密码错误")
		return &types.LoginResp{
			Code: 200,
			Msg: "密码错误",
		}, nil
	}
	jwtUtil := &util.JWTUtil{
		JWTConfig: l.svcCtx.Config.JWTConfig,
	}
	token, err := jwtUtil.GenerateToken(user.Id, user.Username)
	if err != nil {
		l.Logger.Error("生成token失败", err)
		return &types.LoginResp{
			Code: 500,
			Msg: "查询用户失败",
		}, nil
	}
	return &types.LoginResp{
		Code: 200,
		Msg: "登录成功",
		Token: token,
	}, nil
}
