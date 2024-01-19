package routes

import (
	"github.com/go-errors/errors"
	"gorm.io/gorm"
	database "tmp/database/controller"
	models "tmp/database/models"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("епта работает")
}

// добавление поста
func AddPost(c *fiber.Ctx) error {
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.DB.Create(&post)
	return c.Status(200).JSON(post)
}

// GetPostByID получение поста по id
func GetPostByID(c *fiber.Ctx) error {
	title := c.Params("title")
	if title == "" {
		return c.Status(400).JSON(fiber.Map{"ошибка":"нужно название поста"})
	}
	var post models.Post 
	result := database.DB.DB.First(&post, "title = ?", title)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"ошибка": "сообщение не найдено("})
		}
		return c.Status(500).JSON(fiber.Map{"ошибка": "внутренняя ошибка сервера("})
	}
	return c.Status(200).JSON(post)
}
//удаление поста по id
// ...
func DeletePostByTitle(c *fiber.Ctx) error {
	title := c.Params("title")
	if title == "" {
		return c.Status(400).JSON(fiber.Map{"ошибка":"нужно название поста"})
	}
	var post models.Post
	result := database.DB.DB.Where("title = ?", title).Delete(&post)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"ошибка": "внутренняя ошибка сервера"})
	}
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"ошибка": "пост не найдено"})
	}
	return c.Status(200).JSON(fiber.Map{"сообщение": "плост успешно удалено"})
}
//...
