package v1

import (
	"fmt"
	"image"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/venqoi/racc-api/utils"
)

func GetRaccoonOfTheDay(c *fiber.Ctx) error {
	var wantsJSON = utils.WantsJSON(c)

	// sets seed for this day
	utils.SetSeed("daily")
	// set index
	var index = utils.GetIndex()

	bytes, err := os.ReadFile("raccs/racc" + fmt.Sprint(index) + ".jpg")

	c.Set("X-Raccoon-Index", fmt.Sprint(index))

	if err != nil {
		println("error while reading racc photo", err.Error())
		if wantsJSON {
			return c.Status(500).JSON(utils.Response{
				Success: false,
				Message: "An error occurred whilst fetching file",
			})
		}

		return c.SendStatus(500)
	}

	if wantsJSON {
		file, err := os.Open("./raccs/racc" + fmt.Sprint(index) + ".jpg")

		if err != nil {
			println(err.Error())
		}

		defer file.Close()

		image, _, err := image.DecodeConfig(file)

		if err != nil {
			println(err.Error())
		}

		return c.JSON(utils.Response{
			Success: true,
			Data: utils.ImageStruct{
				URL:    utils.BaseURL(c) + "/v1/raccoon/" + fmt.Sprint(index),
				Index:  index,
				Width:  image.Width,
				Height: image.Height,
				Alt:    utils.GetAlti(index),
			},
		})
	}

	c.Set("Content-Type", "image/jpeg")
	return c.Send(bytes)
}
