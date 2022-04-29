package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mgo "happy-car/shared/mongo"
	mongotesting "happy-car/shared/mongo/testing"
	"os"
	"testing"
)

var mongoURI string

func TestResolveAccountID(t *testing.T) {
	// start container
	c := context.Background()

	// 自己会有一个超时机制
	mc, err := mongo.Connect(c, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}
	m := NewMongo(mc.Database("happycar"))

	_, err = m.collection.InsertMany(c, []interface{}{
		bson.M{
			mgo.IDField: mustObjectID("626956e6f80926a1a36111c5"),
			openIDField: "openid_1",
		},
		bson.M{
			mgo.IDField: mustObjectID("626956e6f80926a1a36111e8"),
			openIDField: "openid_2",
		},
	})
	if err != nil {
		t.Fatalf("cannot insert initial values: %v", err)
	}

	// 固定ID
	// 子测试
	m.newObjectID = func() primitive.ObjectID {
		return mustObjectID("62680b0a53a88506efd364ae")
	}

	cases := []struct {
		name   string
		openID string
		wantID string
	}{
		{
			name:   "existing_user",
			openID: "openid_1",
			wantID: "626956e6f80926a1a36111c5",
		},
		{
			name:   "another_existing_user",
			openID: "openid_2",
			wantID: "626956e6f80926a1a36111e8",
		},
		{
			name:   "new_user",
			openID: "openid_3",
			wantID: "62680b0a53a88506efd364ae",
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			id, err := m.ResolveAccountID(context.Background(), cc.openID)
			if err != nil {
				t.Fatalf("cannot resolve account id: %v", err)
			}
			if id != cc.wantID {
				t.Fatalf("got %s, want %s", id, cc.wantID)
			}
		})
	}

	id, err := m.ResolveAccountID(c, "123")
	if err != nil {
		t.Errorf("faild resolve account id for 123: %v", err)
	} else {
		want := "62680b0a53a88506efd364ae"
		if id != want {
			t.Errorf("resolve account id for 123: want %q, got %q", want, id)
		}
	}
}

func mustObjectID(hex string) primitive.ObjectID {
	objID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		panic(err)
	}
	return objID
}

// 这个m就是上面的测试
func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m, &mongoURI))
}
