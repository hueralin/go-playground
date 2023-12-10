package router

import (
	"github.com/gin-gonic/gin"
	"go-playground/gorm/model"
	"net/http"
	"strconv"
)

func InitUserRouter(r *gin.Engine) {
	ur := r.Group("/user")
	ur.POST("", CreateUser)
	ur.POST("/list", GetUserList)
	ur.GET("/:id", GetUserById)
	ur.DELETE("/:id", DeleteUserById)
	ur.POST("/age", UpdateUserAgeById)
}

func CreateUser(c *gin.Context) {
	var userDto model.CreateUserDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user := model.User{
		Username: userDto.Username,
		Password: userDto.Password,
		Age:      userDto.Age,
	}
	err := model.CreateUser(&user)
	if err != nil {
		c.JSON(200, gin.H{"code": -1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "", "data": user})
}

func GetUserById(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	user, err := model.GetUserById(userId)
	if err != nil {
		c.JSON(200, gin.H{"code": -1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "", "data": user})
}

func GetUserList(c *gin.Context) {
	var param model.GetUserListDto
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	users, err := model.GetUserList(param)
	if err != nil {
		c.JSON(200, gin.H{"code": -1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "",
		"data": gin.H{
			"data":  users,
			"total": len(users),
		},
	})
}

func UpdateUserAgeById(c *gin.Context) {
	var updateUserAgeDto model.UpdateUserAgeDto
	if err := c.ShouldBindJSON(&updateUserAgeDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "msg": err.Error(), "data": nil})
		return
	}
	if err := model.UpdateUserAgeById(updateUserAgeDto.Id, updateUserAgeDto.Age); err != nil {
		c.JSON(200, gin.H{"code": -1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "", "data": nil})
}

func DeleteUserById(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUserById(userId); err != nil {
		c.JSON(200, gin.H{"code": -1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "", "data": nil})
}
