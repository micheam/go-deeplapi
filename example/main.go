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
