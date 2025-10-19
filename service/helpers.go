package service

import (
	"errors"

	"github.com/Meduzz/helper/fp/slice"
)

var (
	services  = make([]Service, 0)
	delegates = make([]Delegate, 0)
)

func AddService(svc Service) {
	services = append(services, svc)
}

func AddDelegate(del Delegate) {
	delegates = append(delegates, del)
}

func Start() error {
	return slice.Fold(services, nil, func(svc Service, agg error) error {
		if agg != nil {
			return agg
		}

		err := svc.Start()

		if err != nil {
			return err
		}

		return slice.Fold(delegates, nil, func(d Delegate, agg2 error) error {
			if agg2 != nil {
				return agg2
			}

			return d.Visit(svc)
		})
	})
}

func Stop() error {
	var sum error = nil

	slice.ForEach(services, func(svc Service) {
		err := svc.Stop()

		if err != nil {
			sum = errors.Join(sum, err)
		}
	})

	return sum
}
