package controllerfactory

import "github.com/BrunoPolaski/login-service/internal/controller"

func GetAuthController() controller.AuthController {
	return controller.NewAuthController()
}
