package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	section := strings.Split(data, ",")

	if len(section) != 3 {
		return 0, "", 0, fmt.Errorf("неверный формат данных")
	}

	steps, err := strconv.Atoi(section[0])
	if err != nil {
		return 0, "", 0, err
	}

	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(section[2])
	if err != nil {
		return 0, "", 0, err
	}

	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("продолжительность должна быть больше 0")
	}

	activity := section[1]

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm
	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	return distance(steps, height) / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, trainingType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)

	var calories float64

	switch trainingType {
	case "Бег":
		calories, err = RunningSpentCalories(
			steps,
			weight,
			height,
			duration,
		)
		if err != nil {
			return "", err
		}

	case "Ходьба":
		calories, err = WalkingSpentCalories(
			steps,
			weight,
			height,
			duration,
		)
		if err != nil {
			return "", err
		}

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		trainingType,
		duration.Hours(),
		dist,
		speed,
		calories,
	), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("недопустимое значение")
	}

	speed := meanSpeed(steps, height, duration)

	minutes := duration.Minutes()

	calories := (weight * speed * minutes) / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("недопустимое значение")
	}

	speed := meanSpeed(steps, height, duration)

	minutes := duration.Minutes()

	calories := (weight * speed * minutes) / minInH

	calories *= walkingCaloriesCoefficient

	return calories, nil
}
