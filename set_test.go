package set

import "testing"
import "reflect"

func InitTestEnv() *Set {
	s := NewSet("string set")
	s.AddString("a")
	s.AddString("b")
	s.AddString("c")
	s.AddInt(1)
	s.AddInt(2)
	s.AddInt(3)
	s.AddFloat(1.0)
	s.AddFloat(2.0)
	s.AddFloat(3.0)
	return s
}

func TestAddString(t *testing.T) {
	s := InitTestEnv()

	if len(s.StringMap) != 3 {
		t.Errorf("item count does not match in string set")
	}

	testMap := map[string]float64{"a": 1.0, "b": 1.0, "c": 1.0}
	if reflect.DeepEqual(s.StringMap, testMap) == false {
		t.Errorf("string set contents incorrect")
	}
}

func TestAddInt(t *testing.T) {
	s := InitTestEnv()
	if len(s.IntMap) != 3 {
		t.Errorf("item count does not match in int set")
	}
	testMap := map[int]float64{1: 1.0, 2: 1.0, 3: 1.0}
	if reflect.DeepEqual(s.IntMap, testMap) == false {
		t.Errorf("int set contents incorrect")
	}
}

func TestAddFloat(t *testing.T) {
	s := InitTestEnv()

	if len(s.FloatMap) != 3 {
		t.Errorf("item count does not match in float set")
	}

	testMap := map[float64]float64{1.0: 1.0, 2.0: 1.0, 3.0: 1.0}
	if reflect.DeepEqual(s.FloatMap, testMap) == false {
		t.Errorf("AddFloat items are incorrect")
	}
}

func TestRemoveString(t *testing.T) {
	s := InitTestEnv()

	testMap := map[string]float64{"a": 1.0, "c": 1.0}
	s.RemoveString("b")

	if reflect.DeepEqual(s.StringMap, testMap) == false {
		t.Errorf("RemoveString incorrect")
	}
}

func TestRemoveInt(t *testing.T) {
	s := InitTestEnv()

	testMap := map[int]float64{1: 1.0, 3: 1.0}
	s.RemoveInt(2)

	if reflect.DeepEqual(s.IntMap, testMap) == false {
		t.Errorf("RemoveInt incorrect")
	}
}

func TestRemoveFloat(t *testing.T) {
	s := InitTestEnv()

	testMap := map[float64]float64{1.0: 1.0, 3.0: 1.0}
	s.RemoveFloat(2.0)

	if reflect.DeepEqual(s.FloatMap, testMap) == false {
		t.Errorf("RemoveFloat incorrect")
	}
}

func TestStringContains(t *testing.T) {
	s := InitTestEnv()
	if s.StringContains("a") == false {
		t.Errorf("StringContains incorrect")
	}
}

func TestIntContains(t *testing.T) {
	s := InitTestEnv()
	if s.IntContains(1) == false {
		t.Errorf("IntContains incorrect")
	}
}

func TestFloatContains(t *testing.T) {
	s := InitTestEnv()
	if s.FloatContains(1.0) == false {
		t.Errorf("FloatContains incorrect")
	}
}

func TestIntersectFloat(t *testing.T) {
	s := InitTestEnv()

	intersectSet := NewSetFromFloatMap(map[float64]float64{1.0: 1.0, 2.0: 1.0, 3.0: 1.0, 4.0: 1.0}, "floatmap")

	s.IntersectFloat(intersectSet, "testMap")

	if len(s.FloatMap) != 3 {
		t.Errorf("intersection cardinality incorrect")
	}
	testMap := map[float64]float64{1.0: 1.0, 2.0: 1.0, 3.0: 1.0}
	if reflect.DeepEqual(s.FloatMap, testMap) == false {
		t.Errorf("intersection incorrect")
	}
}
func TestIntersectString(t *testing.T) {
	s := InitTestEnv()

	intersectSet := NewSetFromStringMap(map[string]float64{"a": 1.0, "b": 1.0, "c": 1.0, "d": 1.0}, "stringmap")

	s.IntersectString(intersectSet, "testMap")

	if len(s.StringMap) != 3 {
		t.Errorf("intersection cardinality incorrect")
	}
	testMap := map[string]float64{"a": 1.0, "b": 1.0, "c": 1.0}
	if reflect.DeepEqual(s.StringMap, testMap) == false {
		t.Errorf("intersection incorrect")
	}
}
func TestIntersectInt(t *testing.T) {
	s := InitTestEnv()

	intersectSet := NewSetFromIntMap(map[int]float64{1: 1.0, 2: 1.0, 3: 1.0, 4: 1.0}, "intmap")

	s.IntersectInt(intersectSet, "testMap")

	if len(s.IntMap) != 3 {
		t.Errorf("intersection cardinality incorrect")
	}
	testMap := map[int]float64{1: 1.0, 2: 1.0, 3: 1.0}
	if reflect.DeepEqual(s.IntMap, testMap) == false {
		t.Errorf("intersection incorrect")
	}
}
func TestUnionString(t *testing.T) {
	s := InitTestEnv()
	testSet := NewSetFromStringMap(map[string]float64{"a": 1.0, "b": 1.0, "c": 1.0, "d": 1.0}, "stringmap")
	s.UnionString(testSet, "string_union")

	if reflect.DeepEqual(s.StringMap, testSet.StringMap) == false {
		t.Errorf("union incorrect")
	}
}

func TestUnionInt(t *testing.T) {
	s := InitTestEnv()
	testSet := NewSetFromIntMap(map[int]float64{1: 1.0, 2: 1.0, 3: 1.0, 4: 1.0}, "intmap")
	s.UnionInt(testSet, "int_union")

	if reflect.DeepEqual(s.IntMap, testSet.IntMap) == false {
		t.Errorf("union incorrect")
	}
}

func TestUnionFloat(t *testing.T) {
	s := InitTestEnv()
	testMap := NewSetFromFloatMap(map[float64]float64{1.0: 1.0, 2.0: 1.0, 3.0: 1.0, 4.0: 1.0}, "floatmap")
	s.UnionFloat(testMap, "float union")

	if reflect.DeepEqual(s.FloatMap, testMap.FloatMap) == false {
		t.Errorf("union incorrect")
	}
}
