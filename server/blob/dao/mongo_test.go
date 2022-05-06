package dao

import (
	mgutil "happy-car/shared/mongo"
	"testing"
)

func TestObj(t *testing.T) {
	t.Logf(mgutil.NewObjID().String())
	t.Logf(mgutil.NewObjID().Hex())
}
