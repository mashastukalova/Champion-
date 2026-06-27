package daysteps

import (
	//spentcalories "command-line-argumentsC:\\Users\\masha\\Dev\\Champion\\internal\\spentcalories\\spentcalories.go"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	//"github.com/Yandex-Practicum/tracker/internal/spentcalories"
	spentcalories "github.com/Yandex-Practicum/tracker/internal/spentcalories"
	//"github.com/mashastukalova/Champion/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	section := strings.Split(data, ",")

	if len(section) != 2 {
		return 0, 0, fmt.Errorf("неверный формат данных")
	}

	steps, err := strconv.Atoi(section[0])
	if err != nil {
		return 0, 0, err
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(section[1])
	if err != nil {
		return 0, 0, err
	}
	if duration <= 0 {
		return 0, 0, fmt.Errorf("продолжительность должна быть больше 0")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	calories, err := spentcalories.WalkingSpentCalories(
		steps,
		weight,
		height,
		duration,
	)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps,
		distanceKm,
		calories,
	)
}
