package main

import (
	"lee-activity-framework/lee-activity-rank/dao"
	"lee-activity-framework/lee-activity-rank/server/http"
)

func main() {
	// MySQL
	dao.InitMySQL()
	// Http
	http.InitRouter()
}
