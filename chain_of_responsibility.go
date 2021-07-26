package main

// 函数式实现的一个责任链模式

type Handler func(string) (string, bool)

func MemoryGetValue(key string) (string, bool) {
	var value string
	var hasValue bool

	// try get value from memory cache
	return value, hasValue
}

func RedisGetValue(key string) (string, bool) {
	var value string
	var hasValue bool

	// try get value from redis
	return value, hasValue
}

func MysqlGetValue(key string) (string, bool) {
	var value string
	var hasValue bool

	// try get value from mysql database
	return value, hasValue
}

func GetValueByChain(key string, functions ...Handler) string {
	for _, f := range functions {
		value, ok := f(key)
		if ok {
			return value
		}
	}
	return ""
}

func main() {
	GetValueByChain("key", MemoryGetValue, RedisGetValue, MysqlGetValue)
}
