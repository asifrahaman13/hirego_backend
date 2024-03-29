package ports

type BaseRepository[T any] interface {
	Create(model T, collection string) (string, error)
	GetAll(collection string) ([]map[string]interface{}, error)
	GetByField(field string, field_value string, collection string) (interface{}, error)
	InsertData(workinforamtion interface{}, collection string) (bool, error)
	GetData(username string, collection string) (interface{}, error)
}
