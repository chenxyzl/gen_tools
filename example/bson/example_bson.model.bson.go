// Code generated by https://github.com/chenxyzl/gsgen; DO NOT EDIT.
// gen_tools version: 1.1.1
// generate time: 2024-06-19 15:50:22
package bson

import (
	"github.com/chenxyzl/gsgen/gsmodel"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *TestA) MarshalBSON() ([]byte, error) {
	var doc = bson.M{"cc": s.cc, "dd": s.dd}
	return bson.Marshal(doc)
}
func (s *TestA) UnmarshalBSON(data []byte) error {
	doc := struct {
		Cc *gsmodel.DList[int]           `bson:"cc"`
		Dd *gsmodel.DMap[string, *TestA] `bson:"dd"`
	}{}
	if err := bson.Unmarshal(data, &doc); err != nil {
		return err
	}
	s.setCc(doc.Cc)
	s.setDd(doc.Dd)
	return nil
}
func (s *TestA) BuildBson(m bson.M, preKey string) {
	dirty := s.GetDirty()
	if dirty == 0 {
		return
	}
	if dirty&(1<<0) != 0 {
		if s.cc == nil {
			gsmodel.AddUnsetDirtyM(m, gsmodel.MakeBsonKey("cc", preKey))
		} else {
			s.cc.BuildBson(m, gsmodel.MakeBsonKey("cc", preKey))
		}
	}
	if dirty&(1<<1) != 0 {
		if s.dd == nil {
			gsmodel.AddUnsetDirtyM(m, gsmodel.MakeBsonKey("dd", preKey))
		} else {
			s.dd.BuildBson(m, gsmodel.MakeBsonKey("dd", preKey))
		}
	}
	return
}
func (s *TestA) CleanDirty() {
	s.DirtyModel.CleanDirty()
	if s.cc != nil {
		s.cc.CleanDirty()
	}
	if s.dd != nil {
		s.dd.CleanDirty()
	}
}
