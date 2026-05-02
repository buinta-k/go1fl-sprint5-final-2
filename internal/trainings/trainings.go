package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	str:=strings.Split(datastring,",")
	if len(str)!=3 {
		return fmt.Errorf("Длина слайса не равна 3")
	}
	t.Steps,err:=strconv.Atoi(str[0]) 
	if err!=nil{
		return fmt.Errorf("Ошибка преобразования string в int")
	}
	t.TrainingType=str[1]
	hours, err2:=time.ParseDuration(str[2])
	if err2!=nil {
		fmt.Errorf("Ошибка преобразования string в duration")
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance:=Distance(t.Steps, t.Height)
	speed:=MeanSpeed(t.Steps, t.Height, t.Duration)
	var calory int
	if t.TrainingType=="Бег" {
		calory=RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	}
	if t.TrainingType=="Ходьба" {
		calory=WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else {
		return "",fmt.Errorf("Неизвестный тип тренировки")
	}
	result:=fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.f2 ч.\nДистанция: %d км.\nСкорость: %.f2 км/ч\nСожгли калорий: %.f2\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calory,
	)
	return result, nil
}
