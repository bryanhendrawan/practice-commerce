package middleware

import (
	"os"
	"practice-commerce/entity"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const (
	HeaderToken = "token"
)

func CreateToken(id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["merchant_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func Authorize() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(entity.CommonResponse{
				Message: "Not authorize",
			})
		},
		SigningKey: jwtware.SigningKey{
			Key: []byte(os.Getenv("JWT_SECRET")),
		},
	})
}

// still not in used
// this is to force request / param merchant_id
func GetMerchantIDFromToken(c *fiber.Ctx) int {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return int(claims["merchant_id"].(float64))
}
