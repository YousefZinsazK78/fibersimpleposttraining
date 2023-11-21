package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/database"
	"github.com/yousefzinsazk78/fiber_post_second_version/internal/handler"
)

//todo : add logger to fiber

func Run(port string, db *sql.DB) {
	var (
		app = fiber.New(fiber.Config{
			ErrorHandler: handler.ErrorHandler,
		})
		dbase  = database.NewDatabase(db)
		userdb = database.NewUserDB(dbase)
		postdb = database.NewPostDB(dbase)
		hndler = handler.NewHandler(userdb, postdb)
		v1     = app.Group("/api/v1")
	)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/hello", hndler.Hello)
	// v1 api -> user api
	v1.Post("/user", hndler.UserInsert)

	//v1 api -> post api
	v1.Post("/post", hndler.PostInsert)
	v1.Get("/posts", hndler.GetPosts)
	v1.Get("/post/id/:id", hndler.GetPostByID)
	v1.Get("/post/title/:title", hndler.GetPostByTitle)
	v1.Put("/post/update/", hndler.PutPost)
	v1.Delete("/post/delete/:id", hndler.DeletePost)

	app.Listen(port)
}
