package main

import "fmt"

type KeyValue struct {
	key   string
	value interface{}
}

type HashTable struct {
	table map[int][]KeyValue
}

func NewHashTable() *HashTable {
	return &HashTable{
		table: make(map[int][]KeyValue),
	}
}

func hashFunction(key string) (hash int) {
	for _, char := range key {
		hash += int(char)
	}
	return hash
}

func (ht *HashTable) Insert(key string, value interface{}) {
	index := hashFunction(key)
	ht.table[index] = append(ht.table[index], KeyValue{key: key, value: value})
}

func (ht *HashTable) Get(key string) (value interface{}) {
	index := hashFunction(key)
	for _, keyValue := range ht.table[index] {
		if keyValue.key == key {
			return keyValue.value
		}
	}
	return nil
}

func (ht *HashTable) Remove(key string) {
	index := hashFunction(key)

	for _, keyValue := range ht.table[index] {
		if keyValue.key == key {
			delete(ht.table, index)
		}
	}
}

func (ht *HashTable) Keys() {
	for index, _ := range ht.table {
		for _, key := range ht.table[index] {
			fmt.Println(key.key)
		}
	}
}

func (ht *HashTable) Values() {
	for index, _ := range ht.table {
		for _, value := range ht.table[index] {
			fmt.Println(value.value)
		}
	}
}
