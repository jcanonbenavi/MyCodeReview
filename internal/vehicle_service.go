package internal

import "errors"

var (
	ErrFieldRequired        = errors.New("field required")
	ErrVehicleAlreadyExists = errors.New("vehicle already exists")
	ErrNotFound             = errors.New("There is no vehicle with this color and year")
	ErrVehicleNotFound      = errors.New("Not found")
	ErrVehicleNotUpdated    = errors.New("Vehicle not updated")
	ErrVelocityOutOfRange   = errors.New("Velocity out of range")
)

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// GetbyID is a method that returns a vehicle by id
	GetByID(id int) (vehicle Vehicle, err error)
	// Save is a method that saves a vehicle
	Save(vehicle *Vehicle) (err error)
	// FindByColorAndYear is a method that returns a map of vehicles by color and year
	FindByColorAndYear(color string, year int) (vehicle map[int]Vehicle, err error)
	// FindByBrandAndYearRange is a method that returns a map of vehicles by brand and year range
	FindByBrandAndYearRange(brand string, yearRange [2]int) (vehicle map[int]Vehicle, err error)
	// VelocityAverageByBrand is a method that returns the average velocity of a vehicle by brand
	VelocityAveragebyBrand(brand string) (average float64, err error)
	// SaveMany is a method that saves many vehicles
	SaveMany(vehicles []Vehicle) (err error)
	// UpdateVehicle is a method that updates a vehicle
	UpdateVehicle(vehicle *Vehicle) (err error)
	//Find by type of FuelType
	FindByFuelType(fueltype string) (vehicle map[int]Vehicle, err error)
	//Delete is a method that deletes a vehicle by id
	Delete(id int) (err error)
	//Find by transmission type
	FindByTransmission(transmission string) (vehicle map[int]Vehicle, err error)
	//Find by capacity average by brand
	CapacityAveragebyBrand(brand string) (average float64, err error)
	// FindQuery is a method that returns a map of vehicles by query
	FindQuery(query map[string]any) (vehicle map[int]Vehicle, err error)
	//FilterByWeight is a method that returns a map of vehicles by weight
	FilterByWeight(query map[string]any) (vehicle map[int]Vehicle, err error)
}
