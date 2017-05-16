package hn

type TNoteArray []TNote

func GetSampleNoteArray() TNoteArray {
	return TNoteArray{
		TNote{Title: "AAA", Content: "aaa aaa 111"},
		TNote{Title: "BBB", Content: "aaa aaa bbb"},
		TNote{Title: "CCC", Content: "aaa aaa ccc"},
	}
}