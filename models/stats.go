package models

type Stats struct {
	Eliminations   int `json:"eliminations"`
	ObjectiveKills int `json:"objective_kills"`
	ObjectiveTime  int `json:"objective_time"`
	Damage         int `json:"damage"`
	Healing        int `json:"healing"`
	Deaths         int `json:"deaths"`
}
