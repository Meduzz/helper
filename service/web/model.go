package web

import (
	"fmt"

	"github.com/Meduzz/helper/service"
	"github.com/gin-gonic/gin"
)

type (
	WebApi interface {
		Setup(router *gin.Engine)
	}

	webDelegate struct{}
)

var _ service.Delegate = webDelegate{}

func (webDelegate) Visit(svc service.Service) error {
	api, ok := svc.(WebApi)

	if ok {
		if server != nil {
			api.Setup(server)
		} else {
			return fmt.Errorf("server not setup")
		}
	}

	return nil
}
