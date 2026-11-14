package delegates

import (
	"github.com/Meduzz/helper/service"
	"github.com/gin-gonic/gin"
)

type (
	WithGin interface {
		Setup(*gin.Engine)
	}

	webApiDelegate struct {
		srv *gin.Engine
	}
)

var _ service.Delegate = (*webApiDelegate)(nil)

func NewWebApiDelegate(srv *gin.Engine) service.Delegate {
	return &webApiDelegate{
		srv: srv,
	}
}

func (d *webApiDelegate) Visit(svc service.Service) error {
	if api, ok := svc.(WithGin); ok {
		api.Setup(d.srv)
	}
	return nil
}
