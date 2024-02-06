package main

import (
	"lee-activity-framework/lee-activity-config/dao"
	"lee-activity-framework/lee-activity-config/server/http"
)

func main() {
	// MySQL
	dao.InitMySQL()
	// Http
	http.InitRouter()
}
