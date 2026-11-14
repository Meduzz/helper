package delegates

import (
	"github.com/Meduzz/helper/service"
	"github.com/gin-gonic/gin"
)

type (
	WithGin interface {
		Setup(*gin.Engine)
	}

	ginDelegate struct {
		srv *gin.Engine
	}
)

var _ service.Delegate = (*ginDelegate)(nil)

func NewGinDelegate(srv *gin.Engine) service.Delegate {
	return &ginDelegate{
		srv: srv,
	}
}

func (d *ginDelegate) Visit(svc service.Service) error {
	if api, ok := svc.(WithGin); ok {
		api.Setup(d.srv)
	}
	return nil
}
