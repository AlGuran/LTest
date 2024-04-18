package routes

import (
	"LTest/src/config"
	"LTest/src/matchmaker"
	"LTest/src/models"
	"encoding/json"
	"github.com/mgutz/logxi/v1"
	"github.com/valyala/fasthttp"
	"time"
)

var (
	logger = log.New("rout_user")
)

func HandleUsers(ctx *fasthttp.RequestCtx) {
	result := models.User{}
	err := json.Unmarshal(ctx.Request.Body(), &result)
	if err != nil {
		logger.Error(err.Error())
		ctx.Error(err.Error(), 500)
		return
	}
	result.CreateTime = time.Now().Unix()
	result.AvgWeight = (result.Latency*config.Config.Latency + result.Skill*config.Config.Skill) /
		(config.Config.Latency + config.Config.Skill)
	logger.Info("Add user", result.Name)
	matchmaker.QueueChan <- result
}
