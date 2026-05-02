package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse()
	ActionInfo()
}

func Info(dataset []string, dp DataParser) {
	for _,val:=range dataset {
		res,err:=dp.Parse(val)
		if err!=nil {
			log.Print("Ошибка париснга")
			continue
		}
		info,err:=dp.ActionInfo()
		if err!=nil {
			log.Print("Ошибка парсинга")
			continue
		}
	}
}
