package matchmaker

import (
	"LTest/src/config"
	"LTest/src/models"
	"encoding/json"
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var (
	groupNumber int64                     = 0
	QueueChan   chan models.User          = make(chan models.User, 20)
	queueMap    map[float64][]models.User = make(map[float64][]models.User)
	muQueue                               = sync.Mutex{}
)

func MatchMake() {
	for {
		select {
		case newUser := <-QueueChan:
			var index float64 = 0
			minDeviation := newUser.AvgWeight * (100 - config.Config.DeviationPercent) / 100
			maxDeviation := newUser.AvgWeight * (100 + config.Config.DeviationPercent) / 100
			muQueue.Lock()
			for k := range queueMap {
				if k >= minDeviation && k <= maxDeviation {
					if index == 0 {
						index = k
					} else {
						if module(newUser.AvgWeight-k) < module(newUser.AvgWeight-index) {
							index = k
						}
					}
				}
			}

			if index == 0 {
				queueMap[newUser.AvgWeight] = append(queueMap[newUser.AvgWeight], newUser)
			} else {
				if len(queueMap[index])+1 == config.Config.GroupSize {
					goodGroup := append(queueMap[index], newUser)
					delete(queueMap, index)
					go getResultGroup(goodGroup)
				} else {
					avgWeightGroup := newUser.AvgWeight
					for _, v := range queueMap[index] {
						avgWeightGroup += v.AvgWeight
					}
					avgWeightGroup = avgWeightGroup / float64(len(queueMap[index])+1)
					queueMap[avgWeightGroup] = append(queueMap[index], newUser)
					if index != avgWeightGroup {
						delete(queueMap, index)
					}
				}
			}
			muQueue.Unlock()
		default:
		}
	}
}

func getResultGroup(goodGroup []models.User) {
	result := models.Result{}
	result.GroupNumber = incGroupNumber()
	now := time.Now().Unix()
	for _, v := range goodGroup {
		timeFrom := float64(now - v.CreateTime)
		result.UserList = append(result.UserList, v.Name)
		if len(result.UserList) == 1 {
			result.MinSkill = v.Skill
			result.MaxSkill = v.Skill
			result.AvgSkill = v.Skill

			result.MinLatency = v.Latency
			result.MaxLatency = v.Latency
			result.AvgLatency = v.Latency

			result.MinTime = timeFrom
			result.MaxTime = timeFrom
			result.AvgTime = timeFrom
		} else {
			result.MinSkill = math.Min(result.MinSkill, v.Skill)
			result.MaxSkill = math.Max(result.MaxSkill, v.Skill)
			result.AvgSkill += v.Skill

			result.MinLatency = math.Min(result.MinLatency, v.Latency)
			result.MaxLatency = math.Max(result.MaxLatency, v.Latency)
			result.AvgLatency += v.Latency

			result.MinTime = math.Min(result.MinTime, timeFrom)
			result.MaxTime = math.Max(result.MaxTime, timeFrom)
			result.AvgTime += timeFrom
		}
	}
	result.AvgSkill = result.AvgSkill / float64(len(goodGroup))
	result.AvgLatency = result.AvgLatency / float64(len(goodGroup))
	result.AvgTime = result.AvgTime / float64(len(goodGroup))

	resultJson, _ := json.Marshal(result)
	fmt.Println(string(resultJson))
}

func module(numb float64) float64 {
	if numb < 0 {
		return numb * -1
	}

	return numb
}

func incGroupNumber() int64 {
	return atomic.AddInt64(&groupNumber, 1)
}
