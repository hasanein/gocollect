package gocollect

import "fmt"

type Collection struct {
	values []interface{}
}

// Creates a collection from a slice
func CollectionFromSlice(sliceOfAnything []interface{}) (*Collection)  {
	this := new(Collection)
	this.values = sliceOfAnything
	return this
}

func CollectionFromElement(element interface{}) (*Collection)  {
	this := new(Collection)
	this.values = append(this.values, element)
	return this
}

func (collection *Collection) Filter(predicate func(value interface{}) bool) *Collection {
	filtered := make([]interface{}, 0)
	for _,v := range collection.values{
		if predicate(v){
			 filtered = append(filtered, v)
		}
	}
	collection.values = filtered
	return collection
}

func (collection *Collection) Map(mapper func(value interface{}) interface{}) *Collection {
	mapped := make([]interface{}, 0)
	for _,v := range collection.values{
		mapped = append(mapped, mapper(v))
	}
	collection.values = mapped
	return collection
}

func (collection *Collection) Count() int {
	sum := 0
	for i := 0; i< len(collection.values); i++{
		sum++
	}
	return sum
}

func (collection *Collection) Unwrap() []interface{} {
	return collection.values
}

// TODO decide if we should keep this or make it more generic by
// simply allowing the support for for-each style
func (collection *Collection) Print() interface{} {
	for _,v := range collection.values{
		fmt.Println(v)
	}
	return nil
}

func (collection *Collection) ForEach(consumer func(value interface{})) {
	for _,v := range collection.values{
		consumer(v)
	}
}

func (collection *Collection) Peek(consumer func(value interface{})) *Collection {
	for _,v := range collection.values{
		consumer(v)
	}
	return collection
}

func (collection *Collection) FindFirst() (interface{}, bool) {
	if len(collection.values) != 0{
		return collection.values[0], true
	}
	return nil, false
}