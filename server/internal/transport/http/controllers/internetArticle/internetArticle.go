package internetArticleController

import (
	internetArticleService "research-data-saver/internal/services/handlers/internetArticle"

	"github.com/gofiber/fiber/v2"
)

type createInternetArticleBody struct {
	Name             string `json:"name"`
	Annotation       string `json:"annotation"`
	Link             string `json:"link"`
	PublishingDate   string `json:"publishingDate"`
	Author           string `json:"author"`
	SearchingMachine string `json:"searchingMachine"`
}

type InternetArticleController struct {
	router  *fiber.App
	service *internetArticleService.InternetArticleService
}

func Init(router *fiber.App, service *internetArticleService.InternetArticleService) *InternetArticleController {
	return &InternetArticleController{
		router:  router,
		service: service,
	}
}

func (controller *InternetArticleController) Start(route string) {
	internetArticleRouter := controller.router.Group(route)

	internetArticleRouter.Post("/create", controller.CreateInternetArticle)
	internetArticleRouter.Get("/all", controller.GetInternetArticles)
}

func (controller *InternetArticleController) CreateInternetArticle(c *fiber.Ctx) error {
	var body createInternetArticleBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	createdArticle, err := controller.service.AddArticle(body.Name, body.Annotation, body.Link, body.PublishingDate, body.Author, body.SearchingMachine)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create internet article",
		})
	}

	return c.Status(fiber.StatusOK).JSON(createdArticle)
}

func (controller *InternetArticleController) GetInternetArticles(c *fiber.Ctx) error {
	var queryName = c.Query("name")
	var querySm = c.Query("sm")
	var queryDate = c.Query("date")

	var articles, err = controller.service.GetAll(queryName, querySm, queryDate)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get internet articles",
		})
	}

	return c.Status(fiber.StatusOK).JSON(articles)
}
