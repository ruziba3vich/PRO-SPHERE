package repo

type NewsRepository interface {
	CreateCategory()
	UpdateCategory()
	GetCategoryById()
	GetCategoryByName()
	GetCategoryByOrderNumber()
}
