package jsontranslator

import "maps"

type Language string

const (
	Automatic          Language = "Automatic"
	Afrikaans          Language = "Afrikaans"
	Albanian           Language = "Albanian"
	Amharic            Language = "Amharic"
	Arabic             Language = "Arabic"
	Armenian           Language = "Armenian"
	Azerbaijani        Language = "Azerbaijani"
	Basque             Language = "Basque"
	Belarusian         Language = "Belarusian"
	Bengali            Language = "Bengali"
	Bosnian            Language = "Bosnian"
	Bulgarian          Language = "Bulgarian"
	Catalan            Language = "Catalan"
	Cebuano            Language = "Cebuano"
	Chichewa           Language = "Chichewa"
	ChineseSimplified  Language = "Chinese (Simplified)"
	ChineseTraditional Language = "Chinese (Traditional)"
	Corsican           Language = "Corsican"
	Croatian           Language = "Croatian"
	Czech              Language = "Czech"
	Danish             Language = "Danish"
	Dutch              Language = "Dutch"
	English            Language = "English"
	Esperanto          Language = "Esperanto"
	Estonian           Language = "Estonian"
	Filipino           Language = "Filipino"
	Finnish            Language = "Finnish"
	French             Language = "French"
	Frisian            Language = "Frisian"
	Galician           Language = "Galician"
	Georgian           Language = "Georgian"
	German             Language = "German"
	Greek              Language = "Greek"
	Gujarati           Language = "Gujarati"
	HaitianCreole      Language = "Haitian Creole"
	Hausa              Language = "Hausa"
	Hawaiian           Language = "Hawaiian"
	Hebrew             Language = "Hebrew"
	Hindi              Language = "Hindi"
	Hmong              Language = "Hmong"
	Hungarian          Language = "Hungarian"
	Icelandic          Language = "Icelandic"
	Igbo               Language = "Igbo"
	Indonesian         Language = "Indonesian"
	Irish              Language = "Irish"
	Italian            Language = "Italian"
	Japanese           Language = "Japanese"
	Javanese           Language = "Javanese"
	Kannada            Language = "Kannada"
	Kazakh             Language = "Kazakh"
	Khmer              Language = "Khmer"
	Korean             Language = "Korean"
	KurdishKurmanji    Language = "Kurdish (Kurmanji)"
	Kyrgyz             Language = "Kyrgyz"
	Lao                Language = "Lao"
	Latin              Language = "Latin"
	Latvian            Language = "Latvian"
	Lithuanian         Language = "Lithuanian"
	Luxembourgish      Language = "Luxembourgish"
	Macedonian         Language = "Macedonian"
	Malagasy           Language = "Malagasy"
	Malay              Language = "Malay"
	Malayalam          Language = "Malayalam"
	Maltese            Language = "Maltese"
	Maori              Language = "Maori"
	Marathi            Language = "Marathi"
	Mongolian          Language = "Mongolian"
	MyanmarBurmese     Language = "Myanmar (Burmese)"
	Nepali             Language = "Nepali"
	Norwegian          Language = "Norwegian"
	Pashto             Language = "Pashto"
	Persian            Language = "Persian"
	Polish             Language = "Polish"
	Portuguese         Language = "Portuguese"
	Punjabi            Language = "Punjabi"
	Romanian           Language = "Romanian"
	Russian            Language = "Russian"
	Samoan             Language = "Samoan"
	Scots_Gaelic       Language = "Scots Gaelic"
	Serbian            Language = "Serbian"
	Sesotho            Language = "Sesotho"
	Shona              Language = "Shona"
	Sindhi             Language = "Sindhi"
	Sinhala            Language = "Sinhala"
	Slovak             Language = "Slovak"
	Slovenian          Language = "Slovenian"
	Somali             Language = "Somali"
	Spanish            Language = "Spanish"
	Sundanese          Language = "Sundanese"
	Swahili            Language = "Swahili"
	Swedish            Language = "Swedish"
	Tajik              Language = "Tajik"
	Tamil              Language = "Tamil"
	Telugu             Language = "Telugu"
	Thai               Language = "Thai"
	Turkish            Language = "Turkish"
	Ukrainian          Language = "Ukrainian"
	Urdu               Language = "Urdu"
	Uzbek              Language = "Uzbek"
	Vietnamese         Language = "Vietnamese"
	Welsh              Language = "Welsh"
	Xhosa              Language = "Xhosa"
	Yiddish            Language = "Yiddish"
	Yoruba             Language = "Yoruba"
	Zulu               Language = "Zulu"
)

