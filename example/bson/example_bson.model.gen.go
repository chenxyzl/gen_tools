// Code generated by https://github.com/chenxyzl/gsgen; DO NOT EDIT.
// gen_tools version: 1.1.1
// generate time: 2024-06-19 15:50:22
package bson

import (
	"encoding/json"
	"fmt"
	"github.com/chenxyzl/gsgen/gsmodel"
)

func (s *TestA) GetCc() *gsmodel.DList[int] {
	return s.cc
}
func (s *TestA) setCc(v *gsmodel.DList[int]) {
	if v != nil {
		v.SetParent(0, s.UpdateDirty)
	}
	s.cc = v
	s.UpdateDirty(1 << 0)
}
func (s *TestA) GetDd() *gsmodel.DMap[string, *TestA] {
	return s.dd
}
func (s *TestA) setDd(v *gsmodel.DMap[string, *TestA]) {
	if v != nil {
		v.SetParent(1, s.UpdateDirty)
	}
	s.dd = v
	s.UpdateDirty(1 << 1)
}
func (s *TestA) String() string {
	doc := struct {
		Cc *gsmodel.DList[int]           `bson:"cc"`
		Dd *gsmodel.DMap[string, *TestA] `bson:"dd"`
	}{s.cc, s.dd}
	return fmt.Sprintf("%v", &doc)
}
func (s *TestA) MarshalJSON() ([]byte, error) {
	doc := struct {
		Cc *gsmodel.DList[int]           `bson:"cc"`
		Dd *gsmodel.DMap[string, *TestA] `bson:"dd"`
	}{s.cc, s.dd}
	return json.Marshal(doc)
}
func (s *TestA) UnmarshalJSON(data []byte) error {
	doc := struct {
		Cc *gsmodel.DList[int]           `bson:"cc"`
		Dd *gsmodel.DMap[string, *TestA] `bson:"dd"`
	}{}
	if err := json.Unmarshal(data, &doc); err != nil {
		return err
	}
	s.setCc(doc.Cc)
	s.setDd(doc.Dd)
	return nil
}
func (s *TestA) Clone() (*TestA, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	ret := TestA{}
	err = json.Unmarshal(data, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
