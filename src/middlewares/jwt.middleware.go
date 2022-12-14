package middlewares

import (
	"microservice/src/configs"
	"microservice/src/messages"

	"github.com/gofiber/fiber/v2"

	jwtMiddleware "github.com/gofiber/jwt/v2"
)

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/jwt
func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningKey: []byte(configs.GetJwtSecretKet()),
		ContextKey: "jwt", // used in private routes
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Return status 401 and failed authentication error.
			if err.Error() == "Missing or malformed JWT" {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"message": messages.JwtMissingToken,
				})
			}

			// Return status 401 and failed authentication error.
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": messages.JwtUnauthorizedToken,
			})
		},
	}

	return jwtMiddleware.New(config)
}
