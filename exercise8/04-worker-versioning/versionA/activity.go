package versiona

import (
	"fmt"
	"math/rand"
)

func RecordDriverLocation(driverID string) (string, error) {
	lat := 1.3521 + (rand.Float64()-0.5)*0.05
	lng := 103.8198 + (rand.Float64()-0.5)*0.05
	return fmt.Sprintf("Driver %s at (%.4f, %.4f)", driverID, lat, lng), nil
}
