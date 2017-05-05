//go:generate goagen bootstrap -d github.com/tikasan/goa-simple-sample/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tikasan/goa-simple-sample/app"
	"github.com/tikasan/goa-simple-sample/controller"
)

func main() {
	// Create service
	service := goa.New("goa simple sample")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "actions" controller
	c := controller.NewActionsController(service)
	app.MountActionsController(service, c)
	// Mount "array" controller
	c2 := controller.NewArrayController(service)
	app.MountArrayController(service, c2)
	// Mount "method" controller
	c3 := controller.NewMethodController(service)
	app.MountMethodController(service, c3)
	// Mount "security" controller
	c4 := controller.NewSecurityController(service)
	app.MountSecurityController(service, c4)
	// Mount "swagger" controller
	c5 := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, c5)
	// Mount "validation" controller
	c6 := controller.NewValidationController(service)
	app.MountValidationController(service, c6)
	// Mount "view" controller
	c7 := controller.NewViewController(service)
	app.MountViewController(service, c7)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
