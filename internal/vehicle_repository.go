package internal

import "errors"

var (
	// ErrAlreadyExists is an error that occurs when a vehicle already exists
	ErrAlreadyExists = errors.New("vehicle already exists")
	// ErrorNotFound is an error that occurs when a vehicle is not found
	ErrorNotFound = errors.New("There is no vehicle with this characteristics")
	// ErrInvalidQuery is an error that occurs when a query is invalid
	ErrInvalidQuery = errors.New("repository: invalid query")
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// GetbyID is a method that returns a vehicle by id
	GetbyID(id int) (vehicle Vehicle, err error)
	// Save is a method that saves a vehicle
	Save(vehicle *Vehicle) (err error)
	//find by color and year
	FindByColorAndYear(color string, year int) (v map[int]Vehicle, err error)
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
	FindQuery(query map[string]any) (vehicles map[int]Vehicle, err error)
	//FilterByWeight is a method that returns a map of vehicles by weight
	FilterByWeight(query map[string]any) (vehicles map[int]Vehicle, err error)
}
