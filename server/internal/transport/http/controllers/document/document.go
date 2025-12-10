package documentController

import (
	"fmt"
	documentService "research-data-saver/internal/services/handlers/document"

	"github.com/gofiber/fiber/v2"
)

type createDocumentBody struct {
	Name           string `json:"name"`
	Annotation     string `json:"annotation"`
	Link           string `json:"link"`
	PublishingDate string `json:"publishingDate"`
	Author         string `json:"author"`
}

type DocumentController struct {
	router  *fiber.App
	service *documentService.DocumentService
}

func Init(router *fiber.App, service *documentService.DocumentService) *DocumentController {
	return &DocumentController{
		router:  router,
		service: service,
	}
}

func (controller *DocumentController) Start(route string) {
	documentRouter := controller.router.Group(route)

	documentRouter.Post("/create", controller.CreateDocument)
	documentRouter.Get("/all", controller.GetDocuments)
}

func (controller *DocumentController) CreateDocument(c *fiber.Ctx) error {
	var body createDocumentBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	createdDocument, err := controller.service.AddDocument(body.Name, body.Annotation, body.Link, body.PublishingDate, body.Author)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create document",
		})
	}

	return c.Status(fiber.StatusOK).JSON(createdDocument)
}

func (controller *DocumentController) GetDocuments(c *fiber.Ctx) error {
	var queryName = c.Query("name")
	var queryDate = c.Query("date")

	var documents, err = controller.service.GetAll(queryName, queryDate)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get documents",
		})
	}

	return c.Status(fiber.StatusOK).JSON(documents)
}
