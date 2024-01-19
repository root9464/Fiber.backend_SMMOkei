package main

import (
	"log"
	_"fmt"
	controller "tmp/database/controller"
	routes "tmp/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func AllRoutes(app *fiber.App) {
	app.Get("/hello", routes.Hello)
	app.Post("/createpost", routes.AddPost)
	app.Get("/getpost/:title", routes.GetPostByID)
	app.Get("/delpost/:title", routes.DeletePostByTitle)
}

func main() {

	controller.ConnectToDB()

	app := fiber.New()
	AllRoutes(app)
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))

	// if err != nil {
	// 	log.Fatalf("Ошибка подключения к бд: %v", err)
	// }
	// fmt.Println("Успешное подключение к бд")
	// db.DB.AutoMigrate(models.Post{})
	//создание юзера
	// user1 := user.User{Name: "ff2", Password: "1233", IsAdmin: true}
	// err = user.CreateUser(db.DB,&user1)
	// if err != nil {
	// 	log.Fatalf("ошибка при создании юзера: %v", err)
	// }
	// fmt.Println("юзер успешно создан")
	//создание таблицы постов
	// if err := post.CreateTablePosts(db.DB); err != nil {
	// 	log.Fatalf("ошибка создания таблицы: %v", err)
	// }
	// fmt.Println("таблица успешно создано")
	//создание поста
	// post1 := conn.Post{Title: "title", Сontent: "content"}
	// err = conn.CreatePost(db.DB, &post1)
	// if err != nil {
	// 	log.Fatalf("ошибка при создании поста: %v", err)
	// }
	// fmt.Println("пост успешно создан")

	// user1, err := conn.GetUserByID(db.DB, 3)
	// if err != nil {
	// 	log.Fatalf("Ошибка при получении пользователя: %v", err)
	// }
	// fmt.Printf("Получен пользователь: %+v\n", user1)

	// post1, err := post.GetPostByID(db.DB, 1)
	// if err != nil {
	// 	log.Fatalf("Ошибка при получении пользователя: %v", err)
	// }
	// fmt.Printf("Получен пользователь: %+v\n", post1)

		// Вызов функции GetPostByID из пакета post
	// postInstance, err := routes.AllPosts(c) // Здесь предполагается, что вы хотите получить пост с ID 1
	// if err != nil {
	// 	log.Fatalf("Ошибка при получении поста: %v", err)
	// }
	// // fmt.Printf("Получен пост: %+v\n", postInstance)


	// вызов по id  вызов:(http://127.0.0.1:3000/post/2)
	app.Get("/getpost/:title", func(c *fiber.Ctx) error {
		return routes.GetPostByID(c)
	})

	//удаление по id  вызов:(http://127.0.0.1:3000/delpost/1)
	app.Get("/delpost/:title", func(c *fiber.Ctx) error {
		return routes.DeletePostByTitle(c)
	})


}
