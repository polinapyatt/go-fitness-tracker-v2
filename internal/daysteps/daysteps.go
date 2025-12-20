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
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("нeверный формат строки: ожидается 2 части, получено %d", len(parts))
	}

	// Шаги
	stepsStr := parts[0]
	trimmedSteps := strings.TrimSpace(stepsStr)
	if trimmedSteps == "" {
		return fmt.Errorf("ошибка парсинга шагов: пустое значение")
	}
	// Проверка на пробелы в начале или конце
	if stepsStr != trimmedSteps {
		return fmt.Errorf("ошибка парсинга шагов: недопустимые пробелы")
	}

	steps, err := strconv.Atoi(trimmedSteps)
	if err != nil {
		return fmt.Errorf("ошибка парсинга шагов: %v", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}
	ds.Steps = steps

	// Продолжительность
	durationStr := parts[1]
	trimmedDuration := strings.TrimSpace(durationStr)
	if trimmedDuration == "" {
		return fmt.Errorf("ошибка парсинки продолжительности: пустое значение")
	}
	// Проверка на пробелы в начале или конце
	if durationStr != trimmedDuration {
		return fmt.Errorf("ошибка парсинки продолжительности: недопустимые пробелы")
	}

	duration, err := time.ParseDuration(trimmedDuration)
	if err != nil {
		return fmt.Errorf("ошибка парсинки продолжительности: %v", err)
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Количество шагов: %d.\n", ds.Steps)
	result += fmt.Sprintf("Дистанция составила %.2f км.\n", distance)
	result += fmt.Sprintf("Вы сожгли %.2f ккал.\n", calories)

	return result, nil
}
