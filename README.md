# Find the Duplicate in an Array of N+1 Integers

> Striver SDE Sheet 010 · Day 2 Arrays — treat values as next-pointers and find the cycle entry with Floyd's tortoise and hare. Go, stdlib only, tested.

## Problem

An array holds n+1 integers, each between 1 and n, so at least one value repeats. Find that repeated value **without modifying the array and in O(1) space** — which rules out sorting and hash sets.

## Approach

Because every value is a valid index, `i -> nums[i]` defines a linked structure starting at index 0. A duplicated value is a node with two incoming edges — i.e. the entry of a cycle. That turns the problem into linked-list cycle detection: Floyd's tortoise and hare meet inside the cycle; restarting one pointer at 0 and stepping both one at a time makes them meet exactly at the cycle entry — the duplicate. (Index 0 can't be part of the cycle itself since no value maps to it, which is what makes the start position safe.)

**Complexity:** O(n) time, O(1) space, array untouched.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
