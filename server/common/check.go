package common

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"server/database"
)

func Check(c *gin.Context, formData interface{}, loginRequire bool) interface{}{
	if loginRequire == true {
		collection := database.MongoDB.Collection("user")
		session := sessions.Default(c)
		username := session.Get("user")
		count, _ := collection.CountDocuments(context.Background(), bson.D{{"username", username}})
		if count == 0 {
			return "用户未登录"
		}
	}
	if formData != nil {
		err := c.ShouldBind(formData)
		return err
	}
	return nil
}
