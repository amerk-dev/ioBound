package worker

import (
	"fmt"
	"ioBound/internal/domain"
	"math/rand"
	"time"
)

func TaskWorker(task *domain.Task) {
	randSecond := rand.Intn(60)
	fmt.Println(randSecond)
	time.Sleep(time.Duration(randSecond) * time.Second)
	task.Status = domain.StatusReady
}
