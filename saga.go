package compensators

type Saga struct {
	compensations []func()
}

func NewSaga() Saga {
	return Saga{}
}

func (s *Saga) AddCompensation(callable func()) {
	s.compensations = append(s.compensations, callable)
}

func (s *Saga) Compensate() {
	for _, callable := range s.compensations {
		callable()
	}
}

func (s *Saga) Exec(callable func()) {
	defer func() {
		if err := recover(); err != nil {
			s.Compensate()

			panic("Saga compensated an error")
		}
	}()

	callable()
}
