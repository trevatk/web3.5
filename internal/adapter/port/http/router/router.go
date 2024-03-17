package router

import (
	"github.com/labstack/echo/v4"

	"go.uber.org/zap"

	tcontroller "github.com/trevatk/go-pkg/http/controller"
	"github.com/trevatk/web3.5/internal/adapter/port/http/controller"
	"github.com/trevatk/web3.5/internal/core/domain"
)

// New echo router
func New(logger *zap.Logger, assessments domain.Assessments) *echo.Echo {

	e := echo.New()

	controllers := []interface{}{
		controller.NewAssessments(logger, assessments),
	}

	api := e.Group("/api")
	v1 := api.Group("/v1")

	for _, c := range controllers {

		if v, ok := c.(tcontroller.RootController); ok {
			v.RegisterRoutesV0(e)
		}

		if v, ok := c.(tcontroller.Controller); ok {
			v.RegisterRoutesV1(v1)
		}
	}

	return e
}
