package http

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	apiGroup := r.Group("/api/activity/prop")
	{
		apiGroup.Any("/ping", ping)
		apiGroup.Any("/get_prop_amount", getPropAmount) // 查询道具数量
		apiGroup.Any("/use_prop", useProp)              // 使用道具
		apiGroup.Any("/grant_prop", grantProp)          // 发放道具
		apiGroup.Any("/exchange_prop", exchangeProp)    // 兑换道具(多换一)
	}
	_ = r.Run(":8082")
}
