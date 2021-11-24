package Interface

import (
	"github.com/sonyarouje/simdb"
)

type ConnectionInterface interface {
	Open(entity simdb.Entity) *simdb.Driver
	Errors() []error
	Insert(entity simdb.Entity) (err error)
	Where(key string, cond string, val interface{}) *simdb.Driver
	Get() *simdb.Driver
	First() *simdb.Driver
	Raw() interface{}
	RawArray() []interface{}
	AsEntity(output interface{}) (err error)
	Update(entity simdb.Entity) (err error)
	Upsert(entity simdb.Entity) (err error)
	Delete(entity simdb.Entity) (err error)
}
