// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	structpb "github.com/golang/protobuf/ptypes/struct"
)

// NewEmptyDynValue returns the zero-valued DynValue.
func NewEmptyDynValue() *DynValue {
	// note: 0 is not a valid parse node identifier.
	return NewDynValue(0, nil)
}

// NewDynValue returns a DynValue that corresponds to a parse node id and value.
func NewDynValue(id int64, val ValueNode) *DynValue {
	return &DynValue{ID: id, Value: val}
}

// DynValue is a dynamically typed value used to describe unstructured content.
// Whether the value has the desired type is determined by where it is used within the Instance or
// Template, and whether there are schemas which might enforce a more rigid type definition.
type DynValue struct {
	ID    int64
	Value ValueNode
}

// ValueNode is a marker interface used to indicate which value types may populate a DynValue's
// Value field.
type ValueNode interface {
	isValueNode()
}

// NewStructValue returns an empty StructValue.
func NewStructValue() *StructValue {
	return &StructValue{
		Fields: []*StructField{},
	}
}

// StructValue declares an object with a set of named fields whose values are dynamically typed.
type StructValue struct {
	Fields []*StructField
}

func (*StructValue) isValueNode() {}

// NewStructField returns a StructField instance with an empty DynValue that refers to the
// specified parse node id and field name.
func NewStructField(id int64, name string) *StructField {
	return &StructField{
		ID:   id,
		Name: name,
		Ref:  NewEmptyDynValue(),
	}
}

// StructField specifies a field name and a reference to a dynamic value.
type StructField struct {
	ID   int64
	Name string
	Ref  *DynValue
}

// NewListValue returns an empty ListValue instance.
func NewListValue() *ListValue {
	return &ListValue{
		Entries: []*DynValue{},
	}
}

// ListValue contains a list of dynamically typed entries.
type ListValue struct {
	Entries []*DynValue
}

func (*ListValue) isValueNode() {}

// BoolValue is a boolean value suitable for use within DynValue objects.
type BoolValue bool

func (BoolValue) isValueNode() {}

// BytesValue is a []byte value suitable for use within DynValue objects.
type BytesValue []byte

func (BytesValue) isValueNode() {}

// DoubleValue is a float64 value suitable for use within DynValue objects.
type DoubleValue float64

func (DoubleValue) isValueNode() {}

// IntValue is an int64 value suitable for use within DynValue objects.
type IntValue int64

func (IntValue) isValueNode() {}

// NullValue is a protobuf.Struct concrete null value suitable for use within DynValue objects.
type NullValue structpb.NullValue

func (NullValue) isValueNode() {}

// StringValue is a string value suitable for use within DynValue objects.
type StringValue string

func (StringValue) isValueNode() {}

// UintValue is a uint64 value suitable for use within DynValue objects.
type UintValue uint64

func (UintValue) isValueNode() {}

// Null is a singleton NullValue instance.
var Null NullValue
