package main

import (
	"encoding/json"
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"gopkg.in/yaml.v3"
	"time"
)

type DateTime int64

func (t DateTime) MarshalJSON() ([]byte, error) {
	sec := int64(t)
	if sec == 0 {
		return json.Marshal("")
	}
	timeValue := time.Unix(sec, 0)
	stamp := timeValue.Format(time.RFC3339)
	return json.Marshal(stamp)
}

func (t *DateTime) UnmarshalJSON(data []byte) error {
	var timeStr interface{}
	err := json.Unmarshal(data, &timeStr)
	if err != nil {
		return err
	}
	if timeStr == "" {
		sec := DateTime(0)
		t = &sec
		return nil
	}
	parsedTime, err := time.Parse(time.RFC3339, timeStr.(string))
	if err != nil {
		return err
	}
	sec := DateTime(parsedTime.Unix())
	t = &sec
	return nil
}

type Model struct {
	ID         int64    `gorm:"primary_key" json:"id"`
	CreateTime DateTime `json:"createTime"`
	UpdateTime DateTime `json:"updateTime"`
}
type Token struct {
	Model
	Value       string `json:"value"`
	Expire      int64  `json:"expire"`
	Description string `json:"description"`
	UserID      int64  `json:"userId"`
	OneOff      bool   `json:"oneOff"`
}
type Data struct {
	Name  string
	Value int
}

func main() {
	data := Data{Name: "example", Value: 42}

	// Test MsgPack serialization
	start := time.Now()
	for i := 0; i < 100000; i++ {
		_, err := msgpack.Marshal(data)
		if err != nil {
			fmt.Println("MsgPack error:", err)
		}
	}
	fmt.Printf("MsgPack serialization took %v\n", time.Since(start))

	// Test YAML serialization
	start = time.Now()
	for i := 0; i < 100000; i++ {
		_, err := yaml.Marshal(data)
		if err != nil {
			fmt.Println("YAML error:", err)
		}
	}
	fmt.Printf("YAML serialization took %v\n", time.Since(start))
	t := Token{
		Model: Model{
			ID:         123,
			CreateTime: 1725542444,
			UpdateTime: 1725542444,
		},
		Value:       "123",
		Expire:      123,
		Description: "123",
		UserID:      123,
		OneOff:      false,
	}
	str, _ := msgpack.Marshal(t)
	fmt.Println(string(str))

	token := &Token{}
	err := msgpack.Unmarshal(str, token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
}
