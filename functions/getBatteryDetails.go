package functions

import (
	"fmt"

	"github.com/distatus/battery"
)

type BatteryDetails struct {
	Percentage      float64 `json:"percentage"`
	State           string  `json:"state"`
	Remaining       float64 `json:"remaining"`
	DesignCapacity  float64 `json:"design_capacity"`
	CurrentCapacity float64 `json:"current_capacity"`
}

func GetBatteryDetails() (*BatteryDetails, error) {
	batteries, err := battery.GetAll()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve battery info: %w", err)
	}

	if len(batteries) == 0 {
		return nil, fmt.Errorf("no battery detected")
	}

	b := batteries[0]
	state := b.State.String()

	percentage := 0.0
	if b.Full > 0 {
		percentage = (b.Current / b.Full) * 100
	}

	remaining := 0.0
	if state == "discharging" && b.ChargeRate > 0 {
		remaining = b.Current / b.ChargeRate
	}

	return &BatteryDetails{
		Percentage:      percentage,
		State:           state,
		Remaining:       remaining,
		DesignCapacity:  b.Design,
		CurrentCapacity: b.Current,
	}, nil
}
