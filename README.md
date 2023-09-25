# IDK

This is just a collection of functions I use to reduce the amount of go code I have to write. Makes files smaller.

## Install

```
go get github.com/danhab99/idk@1.0
```

```go
// Recommended import line
import (
  _ "github.com/danhab99/idk"
)
```

## Functions

| name       | usage                                                    |
|------------|----------------------------------------------------------|
| Check      | Insert a function call that returns a value and an error |
| Check0     | Insert a function call that returns an error             |
| Accumulate | Collect the content of an array                          |
| Min        | Generic numeric min func                                 |
| Max        | Generic numeric max func                                 |

#### Ok but why tho

Look the whole thing with go error handling is that there isn't any. The only thing you can really do is panic and defer recover calls. Go encourages you to return errors to the stack to let the programmer decide what to do. But sometimes you just don't want to do anything with the error, there shouldn't be an error but if there is than panic. So why should we have to keep typing all that, or use snippits. Just wrap it in a function call and shorten the line. Trust me you stop noticing all the `Check`s after awhile.
