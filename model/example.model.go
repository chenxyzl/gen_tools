package model

import "gen_tools/model/mdata"

type TestA struct {
	id               uint64 `bson:"_id"`
	a                int64  `bson:"a"`
	b                int32  `bson:"b"`
	mdata.DirtyModel `bson:"-"`
}
type TestB struct {
	id               uint64                      `bson:"_id"`
	m                string                      `bson:"m"`
	n                *TestA                      `bson:"n"` // 内嵌结构体类型 A
	c                *mdata.MList[*TestA]        `bson:"c"`
	d                *mdata.MMap[string, *TestA] `bson:"d"`
	mdata.DirtyModel `bson:"-"`
}
type TestC struct {
	id               uint64 `bson:"_id"`
	x                string `bson:"x"`
	y                *TestB `bson:"y"` // 内嵌结构体类型 B
	mdata.DirtyModel `bson:"-"`
}
