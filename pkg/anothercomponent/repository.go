package anothercomponent

type Repository interface {
	FindAllOfComponent() []string
}

type MemoryRepository struct{}

func (r *MemoryRepository) FindAllOfComponent() []string {
	var components []string

	components = append(components, "appended1")
	components = append(components, "componet2")

	return components
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{}
}
