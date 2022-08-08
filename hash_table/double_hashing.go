package hash_table

import "fmt"

type Person struct {
	name string
	age  int
}

type KeyValue struct {
	key   string
	value interface{}
}

type HashMap struct {
	size  int
	table []*KeyValue
}

func NewHashMap(size int) HashMap {
	return HashMap{
		size:  size,
		table: make([]*KeyValue, size),
	}
}

func (m *HashMap) findIndex(key string) (bool, int) {
	hash1 := hash(key)
	index1 := hash1 % m.size
	if m.table[index1] == nil {
		return false, index1
	}
	if m.table[index1].key == key {
		return true, index1
	}

	// we have a collision - let's start double hashing
	hash2 := hash(key + string('d'))
	stride := hash2%(m.size-1) + 1
	index2 := (index1 + stride) % m.size
	for index2 != index1 {
		if m.table[index2] == nil {
			return false, index2
		}
		if m.table[index2].key == key {
			return true, index2
		}
		index2 = (index2 + stride) % m.size
	}

	// table is full
	return false, -1
}

func (m *HashMap) Insert(key string, value interface{}) {
	found, index := m.findIndex(key)
	if index == -1 {
		fmt.Println("could not insert into the map because it is full")
		return
	}
	if found {
		m.table[index].value = value
	} else {
		m.table[index] = &KeyValue{key, value}
	}
}

func (m *HashMap) Search(key string) interface{} {
	found, index := m.findIndex(key)
	if found {
		return m.table[index].value
	}
	return nil
}

func hash(s string) int {
	hash := 5381
	for _, c := range s {
		hash = hash*33 + int(c)
	}
	return hash
}

func DoubleHashing() {
	// the size of the map must be a prime number
	hashMap := NewHashMap(7)

	p1 := Person{"Carlos", 36}
	p2 := Person{"Marcia", 39}
	p3 := Person{"Rebeca", 2}
	p4 := Person{"Isadora", 1}
	p5 := Person{"Silvio", 77}
	p6 := Person{"Marlene", 65}
	p7 := Person{"Belinha", 16}

	hashMap.Insert(p1.name, p1.age)
	fmt.Println(hashMap.table)
	hashMap.Insert(p2.name, p2.age)
	fmt.Println(hashMap.table)
	hashMap.Insert(p3.name, p3.age)
	fmt.Println(hashMap.table)
	hashMap.Insert(p4.name, p4.age)
	fmt.Println(hashMap.table)
	hashMap.Insert(p5.name, p5.age)
	fmt.Println(hashMap.table)
	hashMap.Insert(p6.name, p6.age)
	fmt.Println(hashMap.table)
	hashMap.Insert(p7.name, p7.age)
	fmt.Println(hashMap.table)

	fmt.Println(p1.name, hashMap.Search(p1.name))
	fmt.Println(p2.name, hashMap.Search(p2.name))
	fmt.Println(p3.name, hashMap.Search(p3.name))
	fmt.Println(p4.name, hashMap.Search(p4.name))
	fmt.Println(p5.name, hashMap.Search(p5.name))
	fmt.Println(p6.name, hashMap.Search(p6.name))
	fmt.Println(p7.name, hashMap.Search(p7.name))

	hashMap.Insert(p1.name, 999)
	fmt.Println(p1.name, hashMap.Search(p1.name))

	p8 := Person{"Renato", 34}
	hashMap.Insert(p8.name, p8.age)
	fmt.Println(p8.name, hashMap.Search(p8.name))
}
