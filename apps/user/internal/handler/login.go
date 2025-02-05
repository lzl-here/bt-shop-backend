package handler

import (
	"context"
	"crypto/sha256"
	"fmt"

	"errors"

	"github.com/google/uuid"
	"github.com/lzl-here/bt-shop-backend/apps/user/internal/constant"
	"github.com/lzl-here/bt-shop-backend/apps/user/internal/domain/model"
	bizerr "github.com/lzl-here/bt-shop-backend/pkg/err"
	ugen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/user"
	"gorm.io/gorm"
)

func (h *UserHandler) NormalLogin(ctx context.Context, req *ugen.NormalLoginReq) (*ugen.NormalLoginRsp, error) {
	user, err := h.rep.GetUser(ctx, &model.User{Username: req.Username})
	if err != nil {
		return nil, err
	}
	// 用户不存在
	if user == nil {
		return nil, bizerr.ErrResourceNotFound
	}
	if !verifyPassword(req.Password, user.Password, user.Salt) {
		return nil, bizerr.ErrPasswordNotMatch
	}

	baseUser, err := user.ToGen()
	if err != nil {
		return nil, err
	}
	return &ugen.NormalLoginRsp{
		Code: 1,
		Msg:  "success",
		Data: &ugen.NormalLoginRsp_Data{
			User: baseUser,
		},
	}, nil
}

func (h *UserHandler) NormalRegister(ctx context.Context, req *ugen.NormalRegisterReq) (*ugen.NormalRegisterRsp, error) {
	if req.Password != req.ConfirmPassword {
		return nil, bizerr.ErrPasswordNotMatch
	}
	user, err := h.rep.GetUser(ctx, &model.User{Username: req.Username})
	// 忽略 "record not found" 错误
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	// 如果用户已存在
	if user != nil {
		return nil, bizerr.ErrUserAlreadyExists
	}
	// 生成盐值
	salt := generateSalt()
	// 对密码进行加盐哈希
	hashedPassword := hashPassword(req.Password, salt)

	// 注册用户，密码hash处理
	user, err = h.rep.CreateUser(ctx, &model.User{
		Username:     req.Username,
		Password:     hashedPassword,
		Salt:         salt,
		RegisterType: constant.RegisterTypeNormal,
		AvatarUrl:    "https://tse4-mm.cn.bing.net/th/id/OIP-C.Wm28iTeZUzxP_FOrlfqZWAHaHa?rs=1&pid=ImgDetMain", // 头像写死，暂时懒得加oss
	})
	if err != nil {
		return nil, err
	}
	return &ugen.NormalRegisterRsp{
		Code: 1,
		Msg:  "success",
		Data: &ugen.NormalRegisterRsp_Data{
			UserId: user.ID,
		},
	}, nil
}

// TODO 登出
func (h *UserHandler) Logout(ctx context.Context, req *ugen.LogoutReq) (*ugen.LogoutRsp, error) {
	return nil, nil
}

// 登录验证时的密码校验
func verifyPassword(inputPassword, storedHash, salt string) bool {
	hashedInput := hashPassword(inputPassword, salt)
	return hashedInput == storedHash
}

func generateSalt() string {
	return uuid.New().String()
}

func hashPassword(password, salt string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password+salt)))
}
