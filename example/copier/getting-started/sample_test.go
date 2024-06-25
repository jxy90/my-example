package main

import (
	"fmt"
	"github.com/jinzhu/copier"
	"strconv"
	"testing"
)

type Address struct {
	City  string
	State string
}

// 简单复制
type Source struct {
	Name    string
	Age     int
	Email   string
	Address Address
	FName   string
}

type Destination struct {
	Name     string
	Age      int
	Email    string
	Address  Address
	FullName string
}

func Test_sample_copy(t *testing.T) {
	source := Source{
		Name:  "John Doe",
		Age:   30,
		Email: "john.doe@example.com",
	}

	var destination Destination

	err := copier.Copy(&destination, &source)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("Source: %+v\n", source)
	fmt.Printf("Destination: %+v\n", destination)
}

// 忽略空值
func Test_ignore_empty(t *testing.T) {
	source := Source{
		Name:  "",
		Age:   30,
		Email: "john.doe@example.com",
	}

	destination := Destination{
		Name: "Default Name",
	}

	err := copier.CopyWithOption(&destination, &source, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return
	}

	fmt.Printf("Source: %+v\n", source)
	fmt.Printf("Destination: %+v\n", destination)
}

// 深度复制
func Test_deep(t *testing.T) {
	source := Source{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}

	var destination Destination

	err := copier.CopyWithOption(&destination, &source, copier.Option{DeepCopy: true})
	if err != nil {
		return
	}

	fmt.Printf("Source: %#v\n", source)
	fmt.Printf("Destination: %#v\n", destination)
}

// 自定义转换器
type Source2 struct {
	Name string
	Age  int
}

type Destination2 struct {
	Name string
	Age  string
}

func Test_custom(t *testing.T) {
	source := Source2{
		Name: "John Doe",
		Age:  30,
	}

	var destination Destination2

	err := copier.CopyWithOption(&destination, &source, copier.Option{
		Converters: []copier.TypeConverter{
			{
				SrcType: 0,
				DstType: "",
				Fn: func(src interface{}) (interface{}, error) {
					if str, ok := src.(int); ok {
						return strconv.Itoa(str), nil
					}
					return src, nil
				},
			},
		},
		IgnoreEmpty: false,
		DeepCopy:    false,
	})

	if err != nil {
		t.Errorf("Copy failed: %v", err)
	}

	fmt.Printf("Source: %+v\n", source)
	fmt.Printf("Destination: %+v\n", destination)

}
