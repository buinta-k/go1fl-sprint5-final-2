package spentenergy

import (
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps<=0 || weight<=0 || height<=0 || duration<=0 {
		return 0, fmt.Errorf("Неккоректный параметр")
	}
	speed:=MeanSpeed(steps,height,duration)
	calory:=(weight*speed)/(duration.Minutes())
	return calory*walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps<=0 || weight<=0 || height<=0 || duration<=0 {
		return 0, fmt.Errorf("Неккоректный параметр")
	}
	speed:=MeanSpeed(steps,height,duration)
	durationInMinutes:=duration.Minutes()
	return =(weight*speed*durationInMinutes)/minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps<0 || duration<=0 {
		return 0
	}
	distance:=Distance(steps,height)
	return distance/duration.Hours()
}

func Distance(steps int, height float64) float64 {
	if steps<=0 || height<=0 {
		return 0
	}
	length:=height*stepLengthCoefficient
	result:=length*float64(steps)
	return result/mInKm
}
