package lru

import (
	"container/list"
)

type Cache struct {
	maxBytes int64
	nBytes   int64
	//双向链表
	ll *list.List
	//键是字符串，值是双向链表中对应节点的指针。
	cache map[string]*list.Element
	//是某条记录被移除时的回调函数，可以为 nil。
	OnEvicted func(key string, value Value)
}

// 键值对 entry 是双向链表节点的数据类型，在链表中仍保存每个值对应的 key 的好处在于，淘汰队首节点时，需要用 key 从字典中删除对应的映射。
type entry struct {
	key   string
	value Value
}

// Value 为了通用性，我们允许值是实现了 Value 接口的任意类型，该接口只包含了一个方法 Len() int，用于返回值所占用的内存大小。
type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		nBytes:    0,
		ll:        list.New(),
		cache:     map[string]*list.Element{},
		OnEvicted: onEvicted,
	}
}

// Get 查找功能
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOldest c.ll.Back() 取到队首节点，从链表中删除。
// delete(c.cache, kv.key)，从字典中 c.cache 删除该节点的映射关系。
// 更新当前所用的内存 c.nBytes。
// 如果回调函数 OnEvicted 不为 nil，则调用回调函数。
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key)) - int64(kv.value.Len())
		//调用
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{
			key:   key,
			value: value,
		})
		c.cache[key] = ele
		c.nBytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.RemoveOldest()
	}
	return
}

// Len the number of cache entries
func (c *Cache) Len() int {
	return c.ll.Len()
}
