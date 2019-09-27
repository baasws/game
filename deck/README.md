# Deck

[![Build Status](https://travis-ci.org/briscola-as-a-service/deck.svg?branch=master)](https://travis-ci.org/briscola-as-a-service/deck)
[![Coverage Status](https://coveralls.io/repos/github/briscola-as-a-service/deck/badge.svg?branch=master)](https://coveralls.io/github/briscola-as-a-service/deck?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/briscola-as-a-service/deck)](https://goreportcard.com/report/github.com/briscola-as-a-service/deck)
[![GoDoc](https://godoc.org/github.com/briscola-as-a-service/deck?status.svg)](https://godoc.org/github.com/briscola-as-a-service/deck)

> Please append `[ci skip]` to commit message, if test execution is not required

## Usage

### 1, 2, 3.. start

```go
// declare a Deck
var d Deck

// you must Shuffle a new deck, huh?
d.Shuffle()

// you Pick a card, like a pillow
card, err := d.Pick()
```

### More...

```go
// get card points
points := d.GetCardPoints(Card{
  Value: 3,
  Semen: "spade",
})
// points = 10
```

### Even more...

```go
// let's drop the `expendable` card, aka `2 di spade`
err := d.Drop()
// err = true, if we Pick-ed a card before. Drop should be done just after the
// Shuffle()
```
