package subscription

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetSubsID() string {
	var Letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	a := make([]rune, 5)
	for i := range a {
		a[i] = Letters[rand.Intn(len(Letters))]
	}
	return "You have subscribed!\nYour subscription ID is " + string(a) + "\n"

}
