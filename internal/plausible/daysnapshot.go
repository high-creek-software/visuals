package plausible

import "time"
import "encoding/json"
import "strings"

type DaySnapshot struct {
    Date PlausibletTime `json:"date"`
    Visitors int `json:"visitors"`
}

type PlausibletTime time.Time

func (t PlausibletTime) MarshalJSON() ([]byte, error) {
	gt := time.Time(t)
	f := gt.Format(time.DateOnly)
	return json.Marshal(f)
}


func (t *PlausibletTime) UnmarshalJSON(b []byte) error {
	input := strings.Trim(string(b), `"`)
	gt, err := time.Parse(time.DateOnly, input)

	if err != nil {
		return err
	}

	*t = PlausibletTime(gt)
	return nil
}