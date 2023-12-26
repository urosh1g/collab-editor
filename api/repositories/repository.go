package repositories

type Repository interface {
	GetAll() ([]interface{}, error)
	GetOne(id int64) (interface{}, error)
}
