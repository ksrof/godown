# Godown
Godown is a tiny-teeny utility that helps you convert your Markdown files into HTML ones.
## Roadmap for 1.0.0
- [ ] Watch for file changes within a specific directory
- [ ] Extract meta data from the Markdown file
- [ ] Support full markdown syntax
- [ ] Save output to a specific directory
## Examples
Right know the tool is just able to convert a level one heading:
```go
package main

import (
	convert "github.com/ksrof/godown"
)

func main() {
	// Read reads a .md file and converts its
	// contents into HTML.
	convert.Read("example.md")
}
```
Things aren't quite optimal yet so don't expect much of it for now.
## License
The MIT License (MIT) - see [`license`](https://github.com/ksrof/godown/blob/main/LICENSE) for more details.



