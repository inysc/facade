package util

import "encoding/json"

type (
	JSON        map[string]any
	Addr[T any] []T
)

func (j JSON) MarshalBinary() (data []byte, err error) { return json.Marshal(j) }
func (j *JSON) UnmarshalBinary(data []byte) error      { return json.Unmarshal(data, j) }

func (a Addr[T]) MarshalBinary() (data []byte, err error) { return json.Marshal(a) }
func (a *Addr[T]) UnmarshalBinary(data []byte) error      { return json.Unmarshal(data, a) }
func (a *Addr[T]) String() string {
	bs, _ := json.Marshal(a)
	return string(bs)
}
