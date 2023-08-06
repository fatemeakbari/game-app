package scheduler

import (
	"fmt"
	"gameapp/model"
	matchingservice "gameapp/service/matching"
	"github.com/go-co-op/gocron"
	"time"
)

type Scheduler struct {
	service matchingservice.Service
}

func New(service matchingservice.Service) Scheduler {

	return Scheduler{
		service: service,
	}
}

func (s *Scheduler) Start(done <-chan bool) {

	sch := gocron.NewScheduler(time.UTC).CronWithSeconds("*/30 * * * * *")

	fmt.Println("start")
	_, err := sch.Do(s.matchPlayer) // every second

	if err != nil {
		fmt.Println("err", err)
	}
	sch.StartBlocking()
	//for {
	//
	//	select {
	//	case <-done:
	//		fmt.Println("return from scheduler")
	//	default:
	//		time.Sleep(1 * time.Second)
	//		fmt.Println("scheduler running time:", time.Now().UnixMilli())
	//	}
	//
	//}
}

func (s *Scheduler) matchPlayer() {

	s.service.MatchWaitingPlayer(model.FootballCategory)
}
