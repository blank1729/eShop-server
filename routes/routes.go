package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	SetUserRoutes(r)

	SetStoreRoutes(r)

	SetCategoryRoutes(r)
	SetProductRoutes(r)
	SetCustomerRoutes(r)

	SetOptionRoutes(r)
	// option variants

	// product variants
	// images routes

	// orders
	SetOrderRoutes(r)
	// carts
	SetCartRoutes(r)
	SetCartItemRoutes(r)
	// address
	SetAddressRoutes(r)
}
