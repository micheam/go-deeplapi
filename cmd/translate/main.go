package main

import (
	"bufio"
	"bytes"
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

	var err error
	ctx := context.Background()
	svc := deeplapi.NewTextTranslatingService(deeplapi.New(*authKey))
	source := deeplapi.ParseLang(*sourceLang)
	target := deeplapi.ParseLang(*targetLang)

	text := flag.Arg(0)
	if len(text) == 0 {
		text, err = scanStdin()
		if err != nil {
			log.Fatal(err)
		}
	}

	result, err := svc.TraslateSingleText(ctx, text, source, target)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result.Text)
}

func scanStdin() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	buf := new(bytes.Buffer)
	for scanner.Scan() {
		buf.Write(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return buf.String(), nil
}
