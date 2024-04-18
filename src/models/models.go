package models

type Configuration struct {
	GroupSize        int     `yaml:"group_size"`
	DeviationPercent float64 `yaml:"deviation_percent"`
	Latency          float64 `yaml:"sp_gr_latency"`
	Skill            float64 `yaml:"sp_gr_skill"`
}

type User struct {
	Name    string  `json:"name"`
	Skill   float64 `json:"skill"`
	Latency float64 `json:"latency"`

	CreateTime int64
	AvgWeight  float64
}

type Result struct {
	GroupNumber int64 `json:"group_number"`

	MinSkill float64 `json:"min_skill"`
	MaxSkill float64 `json:"max_skill"`
	AvgSkill float64 `json:"avg_skill"`

	MinLatency float64 `json:"min_latency"`
	MaxLatency float64 `json:"max_latency"`
	AvgLatency float64 `json:"avg_latency"`

	MinTime float64 `json:"min_time"`
	MaxTime float64 `json:"max_time"`
	AvgTime float64 `json:"avg_time"`

	UserList []string `json:"user_list"`
}
