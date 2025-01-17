package jpushclient

import (
	"encoding/json"
	"time"
)

type Schedule struct {
	Cid     string                 `json:"cid"`
	Name    string                 `json:"name"`
	Enabled bool                   `json:"enabled"`
	Trigger map[string]interface{} `json:"trigger"`
	Push    *PayLoad               `json:"push"`
}

func NewSchedule(name, cid string, enabled bool, push *PayLoad) *Schedule {
	return &Schedule{
		Cid:     cid,
		Name:    name,
		Enabled: enabled,
		Push:    push,
	}
}
func (schedule *Schedule) SingleTrigger(t time.Time) {
	schedule.Trigger = map[string]interface{}{
		"single": map[string]interface{}{
			"time": t.Format("2006-01-02 15:04:05"),
		},
	}

}

func (schedule *Schedule) PeriodicalTrigger(start time.Time, end time.Time, time time.Time, timeUnit string, frequency int, point []string) {
	schedule.Trigger = map[string]interface{}{
		"periodical": map[string]interface{}{
			"start":     start.Format("2006-01-02 15:04:05"),
			"end":       end.Format("2006-01-02 15:04:05"),
			"time":      start.Format("15:04:05"),
			"time_unit": timeUnit,
			"frequency": frequency,
			"point":     point,
		},
	}
}
func (schedule *Schedule) ToBytes() ([]byte, error) {
	content, err := json.Marshal(schedule)
	if err != nil {
		return nil, err
	}
	return content, nil
}
