package util

import (
	"database/sql/driver"
	"encoding/json"
)

type (
	JSON        map[string]any
	Addr[T any] []T
)

func (j JSON) MarshalBinary() (data []byte, err error) { return json.Marshal(j) }
func (j *JSON) UnmarshalBinary(data []byte) error      { return json.Unmarshal(data, j) }

func (a Addr[T]) MarshalBinary() (data []byte, err error) { return json.Marshal(a) }
func (a *Addr[T]) UnmarshalBinary(data []byte) error      { return json.Unmarshal(data, a) }

func (a Addr[T]) Value() (driver.Value, error) {
	bs, err := json.Marshal(a)
	return string(bs), err
}

func (a *Addr[T]) Scan(value any) error {
	switch value := value.(type) {
	case string:
		return json.Unmarshal([]byte(value), a)
	case []byte:
		return json.Unmarshal(value, a)
	}
	return nil
}
