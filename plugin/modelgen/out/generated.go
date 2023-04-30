// Code generated by github.com/willdhorn/gqlgen, DO NOT EDIT.

package out

import (
	"fmt"
	"io"
	"strconv"

	"github.com/willdhorn/gqlgen/graphql"
)

type A interface {
	IsA()
	GetA() string
}

type ArrayOfA interface {
	IsArrayOfA()
	GetTrickyField() []A
	GetTrickyFieldPointer() []A
}

type B interface {
	IsB()
	GetB() int
}

type C interface {
	IsA()
	IsC()
	GetA() string
	GetC() bool
}

type D interface {
	IsA()
	IsB()
	IsD()
	GetA() string
	GetB() int
	GetD() *string
}

type FooBarer interface {
	IsFooBarer()
	GetName() string
}

// InterfaceWithDescription is an interface with a description
type InterfaceWithDescription interface {
	IsInterfaceWithDescription()
	GetName() *string
}

type MissingInterface interface {
	IsMissingInterface()
	GetName() *string
}

type MissingUnion interface {
	IsMissingUnion()
}

// UnionWithDescription is an union with a description
type UnionWithDescription interface {
	IsUnionWithDescription()
}

type X interface {
	IsX()
	GetId() string
}

type CDImplemented struct {
	A string  `json:"a" database:"CDImplementeda"`
	B int     `json:"b" database:"CDImplementedb"`
	C bool    `json:"c" database:"CDImplementedc"`
	D *string `json:"d,omitempty" database:"CDImplementedd"`
}

func (CDImplemented) IsC()              {}
func (this CDImplemented) GetA() string { return this.A }
func (this CDImplemented) GetC() bool   { return this.C }

func (CDImplemented) IsA() {}

func (CDImplemented) IsD() {}

func (this CDImplemented) GetB() int     { return this.B }
func (this CDImplemented) GetD() *string { return this.D }

func (CDImplemented) IsB() {}

type CyclicalA struct {
	FieldOne   *CyclicalB `json:"field_one,omitempty" database:"CyclicalAfield_one"`
	FieldTwo   *CyclicalB `json:"field_two,omitempty" database:"CyclicalAfield_two"`
	FieldThree *CyclicalB `json:"field_three,omitempty" database:"CyclicalAfield_three"`
	FieldFour  string     `json:"field_four" database:"CyclicalAfield_four"`
}

type CyclicalB struct {
	FieldOne   *CyclicalA `json:"field_one,omitempty" database:"CyclicalBfield_one"`
	FieldTwo   *CyclicalA `json:"field_two,omitempty" database:"CyclicalBfield_two"`
	FieldThree *CyclicalA `json:"field_three,omitempty" database:"CyclicalBfield_three"`
	FieldFour  *CyclicalA `json:"field_four,omitempty" database:"CyclicalBfield_four"`
	FieldFive  string     `json:"field_five" database:"CyclicalBfield_five"`
}

type FieldMutationHook struct {
	Name     *string       `json:"name,omitempty" anotherTag:"tag" database:"FieldMutationHookname"`
	Enum     *ExistingEnum `json:"enum,omitempty" yetAnotherTag:"12" database:"FieldMutationHookenum"`
	NoVal    *string       `json:"noVal,omitempty" yaml:"noVal" repeated:"true" database:"FieldMutationHooknoVal"`
	Repeated *string       `json:"repeated,omitempty" someTag:"value" repeated:"true" database:"FieldMutationHookrepeated"`
}

type ImplArrayOfA struct {
	TrickyField        []*CDImplemented `json:"trickyField" database:"ImplArrayOfAtrickyField"`
	TrickyFieldPointer []*CDImplemented `json:"trickyFieldPointer,omitempty" database:"ImplArrayOfAtrickyFieldPointer"`
}

