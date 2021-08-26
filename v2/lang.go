package deeplapi

import (
	"encoding/json"
	"strings"
)

type Lang int

const (
	LangNotSet     Lang = iota // NotSet
	LangBulgarian              // BG
	LangCzech                  // CS
	LangDanish                 // DA
	LangGerman                 // DE
	LangGreek                  // EL
	LangEnglish                // EN
	LangSpanish                // ES
	LangEstonian               // ET
	LangFinnish                // FI
	LangFrench                 // FR
	LangHungarian              // HU
	LangItalian                // IT
	LangJapanese               // JA
	LangLithuanian             // LT
	LangLatvian                // LV
	LangDutch                  // NL
	LangPolish                 // PL
	LangPortuguese             // PT
	LangRomanian               // RO
	LangRussian                // RU
	LangSlovak                 // SK
	LangSlovenian              // SL
	LangSwedish                // SV
	LangChinese                // ZH
)

func (l Lang) String() string {
	return [...]string{
		"NotSet", "BG", "CS", "DA", "DE",
		"EL", "EN", "ES", "ET", "FI",
		"FR", "HU", "IT", "JA", "LT",
		"LV", "NL", "PL", "PT", "RO",
		"RU", "SK", "SL", "SV", "ZH"}[l]
}

func ParseLang(s string) Lang {
	switch strings.ToLower(s) {
	default:
		return LangNotSet
	case "bg":
		return LangBulgarian
	case "cs":
		return LangCzech
	case "da":
		return LangDanish
	case "de":
		return LangGerman
	case "el":
		return LangGreek
	case "en":
		return LangEnglish
	case "es":
		return LangSpanish
	case "et":
		return LangEstonian
	case "FI":
		return LangFinnish
	case "FR":
		return LangFrench
	case "HU":
		return LangHungarian
	case "IT":
		return LangItalian
	case "JA":
		return LangJapanese
	case "LT":
		return LangLithuanian
	case "LV":
		return LangLatvian
	case "NL":
		return LangDutch
	case "PL":
		return LangPolish
	case "PT":
		return LangPortuguese
	case "RO":
		return LangRomanian
	case "RU":
		return LangRussian
	case "SK":
		return LangSlovak
	case "SL":
		return LangSlovenian
	case "SV":
		return LangSwedish
	case "ZH":
		return LangChinese
	}
}

func (l *Lang) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*l = ParseLang(s)
	return nil
}

func (l Lang) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}
