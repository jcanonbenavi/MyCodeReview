package handler

import (
	"app/internal"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

type VehicleRequestJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get vehicle id from url
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid vehicle id",
			})
			return
		}
		// process
		vehicle, err := h.sv.GetByID(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrNotFound):
				response.Text(w, http.StatusNotFound, "ID not found")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}
		// response
		data := VehicleRequestJSON{
			ID:              vehicle.Id,
			Brand:           vehicle.Brand,
			Model:           vehicle.Model,
			Registration:    vehicle.Registration,
			Color:           vehicle.Color,
			FabricationYear: vehicle.FabricationYear,
			Capacity:        vehicle.Capacity,
			MaxSpeed:        vehicle.MaxSpeed,
			FuelType:        vehicle.FuelType,
			Transmission:    vehicle.Transmission,
			Weight:          vehicle.Weight,
			Height:          vehicle.Height,
			Length:          vehicle.Length,
			Width:           vehicle.Width,
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Save is a method that returns a handler for the route POST /vehicles

func (h *VehicleDefault) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - decode request body
		var req VehicleJSON
		// decode request body and convert it to VehicleJSON
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}

		// create vehicle - serialize request body to vehicle
		dimensions := internal.Dimensions{
			Height: req.Height,
			Length: req.Length,
			Width:  req.Width,
		}

		attributes := internal.VehicleAttributes{
			Brand:           req.Brand,
			Model:           req.Model,
			Registration:    req.Registration,
			Color:           req.Color,
			FabricationYear: req.FabricationYear,
			Capacity:        req.Capacity,
			MaxSpeed:        req.MaxSpeed,
			FuelType:        req.FuelType,
			Transmission:    req.Transmission,
			Weight:          req.Weight,
			Dimensions:      dimensions,
		}
		vehicle := internal.Vehicle{
			VehicleAttributes: attributes,
		}

		// process
		// - save vehicle
		if err := h.sv.Save(&vehicle); err != nil {
			switch {
			// if vehicle already exists, return 409 Conflict
			case errors.Is(err, internal.ErrVehicleAlreadyExists):
				response.Text(w, http.StatusConflict, "vehicle already exists")
			// if vehicle is invalid, return 400 Bad Request
			case errors.Is(err, internal.ErrFieldRequired):
				response.Text(w, http.StatusBadRequest, "Invalid body")
			// if vehicle is invalid, return 400 Bad Request
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}
		// response
		data := VehicleRequestJSON{
			ID:              vehicle.Id,
			Brand:           vehicle.Brand,
			Model:           vehicle.Model,
			Registration:    vehicle.Registration,
			Color:           vehicle.Color,
			FabricationYear: vehicle.FabricationYear,
			Capacity:        vehicle.Capacity,
			MaxSpeed:        vehicle.MaxSpeed,
			FuelType:        vehicle.FuelType,
			Transmission:    vehicle.Transmission,
			Weight:          vehicle.Weight,
			Height:          vehicle.Height,
			Length:          vehicle.Length,
			Width:           vehicle.Width,
		}
		// return 201 Created
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// FindByColorAndYear is a method that returns a handler for the route GET /vehicles/color/{color}/year/{year}
func (h *VehicleDefault) FindByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		color := chi.URLParam(r, "color")
		// convert year to int
		year, _ := strconv.Atoi(chi.URLParam(r, "year"))
		// process
		vehicles, err := h.sv.FindByColorAndYear(color, year)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrNotFound):
				response.Text(w, http.StatusNotFound, "Vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}
		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Success",
			"data":    vehicles,
		})

	}
}

// FindByBrandAndYearRange is a method that returns a handler for the route GET /vehicles/brand/{brand}/year/{start_year}/{end_year}
func (h *VehicleDefault) FindByBrandAndYearRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		brand := chi.URLParam(r, "brand")
		// convert year to int
		start, _ := strconv.Atoi(chi.URLParam(r, "start_year"))
		end, _ := strconv.Atoi(chi.URLParam(r, "end_year"))
		yearRange := [2]int{start, end}
		// process
		vehicles, err := h.sv.FindByBrandAndYearRange(brand, yearRange)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrNotFound):
				response.Text(w, http.StatusNotFound, "Vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Success",
			"data":    vehicles,
		})
	}
}

// VelocityAveragebyBrand is a method that returns a handler for the route GET /vehicles/velocity-average/:brand
func (h *VehicleDefault) VelocityAveragebyBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		brand := chi.URLParam(r, "brand")
		// convert year to int
		vehicles, err := h.sv.VelocityAveragebyBrand(brand)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrNotFound):
				response.Text(w, http.StatusNotFound, "Vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}
		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"Velocity average by " + brand + " branch": vehicles,
		})
	}
}

