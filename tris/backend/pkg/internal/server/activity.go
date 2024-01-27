package server

import (
	"fmt"
	"time"
)

var ErrIDNotFound = fmt.Errorf("ID not found")

type Activity struct {
	Time   time.Time `json:"time"`
	Layout []string  `json:"layout"`
	ID     uint64    `json:"id"`
	Turn   string    `json:"turn"`
}

type Activities struct {
	activities []Activity
}

func (c *Activities) Insert(activity Activity) uint64 {
	activity.ID = uint64(len(c.activities))
	c.activities = append(c.activities, activity)
	return activity.ID
}

func (c *Activities) Retrieve(id uint64) (Activity, error) {
	if id >= uint64(len(c.activities)) {
		return Activity{}, ErrIDNotFound
	}
	return c.activities[id], nil
}
