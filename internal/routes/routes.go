package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/database"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/handler"
)

func Run(port string, db *sql.DB) {
	var (
		app    = fiber.New()
		dbase  = database.NewDatabase(db)
		userdb = database.NewUserDB(dbase)
		postdb = database.NewPostDB(dbase)
		hndler = handler.NewHandler(userdb, postdb)
		v1     = app.Group("/api/v1")
	)

	app.Get("/hello", hndler.Hello)
	// v1 api -> user api
	v1.Post("/user", hndler.UserInsert)

	//v1 api -> post api
	v1.Post("/post", hndler.PostInsert)

	app.Listen(port)
}
