package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	ldb "github.com/syndtr/goleveldb/leveldb"
	"log"
)

const(
	path = "database.db"
)

type Person struct {
	Name string
	Age string
	Gender string
}

func main() {
	db := createdb()
	A := testPerson()


	Encode(A, db)
	decodedperson := getperson("Brandon", db) //decode("Brandon", db)
	fmt.Println(decodedperson)
}

func testPerson() *Person {
	A := Person{"Brandon", "18", "Male"}
	return &A
}

func createdb() *ldb.DB{
	db, err := ldb.OpenFile(path, nil)
	if err != nil {
		log.Panicln(err)
	}

	return db
}

func Encode(chosenperson *Person, db *ldb.DB) {
	err := db.Put(
		[]byte(chosenperson.Name),
		chosenperson.Serialize(),
		nil)
	if err != nil {log.Panic(err)}
}

func getperson(name string, db *ldb.DB) *Person{
	rawdata, err := db.Get([]byte(name), nil)
	if err != nil {log.Panic(err)}

	finish := Deserialize(rawdata)
	return finish
}

func (p *Person) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(p)
	if err != nil {log.Panic(err)}

	return result.Bytes()
}

func Deserialize(d []byte) *Person {
	var block Person

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {log.Panic(err)}

	return &block
}