package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

// ComplexStruct 是一个复杂的结构体示例
type ComplexStruct struct {
	ID        int
	Name      string
	Timestamp time.Time
	Details   struct {
		Description string
		Attributes  map[string]interface{}
	}
	Items []struct {
		ItemID   int
		ItemName string
		Quantity int
	}
}

// 生成一个复杂的结构体实例
func generateComplexStruct() ComplexStruct {
	return ComplexStruct{
		ID:        1,
		Name:      "Test Struct",
		Timestamp: time.Now(),
		Details: struct {
			Description string
			Attributes  map[string]interface{}
		}{
			Description: "This is a detailed description",
			Attributes: map[string]interface{}{
				"key1": "value1",
				"key2": 42,
				"key3": []string{"a", "b", "c"},
			},
		},
		Items: []struct {
			ItemID   int
			ItemName string
			Quantity int
		}{
			{ItemID: 101, ItemName: "Item 1", Quantity: 10},
			{ItemID: 102, ItemName: "Item 2", Quantity: 20},
		},
	}
}

func BenchmarkFmtPrintf(b *testing.B) {
	cs := generateComplexStruct()
	for i := 0; i < b.N; i++ {
		fmt.Printf("print %+v\n", cs)
	}
}

func BenchmarkJSONMarshal(b *testing.B) {
	cs := generateComplexStruct()
	for i := 0; i < b.N; i++ {
		str, _ := json.Marshal(cs)
		fmt.Printf("json %s\n", str)
	}
}

func main() {
	// 运行基准测试
	testing.Benchmark(BenchmarkFmtPrintf)
	testing.Benchmark(BenchmarkJSONMarshal)
}
