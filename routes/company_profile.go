package routes

import (
	"github.com/gofiber/fiber/v2"

	companyProfileSvc "github.com/svede-frontend/internal/handler/company_profile"
)

func BuildCompanyProfileRoutes(service fiber.Router) {
	service.Get("/", companyProfileSvc.CompanyProfileHandler)
}
