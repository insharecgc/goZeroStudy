// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"
	"database/sql"
	"time"

	"user-api/internal/model"
	"user-api/internal/svc"
	"user-api/internal/types"
	"user-api/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *util.RestResponse, err error) {
	// todo: add your logic here and delete this line
	userModel := model.NewUsersModel(l.svcCtx.MySql)
	user, err := userModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error("查询失败", err)
		return util.ErrorWithCodeMsg(500, "查询失败"), nil
	}
	if user != nil {
		l.Logger.Error("用户已存在")
		return util.Error(util.ErrUserExist), nil
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		l.Logger.Error("密码加密失败", err)
		return util.ErrorWithCodeMsg(500, "密码加密失败"), nil
	}
	newUser := &model.Users{
		Username:  req.Username,
		Password:  hashedPassword,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}
	_, err = userModel.Insert(l.ctx, newUser)
	if err != nil {
		l.Logger.Error("创建用户失败", err)
		return util.ErrorWithCodeMsg(500, "创建用户失败"), nil
	}

	return util.Success(nil), nil
}
