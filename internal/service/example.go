package service

type ExampleService struct{}

func NewExampleService() *ExampleService {
	return &ExampleService{}
}

func (s *ExampleService) Greet() string {
	return "Hello, Service Layer!"
}