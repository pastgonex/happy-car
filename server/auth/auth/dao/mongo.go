package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mgo "happy-car/shared/mongo"
)

const openIDField = "open_id"

// Mongo defines a mongo dao(data access object).
type Mongo struct {
	collection  *mongo.Collection
	newObjectID func() primitive.ObjectID
}

// NewMongo 责任划分 creates a new mongo dao.
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		collection:  db.Collection("account"),
		newObjectID: primitive.NewObjectID,
	}
}

//function resolveOpenID(open_id) {
//	return db.account.findAndModify({
//		query: {open_id: open_id},
//		update: {$set: {open_id: open_id}},
//		upsert: true,
//		returnNewDocument: true, // 暂时替代new的方案， 等待下一版本的 driver修复
//		// new: true
//	});
//}
//resolveOpenID('111');

func (m *Mongo) ResolveAccountID(c context.Context, openID string) (string, error) {
	//m.collection.InsertOne(c, bson.M{
	//	mgo.IDField: m.newObjectID(),
	//	openIDField: openID,
	//})
	insertedID := m.newObjectID()
	res := m.collection.FindOneAndUpdate(
		c,
		bson.M{
			openIDField: openID,
		},
		mgo.SetOnInsert(bson.M{
			mgo.IDField: insertedID,
			openIDField: openID,
		}),
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After),
	)
	if err := res.Err(); err != nil {
		return "", fmt.Errorf("cannot findOneAndUpdate: %v", err)
	}
	//var row struct {
	//	ID primitive.ObjectID `bson:"_id"`
	//}
	row := mgo.ObjID{}
	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("cannot decode %v", err)
	}
	return row.ID.Hex(), nil
}
