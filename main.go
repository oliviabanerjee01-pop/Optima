
package main

import "fmt"

type Database struct {
    data map[string]string
}

func NewDatabase() *Database {
    return &Database{
        data: make(map[string]string),
    }
}

func (db *Database) Put(key, value string) {
    db.data[key] = value
}

func (db *Database) Get(key string) (string, bool) {
    value, exists := db.data[key]
    return value, exists
}

func (db *Database) Delete(key string) bool {
    _, exists := db.data[key]

    if exists {
        delete(db.data, key)
    }

    return exists
}

func main() {
    db := NewDatabase()

    db.Put("language", "Go")
    db.Put("project", "Optima")

    value, exists := db.Get("language")
    fmt.Println("Value:", value, "Exists:", exists)

    deleted := db.Delete("project")
    fmt.Println("Deleted:", deleted)

    _, exists = db.Get("project")
    fmt.Println("Project still exists:", exists)
}




