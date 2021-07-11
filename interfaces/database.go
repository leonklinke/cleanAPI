package interfaces

type DatabaseHandler interface {
	Begin() (Tx, error)
	Query(query string, args ...interface{}) (Row, error)
	Exec(query string, args ...interface{}) (Result, error)
}

type Tx interface {
} //(*sql.Tx, error)
type Row interface {
} //(*sql.Row, error)
type Result interface {
} //(sql.Result, error)

func NewDatabaseHandler() *DatabaseHandler {
	return new(DatabaseHandler)
}
