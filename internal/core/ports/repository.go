package ports


type BaseRepository[T any] interface {
	Create(model T) (string, error)
	GetByEmail(username string) (interface{} ,error)
	InsertData(username string, workinforamtion interface{}, collection string) (string, error)
	GetData(username string, collection string) (interface{}, error)
}
