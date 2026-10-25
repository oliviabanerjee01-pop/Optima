# Optima: Day 1 Notes

## Goal
Build a distributed key-value database in Go.

## Initial operations
- Put
- Get
- Delete

## Initial storage
An in-memory Go map.

## Current limitations
- Data disappears when the process exits.
- No network API.
- No replication.
- No concurrent access protection yet.

## Questions
- What is a storage engine?
- How does persistence work?
- How do multiple database nodes coordinate?