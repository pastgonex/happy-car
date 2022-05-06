package dao

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/testing/protocmp"
	rentalpb "happy-car/rental/api/gen/v1"
	"happy-car/shared/id"
	mgutil "happy-car/shared/mongo"
	"happy-car/shared/mongo/objid"
	mongotesting "happy-car/shared/mongo/testing"
	"os"
	"testing"
)

/*
测试， 一次劳动，然后整个项目的生命周期， 我们就都可以使用他们
*/

func TestCreateTrip(t *testing.T) {
	// start container
	c := context.Background()

	// 自己会有一个超时机制
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}
	db := mc.Database("happycar")
	err = mongotesting.SetupIndexes(c, db)
	if err != nil {
		t.Fatalf("cannot setup indexes: %v", err)
	}
	m := NewMongo(db)

	// 想测试 tripId这个trip能不能带着 accountId 和 tripStatus成功，如果成功：返回trip
	cases := []struct {
		name       string
		tripId     string
		accountId  string
		tripStatus rentalpb.TripStatus
		wantErr    bool
	}{
		{
			name:       "finished",
			tripId:     "626eb0fd458fedbb7d1c4d5e",
			accountId:  "account1",
			tripStatus: rentalpb.TripStatus_FINISHED,
			wantErr:    false,
		},
		{
			name:       "another_finished",
			tripId:     "626eb0fd458fedbb7d1c4d5f",
			accountId:  "account1",
			tripStatus: rentalpb.TripStatus_FINISHED,
			wantErr:    false,
		},
		{
			name:       "in_progress",
			tripId:     "626eb0fd458fedbb7d1c4d6e",
			accountId:  "account1",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
			wantErr:    false,
		},
		{
			// 第二个 in_progress 应该出错，只能有一个正在进行的行程
			name:       "another_in_progress",
			tripId:     "626eb0fd458fedbb7d1c4d6f",
			accountId:  "account1",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
			wantErr:    true,
		},
		{
			// 另一个用户创建了一个正在进行的行程，是可以的
			name:       "in_progress_by_another_account",
			tripId:     "626eb0fd458fedbb7d1c4d50",
			accountId:  "account2",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
			wantErr:    false,
		},
	}

	for _, cc := range cases {
		//mgutil.NewObjID = func() primitive.ObjectID {
		//	return objid.MustFromID(id.TripID(cc.tripId))
		//}
		mgutil.NewObjIDWithValue(id.TripID(cc.tripId))
		tr, err := m.CreateTrip(c, &rentalpb.Trip{
			AccountId: cc.accountId,
			Status:    cc.tripStatus,
		})
		if cc.wantErr { // want error, but no error.
			if err == nil {
				t.Errorf("%s: error expected; got none", cc.name)
			}
			continue
		}
		if err != nil {
			t.Errorf("%s: error creating trip: %v", cc.name, err)
			continue
		}
		if tr.ID.Hex() != cc.tripId {
			t.Errorf("%s: incorrect trip id; want trip id %q; got %q", cc.name, cc.tripId, tr.ID.Hex())
		}
	}
}

func TestGetTrip(t *testing.T) {
	// start container
	c := context.Background()

	// 自己会有一个超时机制
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}
	m := NewMongo(mc.Database("happycar"))

	acct := id.AccountID("account1")
	mgutil.NewObjID = primitive.NewObjectID
	tripRecord, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: acct.String(),
		CarId:     "car1",
		Start: &rentalpb.LocationStatus{
			Location: &rentalpb.Location{
				Latitude:  30,
				Longitude: 120,
			},
			PoiName: "start point",
		},
		End: &rentalpb.LocationStatus{
			Location: &rentalpb.Location{
				Latitude:  35,
				Longitude: 115,
			},
			FeeCent:  10000,
			KmDriven: 35,
			PoiName:  "end point",
		},
		Status: rentalpb.TripStatus_FINISHED,
	})
	if err != nil {
		t.Fatalf("cannot create trip: %v", err)
	}
	t.Logf("inserted row %s with updatedat %v", tripRecord.ID, tripRecord.UpdatedAt)

	got, err := m.GetTrip(c, objid.ToTripID(tripRecord.ID), acct)
	if err != nil {
		t.Errorf("cannot get trip: %v", err)
	}

	// 故意出错
	//got.Trip.Start.PoiName = "bad start"

	// 第三个参数不加会出错， 记住就好
	if diff := cmp.Diff(tripRecord, got, protocmp.Transform()); diff != "" {
		t.Errorf("result differs; -want +got: %s", diff)
	}
}

