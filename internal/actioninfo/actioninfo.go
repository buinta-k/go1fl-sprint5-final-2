package actioninfo

import (
	"fmt"
	"log"
)
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string,error)
}

func Info(dataset []string, dp DataParser) {
	for _,val:=range dataset {
		err:=dp.Parse(val)
		if err!=nil {
			log.Print("Ошибка париснга")
			continue
		}
		info,err:=dp.ActionInfo()
		if err!=nil {
			log.Print("Ошибка парсинга")
			continue
		}
		fmt.Print(info)
	}
}