func (ImplArrayOfA) IsArrayOfA() {}
func (this ImplArrayOfA) GetTrickyField() []A {
	if this.TrickyField == nil {
		return nil
	}
	interfaceSlice := make([]A, 0, len(this.TrickyField))
	for _, concrete := range this.TrickyField {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}
func (this ImplArrayOfA) GetTrickyFieldPointer() []A {
	if this.TrickyFieldPointer == nil {
		return nil
	}
	interfaceSlice := make([]A, 0, len(this.TrickyFieldPointer))
	for _, concrete := range this.TrickyFieldPointer {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type MissingInput struct {
	Name          *string                           `json:"name,omitempty" database:"MissingInputname"`
	Enum          *MissingEnum                      `json:"enum,omitempty" database:"MissingInputenum"`
	NonNullString string                            `json:"nonNullString" database:"MissingInputnonNullString"`
	NullString    graphql.Omittable[*string]        `json:"nullString,omitempty" database:"MissingInputnullString"`
	NullEnum      graphql.Omittable[*MissingEnum]   `json:"nullEnum,omitempty" database:"MissingInputnullEnum"`
	NullObject    graphql.Omittable[*ExistingInput] `json:"nullObject,omitempty" database:"MissingInputnullObject"`
}

type MissingTypeNotNull struct {
	Name     string               `json:"name" database:"MissingTypeNotNullname"`
	Enum     MissingEnum          `json:"enum" database:"MissingTypeNotNullenum"`
	Int      MissingInterface     `json:"int" database:"MissingTypeNotNullint"`
	Existing *ExistingType        `json:"existing" database:"MissingTypeNotNullexisting"`
	Missing2 *MissingTypeNullable `json:"missing2" database:"MissingTypeNotNullmissing2"`
}

func (MissingTypeNotNull) IsMissingInterface()   {}
func (this MissingTypeNotNull) GetName() *string { return &this.Name }

func (MissingTypeNotNull) IsExistingInterface() {}

func (MissingTypeNotNull) IsMissingUnion() {}

func (MissingTypeNotNull) IsExistingUnion() {}

type MissingTypeNullable struct {
	Name     *string             `json:"name,omitempty" database:"MissingTypeNullablename"`
	Enum     *MissingEnum        `json:"enum,omitempty" database:"MissingTypeNullableenum"`
	Int      MissingInterface    `json:"int,omitempty" database:"MissingTypeNullableint"`
	Existing *ExistingType       `json:"existing,omitempty" database:"MissingTypeNullableexisting"`
	Missing2 *MissingTypeNotNull `json:"missing2,omitempty" database:"MissingTypeNullablemissing2"`
}

func (MissingTypeNullable) IsMissingInterface()   {}
func (this MissingTypeNullable) GetName() *string { return this.Name }

func (MissingTypeNullable) IsExistingInterface() {}

func (MissingTypeNullable) IsMissingUnion() {}

func (MissingTypeNullable) IsExistingUnion() {}

type NotCyclicalA struct {
	FieldOne string `json:"FieldOne" database:"NotCyclicalAFieldOne"`
	FieldTwo int    `json:"FieldTwo" database:"NotCyclicalAFieldTwo"`
}

type NotCyclicalB struct {
	FieldOne string        `json:"FieldOne" database:"NotCyclicalBFieldOne"`
	FieldTwo *NotCyclicalA `json:"FieldTwo" database:"NotCyclicalBFieldTwo"`
}

type Recursive struct {
	FieldOne   *Recursive `json:"FieldOne" database:"RecursiveFieldOne"`
	FieldTwo   *Recursive `json:"FieldTwo" database:"RecursiveFieldTwo"`
	FieldThree *Recursive `json:"FieldThree" database:"RecursiveFieldThree"`
	FieldFour  string     `json:"FieldFour" database:"RecursiveFieldFour"`
}

type RenameFieldTest struct {
	GOODnaME   string `json:"badName" database:"RenameFieldTestbadName"`
	OtherField string `json:"otherField" database:"RenameFieldTestotherField"`
}

// TypeWithDescription is a type with a description
type TypeWithDescription struct {
	Name *string `json:"name,omitempty" database:"TypeWithDescriptionname"`
}

func (TypeWithDescription) IsUnionWithDescription() {}

type Xer struct {
	Id   string `json:"Id" database:"XerId"`
	Name string `json:"Name" database:"XerName"`
}

func (Xer) IsX()               {}
func (this Xer) GetId() string { return this.Id }

type FooBarr struct {
	Name string `json:"name" database:"_Foo_Barrname"`
}

func (FooBarr) IsFooBarer()          {}
func (this FooBarr) GetName() string { return this.Name }

// EnumWithDescription is an enum with a description
type EnumWithDescription string

const (
	EnumWithDescriptionCat EnumWithDescription = "CAT"
	EnumWithDescriptionDog EnumWithDescription = "DOG"
)

var AllEnumWithDescription = []EnumWithDescription{
	EnumWithDescriptionCat,
	EnumWithDescriptionDog,
}

func (e EnumWithDescription) IsValid() bool {
	switch e {
	case EnumWithDescriptionCat, EnumWithDescriptionDog:
		return true
	}
	return false
}

func (e EnumWithDescription) String() string {
	return string(e)
}

func (e *EnumWithDescription) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnumWithDescription(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnumWithDescription", str)
	}
	return nil
}

func (e EnumWithDescription) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MissingEnum string

const (
	MissingEnumHello   MissingEnum = "Hello"
	MissingEnumGoodbye MissingEnum = "Goodbye"
)

var AllMissingEnum = []MissingEnum{
	MissingEnumHello,
	MissingEnumGoodbye,
}

func (e MissingEnum) IsValid() bool {
	switch e {
	case MissingEnumHello, MissingEnumGoodbye:
		return true
	}
	return false
}

func (e MissingEnum) String() string {
	return string(e)
}

func (e *MissingEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MissingEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MissingEnum", str)
	}
	return nil
}

func (e MissingEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
