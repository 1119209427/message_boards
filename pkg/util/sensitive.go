package util

import (
	"github.com/importcjj/sensitive"
)

var Filter *sensitive.Filter

const WordDictPath = "./document/sensitiveDict.txt"

func InitFilter() {
	Filter = sensitive.New()
	err := Filter.LoadWordDict(WordDictPath)
	if err != nil {
		LogrusObj.Infoln(err)
	}
}
