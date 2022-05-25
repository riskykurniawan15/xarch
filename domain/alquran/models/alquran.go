package models

type ReadChapter struct {
	Chapter Chapter `json:"chapter"`
	Verse   []Verse `json:"verse"`
}

type Chapter struct {
	ID                   int    `json:"id"`
	Bismillah            bool   `json:"bismillah"`
	ChapterNumber        int    `json:"chapter_number"`
	ChapterType          string `json:"chapter_type"`
	ChapterVerse         int    `json:"chapter_verse"`
	ChapterName          string `json:"chapter_name"`
	ChapterNameArabic    string `json:"chapter_name_arabic"`
	ChapterNameComplex   string `json:"chapter_name_complex"`
	ChapterNameTranslate struct {
		Indonesia string `json:"indonesia"`
		English   string `json:"english"`
	} `json:"chapter_name_translate"`
}

type Verse struct {
	ID             int    `json:"id"`
	VerseNumber    int    `json:"verse_number"`
	VerseQuran     string `json:"verse_quran"`
	VerseLatin     string `json:"verse_latin"`
	VerseTranslate struct {
		Indonesia string `json:"indonesia"`
		English   string `json:"english"`
	} `json:"verse_translate"`
}
