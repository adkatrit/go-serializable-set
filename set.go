/*
	Author: Jarrod Parker
*/
package set

import "errors"

type Set struct {
	Name      string
	IntMap    map[int]float64
	FloatMap  map[float64]float64
	StringMap map[string]float64
}

func (s *Set) AddString(str string) (err error) {
	if _, ok := s.StringMap[str]; ok {
		err = errors.New("string value already exists in " + s.Name)
	} else {
		err = nil
		s.StringMap[str] = 1.0
	}
	return err
}

func (s *Set) RemoveString(str string) (err error) {
	if _, ok := s.StringMap[str]; ok {
		delete(s.StringMap, str)
		err = nil
	} else {
		err = errors.New("string value doesn't exists in " + s.Name)
	}
	return err
}

func (s *Set) StringContains(str string) bool {
	ok := false
	if _, ok = s.StringMap[str]; ok {
	}
	return ok
}

func (s *Set) AddInt(i int) (err error) {
	if _, ok := s.IntMap[i]; ok {
		err = errors.New("integer value already exists in " + s.Name)
	} else {
		err = nil
		s.IntMap[i] = 1.0
	}
	return err
}

func (s *Set) RemoveInt(i int) (err error) {
	if _, ok := s.IntMap[i]; ok {
		delete(s.IntMap, i)
		err = nil
	} else {
		err = errors.New("int value doesn't exists in " + s.Name)
	}
	return err
}

func (s *Set) IntContains(i int) bool {
	ok := false
	if _, ok = s.IntMap[i]; ok {
	}
	return ok
}

func (s *Set) AddFloat(f float64) (err error) {
	if _, ok := s.FloatMap[f]; ok {
		err = errors.New("float value already exists in " + s.Name)
	} else {
		err = nil
		s.FloatMap[f] = 1.0
	}
	return err
}

func (s *Set) FloatContains(f float64) bool {
	ok := false
	if _, ok = s.FloatMap[f]; ok {
	}
	return ok
}

func (s *Set) RemoveFloat(f float64) (err error) {
	if _, ok := s.FloatMap[f]; ok {
		delete(s.FloatMap, f)
		err = nil
	} else {
		err = errors.New("float value doesn't exists in " + s.Name)
	}
	return err
}

func (s *Set) IntersectString(f *Set, name string) *Set {
	k := NewSet(name)

	for i := range f.StringMap {
		if s.StringContains(i) {
			k.AddString(i)
		}
	}
	s.StringMap = k.StringMap

	return s
}

func (s *Set) IntersectInt(f *Set, name string) *Set {
	k := NewSet(name)

	for i := range f.IntMap {
		if s.IntContains(i) {
			k.AddInt(i)
		}
	}
	s.IntMap = k.IntMap

	return s
}

func (s *Set) IntersectFloat(f *Set, name string) *Set {
	k := NewSet(name)

	for i := range f.FloatMap {
		if s.FloatContains(i) {
			k.AddFloat(i)
		}
	}
	s.FloatMap = k.FloatMap

	return s
}

func (s *Set) UnionString(str *Set, name string) *Set {
	for i := range str.StringMap {
		s.AddString(i)
	}
	return s
}

func (s *Set) UnionInt(in *Set, name string) *Set {
	for i := range in.IntMap {
		s.AddInt(i)
	}
	return s
}

func (s *Set) UnionFloat(f *Set, name string) *Set {
	for i := range f.FloatMap {
		s.AddFloat(i)
	}
	return s
}
func NewSetFromStringMap(m map[string]float64, name string) *Set {
	return &Set{
		Name:      name,
		StringMap: m,
		IntMap:    map[int]float64{},
		FloatMap:  map[float64]float64{},
	}
}
func NewSetFromIntMap(m map[int]float64, name string) *Set {
	return &Set{
		Name:      name,
		StringMap: map[string]float64{},
		IntMap:    m,
		FloatMap:  map[float64]float64{},
	}
}
func NewSetFromFloatMap(m map[float64]float64, name string) *Set {
	return &Set{
		Name:      name,
		StringMap: map[string]float64{},
		IntMap:    map[int]float64{},
		FloatMap:  m,
	}
}

func NewSet(name string) *Set {
	set := Set{
		Name:      name,
		StringMap: map[string]float64{},
		IntMap:    map[int]float64{},
		FloatMap:  map[float64]float64{},
	}
	return &set
}
