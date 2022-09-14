package deeplapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const PathTextTranslating = "translate"

// See Deepl API Document for Details.
// https://www.deepl.com/docs-api/translating-text/

// TODO:
// - [ ] [Multiple sentences](https://www.deepl.com/docs-api/translating-text/multiple-sentences/)
// - [ ] [large-volumes](https://www.deepl.com/docs-api/translating-text/large-volumes/)
// - [ ] [In-Text Markup Support](https://www.deepl.com/docs-api/translating-text/in-text-markup/)

// TextTranslatingService ...
//
// Request Parameters:
//
//	*text*
//
//	(Required) Text to be translated. Only UTF8-encoded plain text is supported.
//	The parameter may be specified multiple times and translations are returned
//	in the same order as they are requested. Each of the parameter values may contain
//	multiple sentences. Up to 50 texts can be sent for translation in one request.
//
//
//	*source_lang*
//
//	(Optional) Language of the text to be translated.
//	If this parameter is omitted, the API will attempt to detect the language of the text and translate it.
//
//
//	*target_lang*
//
//	(Required) The language into which the text should be translated.
//
//
//	*split_sentences*
//
//	(Optional) Sets whether the translation engine should first split the input into sentences.
//	This is enabled by default. Possible values are:
//
//	    * "0"           - no splitting at all, whole input is treated as one sentence
//	    * "1" (default) - splits on interpunction and on newlines
//	    * "nonewlines"  - splits on interpunction only, ignoring newlines
//
//	For applications that send one sentence per text parameter, it is advisable to set split_sentences=0,
//	in order to prevent the engine from splitting the sentence unintentionally.
//
//
//	*preserve_formatting*
//
//	(Optional) Sets whether the translation engine should respect the original formatting,
//	even if it would usually correct some aspects. Possible values are:
//
//	    * "0" (default)
//	    * "1"
//
//	The formatting aspects affected by this setting include:
//
//	    * Punctuation at the beginning and end of the sentence
//	    * Upper/lower case at the beginning of the sentence
//
//
//	*formality*
//
//	(Optional) Sets whether the translated text should lean towards formal or informal language.
//	This feature currently only works for target languages
//	"DE" (German), "FR" (French), "IT" (Italian), "ES" (Spanish),
//	"NL" (Dutch) , "PL" (Polish), "PT-PT"       , "PT-BR" (Portuguese) and "RU" (Russian).
//
//	Possible options are:
//
//	    * "default" (default)
//	    * "more" - for a more formal language
//	    * "less" - for a more informal language
//
//
//	*glossary_id*
//
//	(Optional) Specify the glossary to use for the translation.  Important: This requires the source_lang parameter
//	to be set and the language pair of the glossary has to match the language pair of the request.
type TextTranslatingService struct {
	client *Client
}

func NewTextTranslatingService(c *Client) *TextTranslatingService {
	return &TextTranslatingService{c}
}

type traslateTextResponse struct {
	Result []traslateTextResult `json:"translations,omitempty"`
}

type traslateTextResult struct {
	Source Lang   `json:"detected_source_language,omitempty"`
	Text   string `json:"text,omitempty"`
}

type TraslateSingleTextResult struct {
	Target, Source Lang
	Text           string
}

func (t TextTranslatingService) TraslateSingleText(ctx context.Context, text string, source, target Lang) (*TraslateSingleTextResult, error) {
	log.Println("base url: " + BaseURL())
	_url := strings.Join([]string{BaseURL(), APIVersion, PathTextTranslating}, "/")
	param := url.Values{
		"text":        []string{text},
		"target_lang": []string{target.String()},
	}
	if source != LangNotSet {
		param.Add("source_lang", source.String())
	}

	// TODO(michea): Change to MethodPost with form
	// form := url.Values{
	// 	"text":        []string{text},
	// 	"target_lang": []string{target.String()},
	// }
	// if source != LangNotSet {
	// 	form.Add("source_lang", source.String())
	// }
	// body := bytes.NewReader([]byte(form.Encode()))

	b, err := t.client.doRequest(ctx, http.MethodGet, _url, nil, param)
	if err != nil {
		return nil, fmt.Errorf("do-request: %w", err)
	}
	resp := new(traslateTextResponse)
	err = json.Unmarshal(b, resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &TraslateSingleTextResult{
		Target: target,
		Source: resp.Result[0].Source,
		Text:   resp.Result[0].Text,
	}, nil
}
