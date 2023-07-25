package model

type Word struct {
	Eng string `json:"eng"`
	Chi string `json:"chi"`
}

type Detail struct {
	Name string `json:"name"`
	Init string `json:"init"`
	Len  int    `json:"len"`
}

func DeleteItem(wordList []Word, flag Word) []Word {
	var tmp []Word
	for _, i := range wordList {
		if flag != i {
			tmp = append(tmp, i)
		}
	}
	return tmp
}

// 定义Site结构体
type Site struct {
	Location string
}
