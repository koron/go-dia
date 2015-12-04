# go-dia

go-dia is emulation of Perl's `while<>` for golang.

Usage example:

```go
package main

import (
	"fmt"

	"github.com/koron/go-dia"
)

func main() {
	err := dia.For(func(line string) error {
		fmt.Printf("%s:%d:%s\n", dia.Filename, dia.LineNum, line)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
```

## Refered

*   [すごく長い行を読む場合の注意点](http://qiita.com/syohex/items/034fc6e377603ae4654a)
*   [Goでフィルタコマンドを怠惰に書く](http://qiita.com/Maki-Daisuke/items/79146b1c447a36ccf306)
