package web

import "github.com/Meduzz/helper/service"

func init() {
	service.AddDelegate(webDelegate{})
}
