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
	Duration time.duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	data:=strings.Split(datastring,",")
	if len(data)!=2 {
		fmt.Errorf("Длина слайса не равна 2")
	}
	steps,err2:=strconv.Atoi(data[0])
	if err2!=nil {
		fmt.Errorf("Ошибка преобразования string в int")
	}
	hours,err3:=time.ParseDuration(data[1])
	if err3!=nil {
		fmt.Errorf("Ошибка преобразования string в duration")
	}

	ds.Steps=steps
	ds.Duration=hours
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance:=Distance(ds.Steps,ds.Height)
	calory, err:=WalkingSpentCalories(ds.Steps,ds.Weight,ds.Height,ds.Duration)
	
	if err!=nil {
		return 0, fmt.Errorf("Ошибка вычисления функции")
	}
	
	result:=fmt.Sprintf(
		"Количество шагов %d.\nДистанция составила %.f2 км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calory
	)
	return result, nil
}
