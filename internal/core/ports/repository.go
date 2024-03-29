package ports


type BaseRepository[T any] interface {
	Create(model T) (string, error)
	GetByUser(username string) (interface{} ,error)
	InsertData(workinforamtion interface{}, collection string) (bool, error)
	GetData(username string, collection string) (interface{}, error)
}
