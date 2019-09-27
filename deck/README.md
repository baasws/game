# Deck

## Usage

### 1, 2, 3.. start

```go
// declare a Deck
d := deck.New()

// you Pick a card, like a pillow
card, err := d.Pick()
```

### More...

```go
// get card points
// TODO:
// points = 10
```

### Even more...

```go
// let's drop the `expendable` card, aka `2 di spade`
err := d.Drop()
// err = true, if we Pick-ed a card before. Drop should be done just after the
// Shuffle()
```
