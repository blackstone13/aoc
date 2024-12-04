package main

type ProblemKey struct {
	Day, Part int
}

type ProblemFunc func()

type ProblemRegistry struct {
	registry map[ProblemKey]ProblemFunc
}

func (r ProblemRegistry) RegisterProblem(day int, part int, fn ProblemFunc) {
	key := ProblemKey{Day: day, Part: part}
	r.registry[key] = fn
}

func (r ProblemRegistry) GetProblem(day int, part int) (ProblemFunc, bool) {
	key := ProblemKey{Day: day, Part: part}
	fn, ok := r.registry[key]
	return fn, ok
}

func NewProblemRegistry() *ProblemRegistry {
	registry := &ProblemRegistry{
		registry: make(map[ProblemKey]ProblemFunc),
	}

	registry.RegisterAll()

	return registry
}
