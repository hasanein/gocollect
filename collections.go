package gocollect

import "fmt"

type Collection struct {
	values []interface{}
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

func (collection *Collection)unwrap() (underlyingArray []interface{}) {
	underlyingArray = collection.values
	return
}

func (collection *Collection) Print() interface{} {
	for _,v := range collection.values{
		fmt.Println(v)
	}
	return nil
}

func CollectionFrom(sliceOfAnything []interface{}) (*Collection)  {
	this := new(Collection)
	this.values = sliceOfAnything
	return this
}