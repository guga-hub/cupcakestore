package routers

import (
	"github.com/guga_hub/cupcakestore/controllers"
	"github.com/guga_hub/cupcakestore/database"
	"github.com/guga_hub/cupcakestore/middlewares"
	"github.com/guga_hub/cupcakestore/repositories"
	"github.com/guga_hub/cupcakestore/services"
	"github.com/gofiber/fiber/v2"
)

type ProductRouter struct {
	productController controllers.ProductController
}

func NewProductRouter() *ProductRouter {
	// Initialize repositories
	productRepository := repositories.NewProductRepository(database.DB)

	// Initialize services with repositories
	productService := services.NewProductService(productRepository)

	// Initialize controllers with services
	productController := controllers.NewProductController(productService)

	return &ProductRouter{
		productController: productController,
	}
}

func (r *ProductRouter) InstallRouters(app *fiber.App) {
	product := app.Group("/products")
	product.Get("/details/:id", r.productController.RenderDetails)

	productAdmin := app.Group("/products").Use(middlewares.LoginAndStaffRequired())
	productAdmin.Get("/create", r.productController.RenderCreate)
	productAdmin.Post("/create", r.productController.Create)
	productAdmin.Get("/json", r.productController.JSONProducts)
	productAdmin.Post("/update/:id", r.productController.Update)
	productAdmin.Get("/delete/:id", r.productController.RenderDelete)
	productAdmin.Post("/delete/:id", r.productController.Delete)
	productAdmin.Get("/", r.productController.RenderProducts)
	productAdmin.Get("/:id", r.productController.RenderProduct)
}
