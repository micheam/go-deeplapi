package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/micheam/go-deeplapi/v2"
)

var (
	authKey    = flag.String("auth-key", os.Getenv("DEEPL_AUTH_KEY"), "Authentication Key for DeepL API")
	sourceLang = flag.String("source", "", "Language of the text to be translated. If it is omitted, it will be determined automatically.")
	targetLang = flag.String("target", deeplapi.LangEnglish.String(), "Language of the text to be translated")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	svc := deeplapi.NewTextTranslatingService(deeplapi.New(*authKey))
	source := deeplapi.ParseLang(*sourceLang)
	target := deeplapi.ParseLang(*targetLang)
	result, err := svc.TraslateSingleText(ctx, flag.Arg(0), source, target)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result.Text)
}