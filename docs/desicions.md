# Decision 001: Initial storage

## Decision
Use a Go map for the first in-memory implementation.

## Reason
A map provides simple key-based insertion and retrieval,
letting us focus on database API behavior first.

## Trade-offs
The data is volatile, and the implementation does not
provide persistence, replication, or concurrency safety.

## Future direction
Replace or extend the in-memory storage with a persistent
storage engine in a later phase