package main

import (
			"github.com/kataras/iris"
	"community-go/controller/follow/content"
	"community-go/controller/follow/replys"
		"community-go/utility"
	"community-go/controller/follow/follow"
		"community-go/controller/follow/comments"
	"github.com/go-redis/redis"
	"fmt"
	"community-go/waterpool"
	"community-go/redisTool"
)

//func init() {
//	//链接本地数据库
//	db, _ := sql.Open("mysql", "root:Anyuechao1@tcp(localhost:3306)/merlot_test?charset=utf8")
//	//db, _ := sql.Open("mysql", "maylol_test:maylol_test@tcp(39.104.82.179:3306)/merlot_test?charset=utf8")
//	//db, _ := sql.Open(config.DBConfig.Dialect, config.DBConfig.URL)
//	err := db.Ping()
//	if err != nil {
//		fmt.Println("数据库链接失败", err.Error())
//	} else {
//		fmt.Println("数据库链接成功")
//	}
//	utility.DB = db
//}

var client = redis.NewClient(&redis.Options{
	Addr: "web-1.maylol.com:19000",
	Password:"",
	DB:0,
})

func init() {
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("redis链接失败")
	} else {
		fmt.Println("redis链接成功")
		redisTool.RedisDB = client
	}
}

func main() {
	app := iris.Default()
	//crs := cors.New(cors.Options{
	//	AllowedOrigins: []string("*"),
	//	AllowCredentials: true,
	//})

	v1 := app.Party("/api/v1").AllowMethods(iris.MethodOptions)
	{
		//887
		v1.Get("/comments/{contentID}", func(ctx iris.Context) {
			comments.GetComments(ctx)
		})
		v1.Get("/contents/{pageNum}", func(ctx iris.Context) {
			content.GetContentList(false, ctx)
		})
		//1001
		v1.Get("/replys/{commentID}",  func(ctx iris.Context) {
			replys.GetReplys(ctx)
		})
		v1.Get("/follow/{userID}/{followID}", func(ctx iris.Context) {
			follow.GetFollow(ctx)
		})
		//?userID=119&followID=23
		v1.Get("/waterpool", func(ctx iris.Context) {
			waterpool.Getwaterpool()
		})
	}
	//app.Run(iris.Addr(":8082"))


	app.Run(iris.Addr(":" + "8082"))
	defer utility.DB.Close()
}

//func followAction(app *iris.Application) {
//
//	//followList.FollowAction(app)
//
//	app.Handle("GET", "followList", func(ctx iris.Context) {
//		//users, userErr := users.GetUser(ctx)
//		//follow, followErr := follow.GetFollow(ctx)
//		//content, contentErr := content.GetContent(ctx)
//		//comment, commentErr := comments.GetComments(content.ContentId)
//		comment, commentErr := comments.GetComments(887)
//		contentList := make([]interface{}, 0)
//		//commentsList := make([]comments.Comments, 0)
//		//commentsList = append(commentsList, comment)
//		//content.CommentList = commentsList
//		contentMap := map[string]interface{}{
//			//"follow":  follow,
//			//"content": content,
//			//"users":   users,
//			"comments" : comment,
//		}
//		contentList = append(contentList, contentMap)
//		//contentDic := map[string]interface{}{
//		//	"contentList": contentList,
//		//}
//		//fmt.Println(contentErr)
//		//fmt.Println(commentErr)
//		//fmt.Println(userErr)
//		responseFunc(ctx, contentList, commentErr)
//
//	})
//}

//func followAction(app *iris.Application) {
//	app.Handle("GET", "follow", func(ctx iris.Context) {
//		//ctx.WriteString("string")
//		res, err := follow.GetFollow(ctx)
//		res1, _ := banner.GetBanner(ctx)
//		group := map[string]interface{} {
//			"follow" : res["followList"],
//			"bannner" : res1["bannerlist"],
//		}
//		responseFunc(ctx, group, err)
//	})
//}