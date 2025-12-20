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
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("неверный формат строки: ожидается 3 части, получено %d", len(parts))
	}
	stepsStr := parts[0]
	trimmedSteps := strings.TrimSpace(stepsStr)
	if trimmedSteps == "" {
		return fmt.Errorf("ошибка парсинга количества шагов: пустое значение")
	}
	if stepsStr != trimmedSteps {
		return fmt.Errorf("ошибка парсинга количества шагов: недопустимые пробелы")
	}
	steps, err := strconv.Atoi(trimmedSteps)
	if err != nil {
		return fmt.Errorf("ошибка парсинга количества шагов: %v", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}
	t.Steps = steps

	trainingTypeStr := parts[1]
	t.TrainingType = strings.TrimSpace(trainingTypeStr)
	durationStr := parts[2]
	trimmedDuration := strings.TrimSpace(durationStr)
	if trimmedDuration == "" {
		return fmt.Errorf("ошибка парсинга продолжительности: пустое значение")
	}
	if durationStr != trimmedDuration {
		return fmt.Errorf("ошибка парсинга продолжительности: недопустимые пробелы")
	}
	duration, err := time.ParseDuration(trimmedDuration)
	if err != nil {
		return fmt.Errorf("ошибка парсинга продолжительности: %v", err)
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег", "бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба", "ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}

	if err != nil {
		return "", err
	}
	result := fmt.Sprintf("Тип тренировки: %s\n", t.TrainingType)
	result += fmt.Sprintf("Длительность: %.2f ч.\n", t.Duration.Hours())
	result += fmt.Sprintf("Дистанция: %.2f км.\n", distance)
	result += fmt.Sprintf("Скорость: %.2f км/ч\n", meanSpeed)
	result += fmt.Sprintf("Сожгли калорий: %.2f\n", calories)

	return result, nil
}
