package repository

import (
	"app/internal"
	"fmt"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle, lastId int) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{
		db:     defaultDb,
		lastId: lastId,
	}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
	//I'm add a lastId to save the last id used in the db
	lastId int
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// GetbyID is a method that returns a vehicle by id
func (r *VehicleMap) GetbyID(id int) (vehicle internal.Vehicle, err error) {

	vehicle, ok := r.db[id]
	if !ok {
		err = internal.ErrorNotFound
		return
	}
	return
}

// Save is a method that saves a vehicle
func (r *VehicleMap) Save(vehicule *internal.Vehicle) (err error) {
	for _, vehicles := range (*r).db {
		if vehicles.Model == vehicule.Model && vehicles.Brand == vehicule.Brand && vehicles.FabricationYear == vehicule.FabricationYear {
			return internal.ErrAlreadyExists
		}
	}
	//I'm incrementing the lastId and then I'm assing it to the vehicle id
	(*r).lastId++
	//assing the lastId to the vehicle id
	vehicule.Id = (*r).lastId
	//id is the key and the value is the vehicle
	(*r).db[vehicule.Id] = *vehicule
	return
}

func (r *VehicleMap) FindByColorAndYear(color string, year int) (vehicle map[int]internal.Vehicle, err error) {
	vehicle = make(map[int]internal.Vehicle)
	//I'm iterating over the db map and I'm comparing the color and the year with the parameters
	for key, value := range (*r).db {
		if value.FabricationYear == year && value.Color == color {
			//if the color and the year are the same, I'm assing the vehicle to the map
			vehicle[key] = value
		}
	}
	//if the vehicle is empty, return an error
	if len(vehicle) == 0 {
		err = internal.ErrorNotFound
	}
	return
}

func (r *VehicleMap) FindByBrandAndYearRange(brand string, yearRange [2]int) (vehicle map[int]internal.Vehicle, err error) {
	vehicle = make(map[int]internal.Vehicle)
	//I'm iterating over the db map and I'm comparing the brand and the year range with the parameters
	for key, value := range (*r).db {
		if value.FabricationYear >= yearRange[0] && value.FabricationYear <= yearRange[1] && value.Brand == brand {
			//if the brand and the year range are the same, I'm assing the vehicle to the map
			vehicle[key] = value
		}
	}
	//if the vehicle is empty, return an error
	if len(vehicle) == 0 {
		err = internal.ErrorNotFound
	}
	return
}

// VelocityAverageByBrand is a method that returns the average velocity of a vehicle by brand
func (r *VehicleMap) VelocityAveragebyBrand(brand string) (average float64, err error) {
	//I'm iterating over the db map and I'm comparing the brand with the parameter
	vehicle := make(map[int]internal.Vehicle)
	velocity := 0.0
	for key, value := range (*r).db {
		if value.Brand == brand {
			vehicle[key] = value
			//I'm calculating the average velocity
			velocity += value.MaxSpeed
			average = velocity / float64(len(vehicle))
		}
	}
	//if the vehicle is empty, return an error
	if len(vehicle) == 0 {
		err = internal.ErrorNotFound
	}
	return
}

// SaveMany is a method that saves many vehicles
func (r *VehicleMap) SaveMany(vehicles []internal.Vehicle) (err error) {
	//I'm iterating over the vehicles slice and I'm saving each vehicle
	for _, vehicle := range vehicles {
		//I'm calling the Save method to save each vehicle
		err = r.Save(&vehicle)
		if err != nil {
			return err
		}
	}
	return
}

// UpdateVehicle is a method that updates a vehicle
func (r *VehicleMap) UpdateVehicle(vehicle *internal.Vehicle) (err error) {
	//I'm checking if the vehicle exists in the db
	_, ok := (*r).db[vehicle.Id]
	if !ok {
		return internal.ErrorNotFound
	}
	//I'm updating the vehicle
	(*r).db[vehicle.Id] = *vehicle
	return
}

// FindByFuelType is a method that returns a map of vehicles by fuel type
func (r *VehicleMap) FindByFuelType(fuelType string) (vehicle map[int]internal.Vehicle, err error) {
	vehicle = make(map[int]internal.Vehicle)
	//I'm iterating over the db map and I'm comparing the fuel type with the parameter
	for key, value := range (*r).db {
		if value.FuelType == fuelType {
			//if the fuel type is the same, I'm assing the vehicle to the map
			vehicle[key] = value
		}
	}
	//if the vehicle is empty, return an error
	if len(vehicle) == 0 {
		err = internal.ErrorNotFound
	}
	return
}

