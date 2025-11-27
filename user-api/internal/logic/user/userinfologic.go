// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"user-api/internal/model"
	"user-api/internal/svc"
	"user-api/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *util.RestResponse, err error) {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return util.Error(util.ErrToken), nil
	}
	userModel := model.NewUsersModel(l.svcCtx.MySql)
	user, err := userModel.FindOne(l.ctx, uint64(userId))
	if err != nil && (errors.Is(err, model.ErrNotFound) || errors.Is(err, sql.ErrNoRows)) {
		return util.Error(util.ErrToken), nil
	}
	if err != nil {
		return util.ErrorWithMsg("查询用户失败"), nil
	}

	return util.Success(user), nil
}
