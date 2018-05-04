# lavaloon
## Examples

```
% cat ./examples/02.lv
(import "fmt")

(defun main ()
  (fmt.Println "hello world"))
% lavaloon ./examples/02.lv
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```
