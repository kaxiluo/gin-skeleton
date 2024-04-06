package demo

import (
	"context"
	"fmt"
	"gin-api/app/common/request"
	"gin-api/app/common/response"
	"gin-api/app/models"
	"gin-api/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Test1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func Test2(c *gin.Context) {
	global.App.Log.Info("test2 is logged")

	var form request.DemoFoo
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	response.Success(c, gin.H{
		"now_time": time.Now(),
		"got_foo":  form.Foo,
	})
}

func TestDb(c *gin.Context) {
	db := global.App.DB
	// 执行原生 SQL 查询
	rows, err := db.Raw("SELECT * FROM users").Rows()
	if err != nil {
		panic(err)
	}
	// 遍历查询结果
	var data []models.User
	for rows.Next() {
		var item models.User
		err := db.ScanRows(rows, &item)
		if err != nil {
			panic(err)
		}
		data = append(data, item)
	}
	rows.Close()
	response.Success(c, data)
}

func TestRedis(c *gin.Context) {
	ctx := context.Background()
	rdb := global.App.Redis
	err := rdb.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", val)
}
