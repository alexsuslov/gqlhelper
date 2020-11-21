package gqlhelper

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Query struct {
	q bson.M
}

func NewQuery() *Query {
	return &Query{
		bson.M{},
	}
}
func (Query *Query) M() bson.M {
	return Query.q
}
func (Query *Query) Add(name string, value *string) *Query {
	if value != nil {
		Query.q[name] = value
	}
	return Query
}

func (Query *Query) AddID(name string, value *string) *Query {
	if value != nil {
		id, err := primitive.ObjectIDFromHex(*value)
		if err == nil {
			Query.q[name] = id
		}
	}
	return Query
}

type FindOptions struct {
	o *options.FindOptions
}

func NewFindOptions() *FindOptions {
	return &FindOptions{
		&options.FindOptions{},
	}
}

func (FindOptions *FindOptions) FindOptions() *options.FindOptions {
	return FindOptions.o
}

func (FindOptions *FindOptions) SetLimit(i *int) *FindOptions {
	if i != nil {
		L := int64(*i)
		FindOptions.o.Limit = &L
	}
	return FindOptions
}

func (FindOptions *FindOptions) SetSkip(i *int) *FindOptions {
	if i != nil {
		L := int64(*i)
		FindOptions.o.Skip = &L
	}
	return FindOptions
}

func (FindOptions *FindOptions) SetOrder(order *string) *FindOptions {
	if order != nil {
		name := *order
		Order := bson.M{name: 1}
		if name[0] == '!' {
			Order[name[1:]] = -1
		}
		FindOptions.o.Sort = Order
	}
	return FindOptions
}

func ObjectID(ID *primitive.ObjectID) (string, error) {
	if ID == nil {
		return "", fmt.Errorf("null ObjectID")

	}
	return ID.Hex(), nil
}

func ObjectIDPoint(ID *primitive.ObjectID) (p *string, err error) {
	if ID == nil {
		err = fmt.Errorf("null ObjectID")
	} else {
		*p = ID.Hex()
	}
	return
}