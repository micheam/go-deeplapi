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

func (l Lang) DispName() string {
	return [...]string{
		"NotSet", "Bulgarian", "Czech", "Danish", "German",
		"Greek", "English", "Spanish", "Estonian", "Finnish",
		"French", "Hungarian", "Italian", "Japanese", "Lithuanian",
		"Latvian", "Dutch", "Polish", "Portuguese", "Romanian",
		"Russian", "Slovak", "Slovenian", "Swedish", "Chinese",
	}[l]
}

var allLang = []Lang{
	LangNotSet,
	LangBulgarian,
	LangCzech,
	LangDanish,
	LangGerman,
	LangGreek,
	LangEnglish,
	LangSpanish,
	LangEstonian,
	LangFinnish,
	LangFrench,
	LangHungarian,
	LangItalian,
	LangJapanese,
	LangLithuanian,
	LangLatvian,
	LangDutch,
	LangPolish,
	LangPortuguese,
	LangRomanian,
	LangRussian,
	LangSlovak,
	LangSlovenian,
	LangSwedish,
	LangChinese,
}

func LangList() <-chan Lang {
	c := make(chan Lang)
	go func() {
		for _, l := range allLang {
			c <- l
		}
		close(c)
	}()
	return c
}

func ParseLang(s string) Lang {
	switch strings.ToUpper(s) {
	default:
		return LangNotSet
	case "BG":
		return LangBulgarian
	case "CS":
		return LangCzech
	case "DA":
		return LangDanish
	case "DE":
		return LangGerman
	case "EL":
		return LangGreek
	case "EN":
		return LangEnglish
	case "ES":
		return LangSpanish
	case "ET":
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
