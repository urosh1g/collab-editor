package repositories

type Repository[T any] interface {
	GetAll() ([]T, error)
	GetOne(id int64) (T, error)
	Create(entity any) (T, error)
	Update(id int64, entity any) (T, error)
	Delete(id int64) (T, error)
}
