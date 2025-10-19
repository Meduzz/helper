package service

type (
	Service interface {
		Start() error
		Stop() error
	}

	Delegate interface {
		Visit(svc Service) error
	}
)
