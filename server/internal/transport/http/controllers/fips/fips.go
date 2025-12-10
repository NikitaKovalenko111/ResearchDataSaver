package fipsController

import (
	fipsService "research-data-saver/internal/services/handlers/fips"

	"github.com/gofiber/fiber/v2"
)

type createFipsBody struct {
	Name           string   `json:"name"`
	Link           string   `json:"link"`
	Type           string   `json:"type"`
	Annotation     string   `json:"annotation"`
	Registration   string   `json:"registration"`
	PublishingDate string   `json:"publishingDate"`
	Applicant      string   `json:"applicant"`
	Address        string   `json:"address"`
	Authors        []string `json:"authors"`
}

type FipsController struct {
	router  *fiber.App
	service *fipsService.FipsService
}

func Init(router *fiber.App, service *fipsService.FipsService) *FipsController {
	return &FipsController{
		router:  router,
		service: service,
	}
}

func (controller *FipsController) Start(route string) {
	fipsRouter := controller.router.Group(route)

	fipsRouter.Post("/create", controller.CreateFips)
	fipsRouter.Get("/all", controller.GetFips)
}

func (controller *FipsController) CreateFips(c *fiber.Ctx) error {
	var body createFipsBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	content, err := controller.service.Create(body.Name, body.Link, body.Type, body.Annotation, body.Registration, body.PublishingDate, body.Applicant, body.Address, body.Authors)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create FIPS",
		})
	}

	return c.Status(fiber.StatusOK).JSON(content)
}

func (controller *FipsController) GetFips(c *fiber.Ctx) error {
	var queryName = c.Query("name")
	var queryFipsType = c.Query("fipsType")
	var queryReg = c.Query("reg")
	var queryDate = c.Query("date")

	fips, err := controller.service.GetAll(queryName, queryFipsType, queryReg, queryDate)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create FIPS",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fips)
}
