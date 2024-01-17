package service

import (
	"app/internal"
	"fmt"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func ValidateVehicle(vehicle *internal.Vehicle) (err error) {
	// - validate required fields
	switch {
	case (*vehicle).Brand == "":
		return fmt.Errorf("%w: Brand", internal.ErrFieldRequired)
	case (*vehicle).Model == "":
		return fmt.Errorf("%w: Model", internal.ErrFieldRequired)
	case (*vehicle).Color == "":
		return fmt.Errorf("%w: color", internal.ErrFieldRequired)
	case (*vehicle).FabricationYear == 0:
		return fmt.Errorf("%w: Year", internal.ErrFieldRequired)
	case (*vehicle).Capacity == 0:
		return fmt.Errorf("%w: Passengers", internal.ErrFieldRequired)
	case (*vehicle).Transmission == "":
		return fmt.Errorf("%w: Transmission", internal.ErrFieldRequired)
	case (*vehicle).MaxSpeed == 0:
		return fmt.Errorf("%w: Max Speed", internal.ErrVelocityOutOfRange)
	default:
		return nil

	}

}

func (s *VehicleDefault) Save(vehicle *internal.Vehicle) (err error) {
	//validate business rules
	if err = ValidateVehicle(vehicle); err != nil {
		return
	}

	// save vehicle
	err = s.rp.Save(vehicle)
	if err != nil {
		switch err {
		case internal.ErrAlreadyExists:
			return fmt.Errorf("%w: %s", internal.ErrVehicleAlreadyExists, err.Error())
		}
		return
	}
	return
}

// GetByID is a method that returns a vehicle by id
func (s *VehicleDefault) GetByID(id int) (vehicle internal.Vehicle, err error) {
	vehicle, err = s.rp.GetbyID(id)
	if err != nil {
		switch err {
		//if the error is not found, return an error
		case internal.ErrorNotFound:
			err = fmt.Errorf("%w: %s", internal.ErrNotFound, err.Error())
		}
		return
	}
	return
}

// FindByColorAndYear is a method that returns a map of vehicles by color and year
func (s *VehicleDefault) FindByColorAndYear(color string, year int) (vehicle map[int]internal.Vehicle, err error) {
	vehicle, err = s.rp.FindByColorAndYear(color, year)
	if err != nil {
		switch err {
		//if the error is not found, return an error
		case internal.ErrorNotFound:
			err = fmt.Errorf("%w: %s", internal.ErrNotFound, err.Error())
		}
		return
	}
	return
}

func (s *VehicleDefault) FindByBrandAndYearRange(brand string, yearRange [2]int) (vehicle map[int]internal.Vehicle, err error) {
	vehicle, err = s.rp.FindByBrandAndYearRange(brand, yearRange)
	if err != nil {
		switch err {
		//if the error is not found, return an error
		case internal.ErrorNotFound:
			err = fmt.Errorf("%w: %s", internal.ErrNotFound, err.Error())
		}
		return
	}
	return
}

// VelocityAveragebyBrand is a method that returns the average velocity by brand
func (s *VehicleDefault) VelocityAveragebyBrand(brand string) (average float64, err error) {
	average, err = s.rp.VelocityAveragebyBrand(brand)
	if err != nil {
		switch err {
		//if the error is not found, return an error
		case internal.ErrorNotFound:
			err = fmt.Errorf("%w: %s", internal.ErrNotFound, err.Error())
		}
		return
	}
	return
}

// SaveMany is a method that saves many vehicles
func (s *VehicleDefault) SaveMany(vehicles []internal.Vehicle) (err error) {
	err = s.rp.SaveMany(vehicles)
	if err != nil {
		switch err {
		case internal.ErrAlreadyExists:
			return fmt.Errorf("%w: %s", internal.ErrVehicleAlreadyExists, err.Error())
		}
		return
	}
	return
}

// UpdateVehicle is a method that updates a vehicle
func (s *VehicleDefault) UpdateVehicle(vehicle *internal.Vehicle) (err error) {
	//validate business rules
	//velocity must be between 0 and 300
	if err = ValidateVehicle(vehicle); err != nil {
		return
	}
	// update vehicle
	err = s.rp.UpdateVehicle(vehicle)
	if err != nil {
		switch err {
		case internal.ErrorNotFound:
			return fmt.Errorf("%w: %s", internal.ErrNotFound, err.Error())
		}
		return
	}
	return
}

// FindByFuelType is a method that returns a map of vehicles by fuel type
func (s *VehicleDefault) FindByFuelType(fueltype string) (vehicle map[int]internal.Vehicle, err error) {
	vehicle, err = s.rp.FindByFuelType(fueltype)
	if err != nil {
		switch err {
		//if the error is not found, return an error
		case internal.ErrorNotFound:
			err = fmt.Errorf("%w:FuelType%s", internal.ErrNotFound, err.Error())
		}
		return
	}
	return
}

// Delete is a method that deletes a vehicle
func (s *VehicleDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	if err != nil {
		switch err {
		case internal.ErrorNotFound:
			return fmt.Errorf("%w: %s", internal.ErrNotFound, err.Error())
		}
		return
	}
	return
}

// Findbytransmission is a method that returns a map of vehicles by transmission
func (s *VehicleDefault) FindByTransmission(transmission string) (vehicle map[int]internal.Vehicle, err error) {
	vehicle, err = s.rp.FindByTransmission(transmission)
	if err != nil {
		switch err {
		//if the error is not found, return an error
		case internal.ErrorNotFound:
			err = fmt.Errorf("%w:Transmission%s", internal.ErrNotFound, err.Error())
		}
		return
	}
	return
}

// CapacityAverageByBrand is a method that returns the average capacity of a vehicle by brand
func (s *VehicleDefault) CapacityAveragebyBrand(brand string) (average float64, err error) {
	average, err = s.rp.CapacityAveragebyBrand(brand)
	if err != nil {
		switch err {
		//if the error is not found, return an error
		case internal.ErrorNotFound:
			err = fmt.Errorf("%w: %s", internal.ErrNotFound, err.Error())
		}
		return
	}
	return
}

// FindQuery is a method that returns a map of vehicles by query
func (s *VehicleDefault) FindQuery(query map[string]any) (vehicle map[int]internal.Vehicle, err error) {
	vehicle, err = s.rp.FindQuery(query)
	if err != nil {
		switch err {
		//if the error is not found, return an error
		case internal.ErrInvalidQuery:
			err = fmt.Errorf("%w: %s", internal.ErrNotFound, err.Error())
		}
		return
	}
	return
}

// FilterByWeight is a method that returns a map of vehicles by weight
func (s *VehicleDefault) FilterByWeight(query map[string]any) (vehicle map[int]internal.Vehicle, err error) {
	vehicle, err = s.rp.FilterByWeight(query)
	if err != nil {
		switch err {
		//if the error is not found, return an error
		case internal.ErrInvalidQuery:
			err = fmt.Errorf("%w: %s", internal.ErrNotFound, err.Error())
		}
		return
	}
	return
}
