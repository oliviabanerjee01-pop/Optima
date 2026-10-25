
package main

import "testing"

func TestPutAndGet(t *testing.T) {
    db := NewDatabase()
    db.Put("language", "Go")

    value, exists := db.Get("language")

    if !exists {
        t.Fatal("expected key to exist")
    }

    if value != "Go" {
        t.Fatalf("expected Go, got %q", value)
    }
}

func TestUpdate(t *testing.T) {
    db := NewDatabase()
    db.Put("language", "Go")
    db.Put("language", "Rust")

    value, _ := db.Get("language")

    if value != "Rust" {
        t.Fatalf("expected Rust, got %q", value)
    }
}

func TestDelete(t *testing.T) {
    db := NewDatabase()
    db.Put("project", "Optima")

    if !db.Delete("project") {
        t.Fatal("expected delete to return true")
    }

    _, exists := db.Get("project")

    if exists {
        t.Fatal("expected key to be absent")
    }
}

func TestMissingKey(t *testing.T) {
    db := NewDatabase()

    _, exists := db.Get("unknown")

    if exists {
        t.Fatal("expected missing key to return false")
    }
}
