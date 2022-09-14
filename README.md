# go-deeplapi

## Usage

### AS CLI TOOL

```sh
$ go install github.com/micheam/go-deeplapi/cmd/translate@latest
```

```sh
$ export DEEPL_AUTH_KEY="<YOUR-DEEPL-AUTH-KEY>"
$ export DEEPL_API_DOMEIN="api-free.deepl.com" # for free plan

$ translate こんにちわ、世界
Hello, World
```

Please run `-help` for detailed usage.

```
$ translate -help
Usage of translate:
  translate [flags] [text]
Flags:
  --api-domain string
  	Domain for DeepL API. (Default) Environment variable: 'DEEPL_API_DOMAIN'
  	Use 'api.deepl.com' if not set.
  --auth-key string
  	Authentication Key for DeepL API. (Default) Environment variable: 'DEEPL_AUTH_KEY'
  --list-lang
  	Display a list of LANG_CODE that can be specified.
  --source string
  	LANG_CODE of the text to be translated. If it is omitted, it will be determined automatically.
  --target string
  	LANG_CODE of the text to be translated. (default 'EN')
Examples:
  $ translate --soruce ja --taget en 'あんなこといいな。できたらいいな'
  $ echo 'あんなこといいな。できたらいいな' | translate # source lang will auto-detected
```

### AS GO MODULE

```sh
$ go get github.com/micheam/go-deeplapi/v2
```

```go
package main

import (
	"context"
	"fmt"

	"github.com/micheam/go-deeplapi/v2"
)

func main() {

	const authKey = "YOUR-DEEPL-AUTH-KEY"
	client := deeplapi.New(authKey)
	service := deeplapi.NewTextTranslatingService(client)

	var (
		text   = "こんにちわ、世界"
		source = deeplapi.LangJapanese
		target = deeplapi.LangEnglish
	)

	res, _ := service.TraslateSingleText(context.Background(), text, source, target)
	fmt.Println(res.Text)
	// Output: Hello, World
}
```

## Requirements

- go 1.19 or higher

## Installation

```sh
$ go install github.com/micheam/go-deeplapi/cmd/translate
```

## License

MIT

## Author

[Michito Maeda](https://github.com/micheam)
