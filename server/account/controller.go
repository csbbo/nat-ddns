package account

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"server/common"
	"server/database"
)

func login(c *gin.Context) {
	form := &LoginParam{}
	if err := c.ShouldBind(form); err != nil {
		common.ResponseError(c, err)
		return
	}

	user := struct {
		Username string
		Password string
	}{}

	collection := database.MongoDB.Collection("user")

	err := collection.FindOne(context.Background(), bson.M{"username": form.Username}).Decode(&user)
	if err != nil {
		common.ResponseError(c, "用户不存在")
		return
	}

	if user.Password != common.Hash(form.Password) {
		common.ResponseError(c, "密码错误")
		return
	}

	session := sessions.Default(c)
	session.Set("user", form.Username)
	session.Save()
	common.ResponseSuccess(c, nil)
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	session.Save()
	common.ResponseSuccess(c, nil)
}

func regist(c *gin.Context) {
	form := &RegistParam{}
	if err := c.ShouldBind(form); err != nil {
		common.ResponseError(c, err)
		return
	}

	collection := database.MongoDB.Collection("user")

	count, _ := collection.CountDocuments(context.Background(), bson.M{"username": form.Username})
	if count != 0 {
		common.ResponseError(c, "用户名已存在")
		return
	}

	data := bson.M{"username": form.Username, "password": common.Hash(form.Password)}
	res, _ := collection.InsertOne(context.Background(), data)
	common.ResponseSuccess(c, map[string]interface{}{"id": res.InsertedID})
}

func Setup(r *gin.Engine) {
	r.POST("/login", login)
	r.POST("/logout", logout)
	r.POST("/regist", regist)
}
