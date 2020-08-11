package frp

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/common"
	"server/database"
)

func frpCreate(c *gin.Context) {
	form := &CreateFRPParm{}
	if err := common.Check(c, form, true); err != nil {
		common.ResponseError(c, err)
		return
	}

	user := struct {
		ID       primitive.ObjectID `bson:"_id"`
		Username string             `bson:"username"`
		Password string             `bson:"password"`
	}{}

	collection := database.MongoDB.Collection("frp")
	type FrpItem struct {
		LocalPort int
		UserID    primitive.ObjectID
	}
	frpItem := FrpItem{form.LocalPort, user.ID}
	result, _ := collection.InsertOne(context.Background(), frpItem)
	common.ResponseSuccess(c, map[string]interface{}{"id": result.InsertedID})
}

func frpDelete(c *gin.Context) {
	if err := common.Check(c, nil, true); err != nil {
		common.ResponseError(c, err)
		return
	}
	id := c.Query("id")

	collection := database.MongoDB.Collection("frp")
	objID, _ := primitive.ObjectIDFromHex(id)
	if count, _ := collection.CountDocuments(context.Background(), bson.M{"_id": objID}); count == 0 {
		common.ResponseError(c, "记录不存在,无法删除")
	}
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		common.ResponseError(c, "删除失败")
		return
	}
	common.ResponseSuccess(c, nil)
}

func frpList(c *gin.Context) {
	if err := common.Check(c, nil, true); err != nil {
		common.ResponseError(c, err)
		return
	}
	username := sessions.Default(c).Get("user")

	user := struct {
		ID       primitive.ObjectID `bson:"_id"`
		Username string
		Password string
	}{}
	collection := database.MongoDB.Collection("user")
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		common.ResponseError(c, "用户不存在")
		return
	}

	collection = database.MongoDB.Collection("frp")
	cur, err := collection.Find(context.Background(), bson.M{"userid": user.ID})
	if err != nil {
		common.ResponseError(c, "查询异常")
		return
	}

	type Frp struct {
		ID        primitive.ObjectID `bson:"_id"`
		LocalPort int
		UserID    string 			 `bson:"userid"`
	}
	var frp Frp
	var results []Frp
	for cur.Next(context.Background()) {
		err = cur.Decode(&frp)
		if err != nil {
			common.ResponseError(c, "解析失败")
			return
		}
		results = append(results, frp)
	}
	common.ResponseSuccess(c, results)
}

func Setup(r *gin.Engine) {
	r.POST("/frp/create", frpCreate)
	r.DELETE("/frp/delete", frpDelete)
	r.GET("/frp/list", frpList)
}
