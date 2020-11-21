package gqlhelper

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"testing"
)
func wrap(s string) *string {
	return &s
}

func trust(v interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return v
}
func TestQuery_AddID(t *testing.T) {
	type fields struct {
		q bson.M
	}
	type args struct {
		name  string
		value *string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Query
	}{
		{
			"add id",
			fields{q: bson.M{}},
			args{"test", wrap("5fb4d75f6a2a6db97ace7229")},
			&Query{q: bson.M{"test": trust(
				primitive.ObjectIDFromHex("5fb4d75f6a2a6db97ace7229"))}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Query := &Query{
				q: tt.fields.q,
			}
			if got := Query.AddID(tt.args.name, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddID() = %v, want %v", got, tt.want)
			}
		})
	}
}