func TestGetTrips(t *testing.T) {
	// start container
	c := context.Background()

	// 自己会有一个超时机制
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}
	m := NewMongo(mc.Database("happycar"))

	rows := []struct {
		id        string
		accountID id.AccountID
		status    rentalpb.TripStatus
	}{
		{
			id:        "626eb0fd428fedbb7d1c4d5e",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FINISHED,
		},
		{
			id:        "626eb2b42851e667d3f454ab",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FINISHED,
		},
		{
			id:        "626eb2b42851e667d3f454af",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FINISHED,
		},
		{
			id:        "626eb6ec42e4717e4b1c730e",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_IN_PROGRESS,
		},
		{
			id:        "626eee424ac340b4ff867a12",
			accountID: "account_id_for_get_trips_1",
			status:    rentalpb.TripStatus_IN_PROGRESS,
		},
	}
	for _, r := range rows {
		mgutil.NewObjIDWithValue(id.TripID(r.id))
		_, err := m.CreateTrip(c, &rentalpb.Trip{
			AccountId: r.accountID.String(),
			Status:    r.status,
		})
		if err != nil {
			t.Fatalf("cannot create rows: %v", err)
		}
	}

	cases := []struct {
		name       string
		accountID  string
		status     rentalpb.TripStatus
		wantCount  int    // 期待获取几个结果
		wantOnlyID string // 只有在wantCount是1的时候
	}{
		{
			name:      "get_all",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_TS_NOT_SPECIFIED,
			wantCount: 4,
		},
		{
			name:       "get_int_progress",
			accountID:  "account_id_for_get_trips",
			status:     rentalpb.TripStatus_IN_PROGRESS,
			wantCount:  1,
			wantOnlyID: "626eb6ec42e4717e4b1c730e",
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			res, err := m.GetTrips(context.Background(), id.AccountID(cc.accountID), cc.status)
			for _, v := range res {
				t.Logf("got trip %+v\n", v)
			}
			if err != nil {
				t.Errorf("cannot get trips: %v", err)
			}
			if cc.wantCount != len(res) {
				t.Errorf("incorrect result count; want %d, got %d", cc.wantCount, len(res))
			}
			if cc.wantOnlyID != "" && len(res) > 0 {
				if cc.wantOnlyID != res[0].ID.Hex() {
					t.Errorf("only_id incorect; want %q; get %q", cc.wantOnlyID, res[0].ID.Hex())
				}
			}
		})
	}
}

// 控制时间戳进行测试
func TestUpdateTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	m := NewMongo(mc.Database("happycar"))
	tid := id.TripID("626eb6ec12e4717e4b1c730e")
	aid := id.AccountID("account_for_update")
	var now int64 = 10000
	mgutil.NewObjIDWithValue(tid)
	mgutil.UpdatedAt = func() int64 {
		return now
	}

	tripRecord, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: aid.String(),
		Status:    rentalpb.TripStatus_IN_PROGRESS,
		Start: &rentalpb.LocationStatus{
			PoiName: "start_poi",
		},
	})
	if err != nil {
		t.Fatalf("cannot create trip: %v", err)
	}

	if tripRecord.UpdatedAt != now {
		t.Fatalf("wrong updatedat; want: %d, got: %d", now, tripRecord.UpdatedAt)
	}

	update := &rentalpb.Trip{
		AccountId: aid.String(),
		Status:    rentalpb.TripStatus_IN_PROGRESS,
		Start: &rentalpb.LocationStatus{
			PoiName: "start_poi_updated",
		},
	}

	cases := []struct {
		name          string
		now           int64
		withUpdatedAt int64
		wantErr       bool
	}{
		{
			name:          "normal_update",
			now:           20000,
			withUpdatedAt: 10000,
			wantErr:       false,
		},
		{
			name:          "update_with_stale_timestamp", // 老的
			now:           30000,
			withUpdatedAt: 10000,
			wantErr:       true,
		},
		{
			name:          "update_with_refetch",
			now:           40000,
			withUpdatedAt: 20000,
			wantErr:       false,
		},
	}

	for _, cc := range cases {
		now = cc.now
		err := m.UpdateTrip(c, tid, aid, cc.withUpdatedAt, update)
		if cc.wantErr {
			if err == nil {
				t.Errorf("want error but got nil")
			} else {
				continue
			}
		} else {
			if err != nil {
				t.Errorf("cannot update trip: %v", err)
			}
		}
		updatedTrip, err := m.GetTrip(c, tid, aid)
		if err != nil {
			t.Errorf("%s: cannot get trip after update: %v", cc.name, err)
		}
		if now != updatedTrip.UpdatedAt {
			t.Errorf("%s: incorrect updatedat. want %d; got %d", cc.name, cc.now, updatedTrip.UpdatedAt)
		}
	}

}

// 这个m就是上面的测试
func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
