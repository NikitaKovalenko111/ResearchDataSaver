package libraryArticleController

import (
	"fmt"
	libraryArticleService "research-data-saver/internal/services/handlers/libraryArticle"

	"github.com/gofiber/fiber/v2"
)

type createLibraryArticleBody struct {
	Name            string   `json:"name"`
	Annotation      string   `json:"annotation"`
	Link            string   `json:"link"`
	PublishingDate  string   `json:"publishingDate"`
	Lang            string   `json:"lang"`
	UDK             string   `json:"udk"`
	PublisherObject string   `json:"publisherObject"`
	Publisher       string   `json:"publisher"`
	Supervisor      string   `json:"supervisor"`
	Authors         []string `json:"authors"`
}

type LibraryArticleController struct {
	router  *fiber.App
	service *libraryArticleService.LibraryArticleService
}

func Init(router *fiber.App, service *libraryArticleService.LibraryArticleService) *LibraryArticleController {
	return &LibraryArticleController{
		router:  router,
		service: service,
	}
}

func (controller *LibraryArticleController) Start(route string) {
	libraryArticleRouter := controller.router.Group(route)

	libraryArticleRouter.Post("/create", controller.CreateLibraryArticle)
	libraryArticleRouter.Get("/all", controller.GetLibraryArticles)
}

func (controller *LibraryArticleController) CreateLibraryArticle(c *fiber.Ctx) error {
	var body createLibraryArticleBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	createdArticle, err := controller.service.AddArticle(body.Name, body.Annotation, body.Link, body.PublishingDate, body.Lang, body.UDK, body.PublisherObject, body.Publisher, body.Supervisor, body.Authors)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create library article",
		})
	}

	return c.Status(fiber.StatusOK).JSON(createdArticle)
}

func (controller *LibraryArticleController) GetLibraryArticles(c *fiber.Ctx) error {
	articles, err := controller.service.GetAll()

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get library articles",
		})
	}

	return c.Status(fiber.StatusOK).JSON(articles)
}
