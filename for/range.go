for key, value := range oldMap {
	newMap[key] = value
}

for key := range m {
	if key.expired() {
		delete(m, key)
	}
}

sum := 0
for _, value := range array {
	sum += value
}

for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}