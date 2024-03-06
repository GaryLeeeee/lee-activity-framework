package main

import (
	"lee-activity-framework/lee-activity-lottery/dao"
	"lee-activity-framework/lee-activity-lottery/server/http"
)

func main() {
	// MySQL
	dao.InitMySQL()
	// Http
	http.InitRouter()
}