// SaveMany is a method that returns a handler for the route POST /vehicles/many
func (h *VehicleDefault) SaveMany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - decode request body
		var req []VehicleJSON
		// decode request body and convert it to VehicleJSON
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}
		//for each vehicle in req, create a vehicle and save it
		for _, vehicle := range req {
			// create vehicle - serialize request body to vehicle
			dimensions := internal.Dimensions{
				Height: vehicle.Height,
				Length: vehicle.Length,
				Width:  vehicle.Width,
			}

			attributes := internal.VehicleAttributes{
				Brand:           vehicle.Brand,
				Model:           vehicle.Model,
				Registration:    vehicle.Registration,
				Color:           vehicle.Color,
				FabricationYear: vehicle.FabricationYear,
				Capacity:        vehicle.Capacity,
				MaxSpeed:        vehicle.MaxSpeed,
				FuelType:        vehicle.FuelType,
				Transmission:    vehicle.Transmission,
				Weight:          vehicle.Weight,
				Dimensions:      dimensions,
			}
			vehicle := internal.Vehicle{
				VehicleAttributes: attributes,
			}
			// process
			// - save vehicle
			if err := h.sv.Save(&vehicle); err != nil {
				switch {
				// if vehicle already exists, return 409 Conflict
				case errors.Is(err, internal.ErrVehicleAlreadyExists):
					response.Text(w, http.StatusConflict, "vehicle already exists")
				// if vehicle is invalid, return 400 Bad Request
				case errors.Is(err, internal.ErrFieldRequired):
					response.Text(w, http.StatusBadRequest, "Invalid body")
				// if vehicle is invalid, return 400 Bad Request
				default:
					response.Text(w, http.StatusInternalServerError, "internal server error")
				}
				return
			}
			// // response
			data := VehicleRequestJSON{
				ID:              vehicle.Id,
				Brand:           vehicle.Brand,
				Model:           vehicle.Model,
				Registration:    vehicle.Registration,
				Color:           vehicle.Color,
				FabricationYear: vehicle.FabricationYear,
				Capacity:        vehicle.Capacity,
				MaxSpeed:        vehicle.MaxSpeed,
				FuelType:        vehicle.FuelType,
				Transmission:    vehicle.Transmission,
				Weight:          vehicle.Weight,
				Height:          vehicle.Height,
				Length:          vehicle.Length,
				Width:           vehicle.Width,
			}
			// return 201 Created
			response.JSON(w, http.StatusCreated, map[string]any{
				"message": "success",
				"data":    data,
			})
		}
	}
}

// UpdateMaxSpeed is a method that returns a handler for the route PUT /vehicles/:id/max-speed
func (h *VehicleDefault) UpdateMaxSpeed() http.HandlerFunc {
	// request
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "Invalid id")
			return
		}
		//check if vehicle exists
		vehicle, err := h.sv.GetByID(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrNotFound):
				response.Text(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}
		// create vehicle - serialize request body to vehicle
		data := VehicleRequestJSON{
			ID:              vehicle.Id,
			Brand:           vehicle.Brand,
			Model:           vehicle.Model,
			Registration:    vehicle.Registration,
			Color:           vehicle.Color,
			FabricationYear: vehicle.FabricationYear,
			Capacity:        vehicle.Capacity,
			MaxSpeed:        vehicle.MaxSpeed,
			FuelType:        vehicle.FuelType,
			Transmission:    vehicle.Transmission,
			Weight:          vehicle.Weight,
			Height:          vehicle.Height,
			Length:          vehicle.Length,
			Width:           vehicle.Width,
		}

		// decode request body and convert it to VehicleJSON
		if err := request.JSON(r, &data); err != nil {
			response.Text(w, http.StatusBadRequest, "Invalid request body")
			return
		}
		// create vehicle - serialize request body to vehicle
		dimensions := internal.Dimensions{
			Height: data.Height,
			Length: data.Length,
			Width:  data.Width,
		}
		attributes := internal.VehicleAttributes{
			Brand:           data.Brand,
			Model:           data.Model,
			Registration:    data.Registration,
			Color:           data.Color,
			FabricationYear: data.FabricationYear,
			Capacity:        data.Capacity,
			MaxSpeed:        data.MaxSpeed,
			FuelType:        data.FuelType,
			Transmission:    data.Transmission,
			Weight:          data.Weight,
			Dimensions:      dimensions,
		}
		vehicleserialize := internal.Vehicle{
			Id:                id,
			VehicleAttributes: attributes,
		}
		// process
		if err := h.sv.UpdateVehicle(&vehicleserialize); err != nil {
			switch {
			case errors.Is(err, internal.ErrNotFound):
				response.Text(w, http.StatusNotFound, "Vehicle not found")
			case errors.Is(err, internal.ErrFieldRequired):
				response.Text(w, http.StatusBadRequest, "Invalid body")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}

		// response
		// - deserialize vehicle to VehicleJSON
		dataResponse := VehicleJSON{
			ID:              vehicle.Id,
			Brand:           vehicle.Brand,
			Model:           vehicle.Model,
			Registration:    vehicle.Registration,
			Color:           vehicle.Color,
			FabricationYear: vehicle.FabricationYear,
			Capacity:        vehicle.Capacity,
			MaxSpeed:        vehicle.MaxSpeed,
			FuelType:        vehicle.FuelType,
			Transmission:    vehicle.Transmission,
			Weight:          vehicle.Weight,
			Height:          vehicle.Height,
			Length:          vehicle.Length,
			Width:           vehicle.Width,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Updated successfully",
			"data":    dataResponse,
		})
	}
}

