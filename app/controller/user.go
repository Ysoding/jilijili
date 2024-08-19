package controller

import (
	"github.com/Ysoding/jilijili/app/controller/errs"
	"github.com/Ysoding/jilijili/app/controller/payload"
	"github.com/Ysoding/jilijili/app/domain"
	"github.com/Ysoding/jilijili/app/service"
	"github.com/Ysoding/jilijili/pkg/ginx"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

type UserController struct {
	svc service.UserService
	log *zap.Logger

	emailRexExp    *regexp.Regexp
	passwordRexExp *regexp.Regexp
}

func NewUserController(svc service.UserService, log *zap.Logger) *UserController {
	return &UserController{svc: svc, log: log, emailRexExp: regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRexExp: regexp.MustCompile(passwordRegexPattern, regexp.None)}
}

func (ctl *UserController) HandlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (ctl *UserController) RegisterRoutes(r *gin.RouterGroup) {
	g := r.Group("/users")

	g.POST("", ginx.WrapBody(ctl.SignUp))
}

func (ctl *UserController) SignUp(ctx *gin.Context, req payload.SignUpReq) (ginx.Result, error) {
	isEmail, err := ctl.emailRexExp.MatchString(req.Email)
	if err != nil {
		return ginx.Result{
			Code: errs.UserInternalServerError,
			Msg:  "系统错误",
		}, nil
	}

	if !isEmail {
		return ginx.Result{
			Code: errs.UserInvalidInput,
			Msg:  "非法邮箱格式",
		}, nil
	}

	if req.Password != req.PasswordConfirm {
		return ginx.Result{
			Code: errs.UserInvalidInput,
			Msg:  "两次输入的密码不相等",
		}, nil
	}

	isPassword, err := ctl.passwordRexExp.MatchString(req.Password)
	if err != nil {
		return ginx.Result{
			Code: errs.UserInternalServerError,
			Msg:  "系统错误",
		}, err
	}

	if !isPassword {
		return ginx.Result{
			Code: errs.UserInvalidInput,
			Msg:  "密码必须包含至少一个字母、一个数字、一个特殊字符，并且长度至少为8个字符",
		}, nil
	}

	err = ctl.svc.SignUp(ctx.Request.Context(), domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	switch err {
	case nil:
		return ginx.Result{
			Msg: "OK",
		}, nil
	case service.ErrUniqueEmail:
		return ginx.Result{
			Code: errs.UserDuplicateEmail,
			Msg:  "邮箱冲突",
		}, nil
	default:
		return ginx.Result{
			Code: errs.UserInternalServerError,
			Msg:  "系统错误",
		}, err
	}
}
