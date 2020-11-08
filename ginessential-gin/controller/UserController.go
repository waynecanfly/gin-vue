package controller

import (
	"ginessential/common"
	"ginessential/dto"
	"ginessential/model"
	"ginessential/response"
	"ginessential/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	DB := common.GetDB()

	requestUser := model.User{}
	//gin提供了一种方式Bind
	c.Bind(&requestUser)
	//获取参数
	name := requestUser.Name
	tel := requestUser.Tel
	pwd := requestUser.Pwd

	//数据验证
	if len(tel) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号码必须11位")
	}
	if len(pwd) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}
	//判断手机号是否存在
	if isTelephoneExist(DB, tel) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}
	log.Println(name, pwd, tel)

	//创建用户
	//为用户密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name: name,
		Tel:  tel,
		Pwd:  string(hasedPassword),
	}
	DB.Create(&newUser)

	response.Success(c, nil, "用户注册成功")
}

func Login(c *gin.Context) {
	DB := common.GetDB()
	//获取参数
	tel := c.PostForm("tel")
	pwd := c.PostForm("pwd")
	//数据验证
	if len(tel) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "The length of the tel-number must be 11",
		})
	}
	if len(pwd) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能少于6位",
		})
		return
	}

	//判断手机号是否存在
	var user model.User
	DB.Where("tel = ?", tel).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "User does not exits",
		})
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(pwd)); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 400,
			"msg":  "wrong password",
		})
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		log.Printf("token generate error :%v", err)
		return
	}

	//返回结果
	response.Success(c, gin.H{"token": token}, "登录成功")

}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"code": 20,
		"data": gin.H{
			"user": dto.ToUserDto(user.(model.User)),
		},
	})
}

func isTelephoneExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("tel = ?", tel).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
