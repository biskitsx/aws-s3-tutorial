package handler

import (
	"log"

	"github.com/biskitsx/aws-s3-tutorial/service"
	"github.com/gofiber/fiber/v2"
)

type UploadController interface {
	UploadImage(c *fiber.Ctx) error
}
type uploadController struct {
	awsService service.AwsService
}

func NewUploadController() uploadController {
	return uploadController{
		awsService: service.NewAwsService(),
	}
}

func (controller *uploadController) UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("upload")
	if err != nil {
		return err
	}
	result, err := controller.awsService.Upload(file)
	if err != nil {
		log.Fatal(err)
		return c.JSON(fiber.Map{
			"msg": "error",
		})
	}
	return c.JSON(fiber.Map{
		"msg": result.Location,
	})
}
