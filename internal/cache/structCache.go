package cache

import "container/list"

func (c *Cache) promoteL1(item *list.Element) {
	c.L1List.MoveToFront(item)
}

func (c *Cache) promoteL2(item *list.Element) {
	c.L2List.MoveToFront(item)
}

func (c *Cache) removeL2(item *list.Element) {
	c.L2List.Remove(item)
	delete(c.L2, item.Value.(*CacheItem).Key)
}

func (c *Cache) addToL1(item *CacheItem) {
	e := c.L1List.PushFront(item)
	c.L1[item.Key] = e
}

func (c *Cache) addToL2(item *CacheItem) {
	e := c.L2List.PushFront(item)
	c.L2[item.Key] = e
}

func (c *Cache) evictL1() {
	item := c.L1List.Back()
	c.L1List.Remove(item)
	delete(c.L1, item.Value.(*CacheItem).Key)
}
