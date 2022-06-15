package pkg

import (
	"reflect"
	"testing"
)

func TestValidateStruct(t *testing.T) {
	ValidateStruct()
}

func TestEntryValue(t *testing.T) {
	u :=User{Aa:&Person{}}
	entryValue(reflect.ValueOf(&u))
}

func TestGetTag(t *testing.T) {
	u :=&User{Aa:&Person{}}
	FindTag("User.Aa.Msg",u)
}

//func TestForeachStruct(t *testing.T) {
//	u2 := User{
//		Name: "adfadfd",
//		Age:  15,
//		AA:   []string{""},
//		Aa:   Person{},
//	}
//	ForeachStruct(&u2)
//}
