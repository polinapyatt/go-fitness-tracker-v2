package actioninfo

import (
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for i, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("Ошибка парсинга строки %d ('%s'): %v", i+1, data, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка формирования информации о тренировке %d: %v", i+1, err)
			continue
		}

		log.Println(info)
		log.Println("---")
	}
}
