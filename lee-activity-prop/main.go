package main

import (
	"lee-activity-framework/lee-activity-prop/dao"
	"lee-activity-framework/lee-activity-prop/server/http"
)

func main() {
	// MySQL
	dao.InitMySQL()
	// Http
	http.InitRouter()
}
