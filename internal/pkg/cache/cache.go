package cache

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	L1         map[string]*list.Element
	L1List     *list.List
	L2         map[string]*list.Element
	L2List     *list.List
	MaxSizeL1  int
	MaxSizeL2  int
	AccessTime time.Duration
	Lock       sync.Mutex
}

type CacheItem struct {
	Key       string
	Value     string
	AccessCnt int
}

func NewCache(maxSizeL1, maxSizeL2 int, accessTime time.Duration) (*Cache, error) {
	return &Cache{
		L1:         make(map[string]*list.Element),
		L1List:     list.New(),
		L2:         make(map[string]*list.Element),
		L2List:     list.New(),
		MaxSizeL1:  maxSizeL1,
		MaxSizeL2:  maxSizeL2,
		AccessTime: accessTime,
	}, nil
}

func (c *Cache) Get(key string) interface{} {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	if item, ok := c.L1[key]; ok {
		c.promoteL1(item)
		return item.Value.(*CacheItem).Value
	}

	if item, ok := c.L2[key]; ok {
		if time.Since(time.Unix(int64(item.Value.(*CacheItem).AccessCnt), 0)) < c.AccessTime {
			c.removeL2(item)
			return nil
		}
		item.Value.(*CacheItem).AccessCnt++
		c.promoteL2(item)
		return item.Value.(*CacheItem).Value
	}

	return nil
}

func (c *Cache) Set(key string, value string) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	if item, ok := c.L1[key]; ok {
		item.Value.(*CacheItem).Value = value
		c.promoteL1(item)
		return
	}

	if item, ok := c.L2[key]; ok {
		item.Value.(*CacheItem).Value = value
		item.Value.(*CacheItem).AccessCnt = 0
		c.promoteL2(item)
		return
	}

	item := &CacheItem{
		Key:       key,
		Value:     value,
		AccessCnt: 0,
	}

	if len(c.L1) < c.MaxSizeL1 {
		c.addToL1(item)
		return
	}

	if len(c.L2) < c.MaxSizeL2 {
		c.addToL2(item)
		return
	}

	c.evictL1()
	c.addToL1(item)
}

func (c *Cache) Print() {
	fmt.Println("Cache Contents:")
	for key, item := range c.L1 {
		fmt.Printf("L1 - Key: %s, Value: %v\n", key, item.Value.(*CacheItem).Value)
	}
	for key, item := range c.L2 {
		fmt.Printf("L2 - Key: %s, Value: %v\n", key, item.Value.(*CacheItem).Value)
	}
}