// DeleteVehicle is a method that deletes a vehicle
func (r *VehicleMap) Delete(id int) (err error) {
	//I'm checking if the vehicle exists in the db
	_, ok := (*r).db[id]
	if !ok {
		return internal.ErrorNotFound
	}
	//I'm deleting the vehicle
	delete((*r).db, id)
	return
}

// FindByTransmission is a method that returns a map of vehicles by transmission
func (r *VehicleMap) FindByTransmission(transmission string) (vehicle map[int]internal.Vehicle, err error) {
	vehicle = make(map[int]internal.Vehicle)
	//I'm iterating over the db map and I'm comparing the transmission with the parameter
	for key, value := range (*r).db {
		if value.Transmission == transmission {
			//if the transmission is the same, I'm assing the vehicle to the map
			vehicle[key] = value
		}
	}
	//if the vehicle is empty, return an error
	if len(vehicle) == 0 {
		err = internal.ErrorNotFound
	}
	return
}

// CapacityAverageByBrand is a method that returns the average capacity of a vehicle by brand
func (r *VehicleMap) CapacityAveragebyBrand(brand string) (average float64, err error) {
	//I'm iterating over the db map and I'm comparing the brand with the parameter
	vehicle := make(map[int]internal.Vehicle)
	capacity := 0.0
	for key, value := range (*r).db {
		if value.Brand == brand {
			vehicle[key] = value
			//I'm calculating the average velocity
			capacity += float64(value.Capacity)
			average = capacity / float64(len(vehicle))
		}
	}
	//if the vehicle is empty, return an error
	if len(vehicle) == 0 {
		err = internal.ErrorNotFound
	}
	return
}

// FindQuery is a method that returns a map of vehicles by query
func (r *VehicleMap) FindQuery(query map[string]any) (vehicles map[int]internal.Vehicle, err error) {
	// create map of vehicles
	vehicles = make(map[int]internal.Vehicle)

	// validate if query is set
	querySet := len(query) > 0
	// iterate over db map
	for id, item := range r.db {
		// if query is set, filter
		if querySet {
			//Search for the minLength in the query map
			minLength, ok := query["min_length"]
			if !ok {
				err = internal.ErrInvalidQuery
				return
			}
			//convert minLength to float64
			minLengthFloat, ok := minLength.(float64)
			if !ok {
				err = internal.ErrInvalidQuery
				return
			}

			//maxlength
			maxLength, ok := query["max_length"]
			if !ok {
				err = internal.ErrInvalidQuery
				return
			}
			maxLengthFloat, ok := maxLength.(float64)
			if !ok {
				err = internal.ErrInvalidQuery
				return
			}

			//minwidth
			minWidth, ok := query["min_width"]
			if !ok {
				err = internal.ErrInvalidQuery
				return
			}
			minWidthFloat, ok := minWidth.(float64)
			if !ok {
				err = internal.ErrInvalidQuery
				return
			}
			//maxwidth
			maxWidth, ok := query["max_width"]
			if !ok {
				err = internal.ErrInvalidQuery
				return
			}
			maxWidthFloat, ok := maxWidth.(float64)
			if !ok {
				err = internal.ErrInvalidQuery
				return
			}
			// filter
			if item.Height < minLengthFloat || item.Height > maxLengthFloat || item.Width < minWidthFloat || item.Width > maxWidthFloat {
				continue
			}
		}

		// by default add item
		vehicles[id] = item
		fmt.Println(vehicles)
	}
	return
}

// FilterByWeight is a method that returns a map of vehicles by weight
func (r *VehicleMap) FilterByWeight(query map[string]any) (vehicles map[int]internal.Vehicle, err error) {
	// create map of vehiclesx
	vehicles = make(map[int]internal.Vehicle)

	// filter by query
	querySet := len(query) > 0
	for id, item := range r.db {
		// if query is set, filter
		if querySet {
			//search for the weight_min in the query map
			weightMin, ok := query["weight_min"]
			if !ok {
				err = internal.ErrInvalidQuery
				return
			}
			weightMinFloat, ok := weightMin.(float64)
			if !ok {
				err = internal.ErrInvalidQuery
				return
			}

			//maxlength
			weightMax, ok := query["weight_max"]
			if !ok {
				err = internal.ErrInvalidQuery
				return
			}
			weightMaxFloat, ok := weightMax.(float64)
			if !ok {
				err = internal.ErrInvalidQuery
				return
			}
			// filter
			if item.Weight < weightMinFloat || item.Weight > weightMaxFloat {
				continue
			}
		}

		// by default add item
		vehicles[id] = item
	}
	return
}
