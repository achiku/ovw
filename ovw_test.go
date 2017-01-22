package ovw

import (
	"reflect"
	"testing"
	"time"

	"github.com/AdrianLungu/decimal"
)

type testStruct struct {
	String   string
	Bool     bool
	Integer  int
	Float    float64
	Time     time.Time
	LargeNum testSubStruct
	Decimal  decimal.Decimal
}

type testSubStruct struct {
	IntPart int
	IsZero  bool
}

func TestMergeOverwrite(t *testing.T) {
	to := testStruct{
		String:  "default string",
		Bool:    true,
		Integer: 100,
		Float:   float64(1.001112),
		Time:    time.Now(),
		LargeNum: testSubStruct{
			IntPart: 1000,
			IsZero:  false,
		},
		Decimal: decimal.NewFromFloat(1000),
	}
	from := testStruct{
		Bool:    false,
		Integer: 3000,
		Time:    time.Now().AddDate(1, 0, 0),
		Decimal: decimal.NewFromFloat(22000),
		LargeNum: testSubStruct{
			IntPart: 222000,
			IsZero:  true,
		},
	}
	var target testStruct
	if err := MergeOverwrite(to, from, &target); err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", target)

	if expected := from.Decimal; target.Decimal.Cmp(expected) != 0 {
		t.Errorf("want %s got %s", expected, target.Decimal)
	}
	if expected := from.Bool; target.Bool != expected {
		t.Errorf("want %t got %t", expected, target.Bool)
	}
	if expected := to.Float; target.Float != expected {
		t.Errorf("want %f got %f", expected, target.Float)
	}
	if expected := to.String; target.String != expected {
		t.Errorf("want %s got %s", expected, target.String)
	}
	if expected := from.LargeNum; !reflect.DeepEqual(target.LargeNum, expected) {
		t.Errorf("want %+v got %+v", expected, target.LargeNum)
	}
}
