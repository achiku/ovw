package ovw

import (
	"testing"
	"time"

	"github.com/fatih/structs"
)

type testStruct struct {
	Bool     bool
	Integer  int
	Float    float64
	Time     time.Time
	LargeNum testSubStruct
}

type testSubStruct struct {
	IntPart int
	IsZero  bool
}

func TestOverWrite(t *testing.T) {
	dst := testStruct{
		LargeNum: testSubStruct{
			IntPart: 10,
			IsZero:  false,
		},
	}
	src := testStruct{
		Bool:    true,
		Integer: 100,
		Float:   float64(1.001112),
		Time:    time.Now(),
		LargeNum: testSubStruct{
			IntPart: 10,
			IsZero:  false,
		},
	}
	if err := OverWrite(dst, src); err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", dst)
}

func TestStruct(t *testing.T) {
	s := struct {
		Bool    bool
		Integer int
	}{
		Bool:    false,
		Integer: 1000,
	}

	st := structs.New(s)
	if err := st.Field("Bool").Set(true); err != nil {
		t.Fatal(err)
	}
	for _, f := range st.Fields() {
		t.Logf("%+v", f.Name())
		if err := f.Set(true); err != nil {
			t.Fatal(err)
		}
	}
}
