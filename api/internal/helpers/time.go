package helpers

import "time"

func TimeNow() time.Time {
	moscow, _ := time.LoadLocation("Europe/Moscow")
	return time.Now().In(moscow)
}
