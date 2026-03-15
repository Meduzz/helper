package block

import (
	"errors"
	"log"
	"sync"

	"github.com/Meduzz/helper/fp/slice"
)

var (
	hooks []Hook
	mu    = &sync.RWMutex{}
)

func init() {
	go func() {
		err := Block(func() error {
			mu.RLock()
			defer mu.RUnlock()

			return slice.Fold(hooks, nil, func(h Hook, agg error) error {
				err := h()

				if err != nil {
					return errors.Join(agg, err)
				}

				return agg
			})
		})

		if err != nil {
			log.Fatalf("Shutdown hooks threw error(s): %v\n", err) // TODO good enough?
		}
	}()
}

func RegisterShutdownHook(it Hook) {
	mu.Lock()
	defer mu.Unlock()

	hooks = append(hooks, it)
}
