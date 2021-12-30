package AbstractFactory

import (
	"errors"
	"fmt"
)

const (
	SportMotorbikeType = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFactory struct {

}

func (m *MotorbikeFactory) GetVehicle(v int) (Vehicle,error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike),nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike),nil
	default:
		return nil,errors.New(fmt.Sprintf("Vehicle of Type %d not recognized\n",m))
	}
}