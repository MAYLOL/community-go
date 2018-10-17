package waterpool

import (
	"community-go/redisTool"
	"fmt"
)

func Getwaterpool() {
	m, err := redisTool.RedisDB.ZRange("data:articles2",1,5).Result()
	if err != nil {
		fmt.Println("redis查询错误")
		return
	}
	println(m)
	//res, err := redisTools.RedisDB.ZRange("data:articles2",int64((in.Page-1)*in.Count),int64((in.Page * in.Count)-1)).Result()


}