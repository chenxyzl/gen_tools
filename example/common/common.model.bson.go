// Code generated by https://github.com/chenxyzl/gsgen; DO NOT EDIT.
// gen_tools version: 1.1.1
// generate time: 2024-06-19 15:44:41
package common

import "go.mongodb.org/mongo-driver/bson"

func (s *Common) MarshalBSON() ([]byte, error) {
	var doc = bson.M{}
	return bson.Marshal(doc)
}
func (s *Common) UnmarshalBSON(data []byte) error {
	doc := struct {
	}{}
	if err := bson.Unmarshal(data, &doc); err != nil {
		return err
	}
	return nil
}
func (s *Common) BuildBson(m bson.M, preKey string) {
	dirty := s.GetDirty()
	if dirty == 0 {
		return
	}
	return
}
func (s *Common) CleanDirty() {
	s.DirtyModel.CleanDirty()
}
