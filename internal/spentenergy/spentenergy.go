package spentenergy

import (
	"errors"
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
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть положительной")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	if meanSpeed <= 0 {
		return 0, errors.New("не удалось рассчитать среднюю скорость")
	}

	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH
	calories *= walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть положительной")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	if meanSpeed <= 0 {
		return 0, errors.New("не удалось рассчитать среднюю скорость")
	}

	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	hours := duration.Hours()
	if hours <= 0 {
		return 0
	}

	return distance / hours

}

func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}

	stepLength := height * stepLengthCoefficient
	distance := float64(steps) * stepLength
	return distance / mInKm
}
