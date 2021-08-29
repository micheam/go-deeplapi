package main

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/micheam/go-deeplapi/v2"
)

var (
	authKey    = flag.String("auth-key", os.Getenv("DEEPL_AUTH_KEY"), "Authentication Key for DeepL API.")
	sourceLang = flag.String("source", "", "LANG_CODE of the text to be translated. If it is omitted, it will be determined automatically.")
	targetLang = flag.String("target", deeplapi.LangEnglish.String(), "LANG_CODE of the text to be translated.")

	langList = flag.Bool("list-lang", false, "Display a list of LANG_CODE that can be specified.")
)

func Usage() {
	// TODO(micheam): Replace with text/template
	fmt.Fprintf(os.Stderr, "Usage of translate:\n")
	fmt.Fprintf(os.Stderr, "  translate [flags] [text]\n")

	fmt.Fprintf(os.Stderr, "Flags:\n")
	fmt.Fprintf(os.Stderr, "  --auth-key string\n")
	fmt.Fprintf(os.Stderr, "  \tAuthentication Key for DeepL API. (Default) Environment variable: 'DEEPL_AUTH_KEY'\n")
	fmt.Fprintf(os.Stderr, "  --list-lang\n")
	fmt.Fprintf(os.Stderr, "  \tDisplay a list of LANG_CODE that can be specified.\n")
	fmt.Fprintf(os.Stderr, "  --source string\n")
	fmt.Fprintf(os.Stderr, "  \tLANG_CODE of the text to be translated. If it is omitted, it will be determined automatically.\n")
	fmt.Fprintf(os.Stderr, "  --target string\n")
	fmt.Fprintf(os.Stderr, "  \tLANG_CODE of the text to be translated. (default 'EN')\n")

	fmt.Fprintf(os.Stderr, "Examples:\n")
	fmt.Fprintf(os.Stderr, "  $ translate --soruce ja --taget en 'あんなこといいな。できたらいいな'\n")
	fmt.Fprintf(os.Stderr, "  $ echo 'あんなこといいな。できたらいいな' | translate # source lang will auto-detected\n")
	// flag.PrintDefaults()
}

func main() {
	log.SetPrefix("translate: ")
	flag.Usage = Usage
	flag.Parse()

	var err error
	ctx := context.Background()

	if *langList {
		printLangList()
		os.Exit(0)
	}

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

func printLangList() {
	for l := range deeplapi.LangList() {
		if l == deeplapi.LangNotSet {
			continue
		}
		fmt.Printf("%s\t%s\n", strings.ToLower(l.String()), l.DispName())
	}
}
