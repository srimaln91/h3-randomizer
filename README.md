# Random H3 Cell Generator

## How to use

```go
package main 

import (
	"github.com/srimaln91/h3-randomizer"
)

func main(){
	r := randomizer.NewRandomizer(MinLongtitude, MaxLongtitude, MinLatitude, MaxLatitude)
    h3Cell, location, err := r.RandomH3(8)
    
	fmt.Println(h3Cell, location.getLatitude(), location.getLongtitude())
}
```
