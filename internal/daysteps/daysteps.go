package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return fmt.Errorf("Длина слайса не равна 2")
	}

	steps, err2 := strconv.Atoi(data[0])
	if err2 != nil {
		return fmt.Errorf("Ошибка преобразования string в int")
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше нуля")
	}

	hours, err3 := time.ParseDuration(data[1])
	if err3 != nil {
		return fmt.Errorf("Ошибка преобразования string в duration")
	}
	if hours <= 0 {
		return fmt.Errorf("длительность должна быть больше нуля")
	}

	ds.Steps = steps
	ds.Duration = hours
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance:=spentenergy.Distance(ds.Steps,ds.Height)
	calory, err:=spentenergy.WalkingSpentCalories(ds.Steps,ds.Weight,ds.Height,ds.Duration)
	
	if err!=nil {
		return "", fmt.Errorf("Ошибка вычисления функции")
	}
	
	result:=fmt.Sprintf(
		"Количество шагов %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calory
	)
	return result, nil
}
