package routing

type service struct{}

func New() *service {
	return &service{}
}

func (s *service) ShipsRoute() []string {
	return []string{"dock #1", "dock #2"}
}
func (s *service) TrucksRoute() []string {
	return []string{"warehouse #1", "warehoue #2", "warehouse #3"}
}
