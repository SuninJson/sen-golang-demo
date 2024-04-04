package main

func main() {
	m := make(map[string]string)
	m["key1"] = "value1"
	m["key2"] = "value2"

	// 定期清理不再使用的键值对
	delete(m, "key1")
}