// FindByFuelType is a method that returns a handler for the route GET /vehicles/fuel-type/:type
func (h *VehicleDefault) FindByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		//find by fuel type
		brand := chi.URLParam(r, "type")
		// process
		vehicles, err := h.sv.FindByFuelType(brand)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrNotFound):
				response.Text(w, http.StatusNotFound, "Not found vehicles with this fuel type")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Find By Fuel Type successfully",
			"data":    vehicles,
		})
	}
}

// Delete is a method that returns a handler for the route DELETE /vehicles/:id
func (h *VehicleDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid id",
				"data":    nil,
			})
			return
		}
		// process
		if err := h.sv.Delete(id); err != nil {
			switch {
			case errors.Is(err, internal.ErrNotFound):
				response.Text(w, http.StatusNotFound, "Vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}
		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Deleted successfully",
		})
	}
}

// FindBy
func (h *VehicleDefault) FindByTransmission() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		//find by fuel type
		transmission := chi.URLParam(r, "type")
		// process
		vehicles, err := h.sv.FindByTransmission(transmission)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrNotFound):
				response.Text(w, http.StatusNotFound, "Not found vehicles with this transmission")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Find By Transmission successfully",
			"data":    vehicles,
		})
	}
}

// CapacityAveragebyBrand is a method that returns a handler for the route GET /vehicles/capacity-average/:brand
func (h *VehicleDefault) CapacityAveragebyBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// request
		//find by fuel type
		brand := chi.URLParam(r, "brand")
		// process
		vehicles, err := h.sv.CapacityAveragebyBrand(brand)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrNotFound):
				response.Text(w, http.StatusNotFound, "Not found vehicles with this brand")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Capacity Average by Brand successfully",
			"data":    vehicles,
		})
	}
}

// UpdateFuelType is a method that returns a handler for the route PUT /vehicles/:id/fuel-type
func (h *VehicleDefault) UpdateFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "Invalid id")
			return
		}
		//check if vehicle exists
		vehicle, err := h.sv.GetByID(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrNotFound):
				response.Text(w, http.StatusNotFound, "Vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}
		// create vehicle - serialize request body to vehicle
		data := VehicleRequestJSON{
			ID:              vehicle.Id,
			Brand:           vehicle.Brand,
			Model:           vehicle.Model,
			Registration:    vehicle.Registration,
			Color:           vehicle.Color,
			FabricationYear: vehicle.FabricationYear,
			Capacity:        vehicle.Capacity,
			MaxSpeed:        vehicle.MaxSpeed,
			FuelType:        vehicle.FuelType,
			Transmission:    vehicle.Transmission,
			Weight:          vehicle.Weight,
			Height:          vehicle.Height,
			Length:          vehicle.Length,
			Width:           vehicle.Width,
		}

		// decode request body and convert it to VehicleJSON
		if err := request.JSON(r, &data); err != nil {
			response.Text(w, http.StatusBadRequest, "Invalid request body")
			return
		}
		// create vehicle - serialize request body to vehicle
		dimensions := internal.Dimensions{
			Height: data.Height,
			Length: data.Length,
			Width:  data.Width,
		}
		attributes := internal.VehicleAttributes{
			Brand:           data.Brand,
			Model:           data.Model,
			Registration:    data.Registration,
			Color:           data.Color,
			FabricationYear: data.FabricationYear,
			Capacity:        data.Capacity,
			MaxSpeed:        data.MaxSpeed,
			FuelType:        data.FuelType,
			Transmission:    data.Transmission,
			Weight:          data.Weight,
			Dimensions:      dimensions,
		}
		vehicleserialize := internal.Vehicle{
			Id:                id,
			VehicleAttributes: attributes,
		}
		// process
		if err := h.sv.UpdateVehicle(&vehicleserialize); err != nil {
			switch {
			case errors.Is(err, internal.ErrNotFound):
				response.Text(w, http.StatusNotFound, "Vehicle not found")
			case errors.Is(err, internal.ErrFieldRequired):
				response.Text(w, http.StatusBadRequest, "Invalid body")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}

		// response
		// - deserialize vehicle to VehicleJSON
		dataResponse := VehicleJSON{
			ID:              vehicle.Id,
			Brand:           vehicle.Brand,
			Model:           vehicle.Model,
			Registration:    vehicle.Registration,
			Color:           vehicle.Color,
			FabricationYear: vehicle.FabricationYear,
			Capacity:        vehicle.Capacity,
			MaxSpeed:        vehicle.MaxSpeed,
			FuelType:        vehicle.FuelType,
			Transmission:    vehicle.Transmission,
			Weight:          vehicle.Weight,
			Height:          vehicle.Height,
			Length:          vehicle.Length,
			Width:           vehicle.Width,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Updated successfully",
			"data":    dataResponse,
		})
	}
}

// FindAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		query := make(map[string]any)
		//request query params -min_length-
		minLength, ok := r.URL.Query()["min_length"]
		if ok {
			//convert string to float64
			minLengthFloat, err := strconv.ParseFloat(minLength[0], 64)
			if err != nil {
				//if error return bad request
				response.Text(w, http.StatusBadRequest, "Invalid min length")
				return
			}
			//if no error add to query map
			query["min_length"] = minLengthFloat
		}
		//request query params -max_length-
		maxLength, ok := r.URL.Query()["max_length"]
		if ok {
			//convert string to float64
			maxLengthFloat, err := strconv.ParseFloat(maxLength[0], 64)
			if err != nil {
				response.Text(w, http.StatusBadRequest, "Invalid max_length")
				return
			}
			//if no error add to query map
			query["max_length"] = maxLengthFloat
		}
		//request query params -min_width-
		minWidth, ok := r.URL.Query()["min_width"]
		if ok {
			//convert string to float64
			minWidthFloat, err := strconv.ParseFloat(minWidth[0], 64)
			if err != nil {
				response.Text(w, http.StatusBadRequest, "Invalid min_width")
				return
			}
			//if no error add to query map
			query["min_width"] = minWidthFloat
		}
		//request query params -max_width-
		maxWidth, ok := r.URL.Query()["max_width"]
		if ok {
			//convert string to float64
			maxWidthFloat, err := strconv.ParseFloat(maxWidth[0], 64)
			if err != nil {
				response.Text(w, http.StatusBadRequest, "Invalid max_width")
				return
			}
			//if no error add to query map
			query["max_width"] = maxWidthFloat
		}
		// process
		//send query to service and get response
		items, err := h.sv.FindQuery(query)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrNotFound):
				response.Text(w, http.StatusNotFound, "Not found vehicles")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}
		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Find by query successfully",
			"data":    items,
		})
	}
}

// FilterByWeight is a method that returns a handler for the route GET /vehicles/filter/weight
func (h *VehicleDefault) FilterByWeight() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		query := make(map[string]any)
		//request query params -weight_min-
		weightMin, ok := r.URL.Query()["weight_min"]
		if ok {
			//convert string to float64
			weightMinFloat, err := strconv.ParseFloat(weightMin[0], 64)
			if err != nil {
				response.Text(w, http.StatusBadRequest, "Invalid min length")
				return
			}
			//if no error add to query map
			query["weight_min"] = weightMinFloat
		}
		//request query params -weight_max-
		weightMax, ok := r.URL.Query()["weight_max"]
		if ok {
			//convert string to float64
			weightMaxFloat, err := strconv.ParseFloat(weightMax[0], 64)
			if err != nil {
				response.Text(w, http.StatusBadRequest, "Invalid max_length")
				return
			}
			//if no error add to query map
			query["weight_max"] = weightMaxFloat
		}
		// process
		items, err := h.sv.FilterByWeight(query)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrNotFound):
				response.Text(w, http.StatusNotFound, "Not found vehicles")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Find by query successfully",
			"data":    items,
		})
	}
}
