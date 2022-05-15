package datatypes

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/netip"
)

func test() {
	var _ = IPPrefix{}
}

type IPPrefix []netip.Prefix

func (IPPrefix) GormDataType() string {
	return "json"
}

func (IPPrefix) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}
func (m IPPrefix) Value() (driver.Value, error) {
	ba, err := m.MarshalJSON()
	return string(ba), err
}
func (m IPPrefix) MarshalJSON() ([]byte, error) {
	t := ([]netip.Prefix)(m)
	return json.Marshal(t)
}
func (m *IPPrefix) UnmarshalJSON(b []byte) error {
	var t = make([]netip.Prefix, 0)
	err := json.Unmarshal(b, &t)
	*m = t
	return err
}
func (m *IPPrefix) Scan(val interface{}) error {
	if val == nil {
		*m = make([]netip.Prefix, 0)
		return nil
	}
	var ba []byte
	switch v := val.(type) {
	case []byte:
		ba = v
	case string:
		ba = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", val))
	}
	t := make([]netip.Prefix, 0)
	err := json.Unmarshal(ba, &t)
	*m = t
	return err
}
