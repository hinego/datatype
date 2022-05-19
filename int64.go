package datatype

import (
	"database/sql/driver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"reflect"
	"sync/atomic"
)

// Int64 is an atomic wrapper around an int64.
type Int64 int64

func NewInt64(i int64) *Int64 {
	data := Int64(i)
	return &data
}
func (Int64) GormDataType() string {
	return "bigint"
}

func (Int64) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "BIGINT"
	case "mysql":
		return "BIGINT"
	case "postgres":
		return "BIGINT"
	}
	return ""
}

// Load atomically loads the wrapped value.
func (r *Int64) Load() int64 {
	v := int64(*r)
	return atomic.LoadInt64(&v)
}

// Add atomically adds to the wrapped int64 and returns the new value.
func (r *Int64) Add(n int64) int64 {
	v := int64(*r)
	return atomic.AddInt64(&v, n)
}

// Sub atomically subtracts from the wrapped int64 and returns the new value.
func (r *Int64) Sub(n int64) int64 {
	v := int64(*r)
	return atomic.AddInt64(&v, -n)
}

// Inc atomically increments the wrapped int64 and returns the new value.
func (r *Int64) Inc() int64 {
	return r.Add(1)
}

// Dec atomically decrements the wrapped int64 and returns the new value.
func (r *Int64) Dec() int64 {
	return r.Sub(1)
}

// CAS is an atomic compare-and-swap.
func (r *Int64) CAS(old, new int64) bool {
	v := int64(*r)
	return atomic.CompareAndSwapInt64(&v, old, new)
}

// Store atomically stores the passed value.
func (r *Int64) Store(n int64) {
	v := int64(*r)
	atomic.StoreInt64(&v, n)
}

// Swap atomically swaps the wrapped int64 and returns the old value.
func (r *Int64) Swap(n int64) int64 {
	v := int64(*r)
	return atomic.SwapInt64(&v, n)
}

func (r *Int64) Value() (driver.Value, error) {
	log.Println("Value", r.Load())
	return r.Load(), nil
}
func (r *Int64) Scan(val interface{}) error {
	log.Println("Scan", val)
	log.Println("Scan", reflect.TypeOf(val))
	return nil
}
