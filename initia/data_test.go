package initia

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestAutoTime(t *testing.T) {
	for i := 0; i < 100; i++ {
		tt := fmt.Sprintf("19%d-%02d-%02d", 60+rand.Intn(40), 1+rand.Intn(12), 1+rand.Intn(28))
		ttt,_ := time.ParseInLocation("2006-01-02", tt, time.Local)
		fmt.Println(tt, ttt.Unix())
	}

}