var AILanguageMap = map[Language]string{
	Automatic:          "auto",
	Afrikaans:          "af-ZA",
	Albanian:           "sq-AL",
	Amharic:            "am-ET",
	Arabic:             "ar-SA",
	Armenian:           "hy-AM",
	Azerbaijani:        "az-AZ",
	Basque:             "eu-ES",
	Belarusian:         "be-BY",
	Bengali:            "bn-BD",
	Bosnian:            "bs-BA",
	Bulgarian:          "bg-BG",
	Catalan:            "ca-ES",
	Cebuano:            "ceb-PH",
	Chichewa:           "ny-MW",
	ChineseSimplified:  "zh-CN",
	ChineseTraditional: "zh-TW",
	Corsican:           "co-FR",
	Croatian:           "hr-HR",
	Czech:              "cs-CZ",
	Danish:             "da-DK",
	Dutch:              "nl-NL",
	English:            "en-US",
	Esperanto:          "eo",
	Estonian:           "et-EE",
	Filipino:           "tl-PH",
	Finnish:            "fi-FI",
	French:             "fr-FR",
	Frisian:            "fy-NL",
	Galician:           "gl-ES",
	Georgian:           "ka-GE",
	German:             "de-DE",
	Greek:              "el-GR",
	Gujarati:           "gu-IN",
	HaitianCreole:      "ht-HT",
	Hausa:              "ha-NG",
	Hawaiian:           "haw-US",
	Hebrew:             "he-IL",
	Hindi:              "hi-IN",
	Hmong:              "hmn",
	Hungarian:          "hu-HU",
	Icelandic:          "is-IS",
	Igbo:               "ig-NG",
	Indonesian:         "id-ID",
	Irish:              "ga-IE",
	Italian:            "it-IT",
	Japanese:           "ja-JP",
	Javanese:           "jv-ID",
	Kannada:            "kn-IN",
	Kazakh:             "kk-KZ",
	Khmer:              "km-KH",
	Korean:             "ko-KR",
	KurdishKurmanji:    "ku-TR",
	Kyrgyz:             "ky-KG",
	Lao:                "lo-LA",
	Latin:              "la",
	Latvian:            "lv-LV",
	Lithuanian:         "lt-LT",
	Luxembourgish:      "lb-LU",
	Macedonian:         "mk-MK",
	Malagasy:           "mg-MG",
	Malay:              "ms-MY",
	Malayalam:          "ml-IN",
	Maltese:            "mt-MT",
	Maori:              "mi-NZ",
	Marathi:            "mr-IN",
	Mongolian:          "mn-MN",
	MyanmarBurmese:     "my-MM",
	Nepali:             "ne-NP",
	Norwegian:          "no-NO",
	Pashto:             "ps-AF",
	Persian:            "fa-IR",
	Polish:             "pl-PL",
	Portuguese:         "pt-PT",
	Punjabi:            "pa-IN",
	Romanian:           "ro-RO",
	Russian:            "ru-RU",
	Samoan:             "sm-WS",
	Scots_Gaelic:       "gd-GB",
	Serbian:            "sr-RS",
	Sesotho:            "st-LS",
	Shona:              "sn-ZW",
	Sindhi:             "sd-PK",
	Sinhala:            "si-LK",
	Slovak:             "sk-SK",
	Slovenian:          "sl-SI",
	Somali:             "so-SO",
	Spanish:            "es-ES",
	Sundanese:          "su-ID",
	Swahili:            "sw-KE",
	Swedish:            "sv-SE",
	Tajik:              "tg-TJ",
	Tamil:              "ta-IN",
	Telugu:             "te-IN",
	Thai:               "th-TH",
	Turkish:            "tr-TR",
	Ukrainian:          "uk-UA",
	Urdu:               "ur-PK",
	Uzbek:              "uz-UZ",
	Vietnamese:         "vi-VN",
	Welsh:              "cy-GB",
	Xhosa:              "xh-ZA",
	Yiddish:            "yi",
	Yoruba:             "yo-NG",
	Zulu:               "zu-ZA",
}

func getLanguageKeyFromValue(from string, data map[Language]string) Language {
	keys := maps.Keys(data)
	for language := range keys {
		if data[language] == from {
			return language
		}
	}
	return ""
}
