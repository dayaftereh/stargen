package accretion

import "container/list"

func Find(collection *list.List, fn func(value interface{}, index int64) bool) *list.Element {
	index := int64(0)
	for e := collection.Front(); e != nil; e = e.Next() {
		found := fn(e.Value, index)
		if found {
			return e
		}
		index++
	}
	return nil
}

func FindByElement(collection *list.List, fn func(element *list.Element, index int64) bool) *list.Element {
	index := int64(0)
	for e := collection.Front(); e != nil; e = e.Next() {
		found := fn(e, index)
		if found {
			return e
		}
		index++
	}
	return nil
}

func FindAfter(last *list.Element, fn func(value interface{}) bool) *list.Element {
	for e := last.Next(); e != nil; e = e.Next() {
		found := fn(e.Value)
		if found {
			return e
		}
	}
	return nil
}

func ForEachElement(collection *list.List, fn func(element *list.Element, index int64)) {
	index := int64(0)
	for e := collection.Front(); e != nil; e = e.Next() {
		fn(e, index)
		index++
	}
}

func ForEach(collection *list.List, fn func(value interface{}, index int64)) {
	index := int64(0)
	for e := collection.Front(); e != nil; e = e.Next() {
		fn(e.Value, index)
		index++
	}
}
