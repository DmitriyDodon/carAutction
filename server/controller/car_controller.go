package controller

import (
	"app/server/httpmodels"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// CreateCar godoc
// @Summary      Create car
// @Description  Create car for auction
// @Accept       json
// @Produce      json
// @Param        params body httpmodels.CarCreateRequest true "request body"
// @Success      201
// @Failure      400  {object}  httpmodels.CustomHttpError "[INCORRECT_REQUEST_BODY], [VALIDATION_FAILED]"
// @Failure      500  {object}  httpmodels.CustomHttpError "[ITERNAL_SERVER_ERROR]"
// @Router       /car [post]
func (con Controller) CreateCar(c echo.Context) error {
	req := new(httpmodels.CarCreateRequest)

	err := c.Bind(req)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, httpmodels.UnprocessableEntity)
	}

	err = req.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, err = con.dbConnection.Execute("insert into cars (id, color, price_in_cents, max_speed_mph, max_speed_kmp, vendor_name, model_name) values (?, ?, ?, ?, ?, ?, ?)",
		uuid.Must(uuid.NewRandom()).String(),
		req.Color,
		req.PriceInCents,
		req.MaxSpeedMPH,
		req.MaxSpeedKMP,
		req.VendorName,
		req.ModelName,
	)

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, httpmodels.ServerError)
	}

	return c.NoContent(http.StatusCreated)
}

// UpdateCar godoc
// @Summary      Update car
// @Description  Update car for auction
// @Accept       json
// @Produce      json
// @Param        carId   path    string  false  "Id of the car"
// @Param        params body httpmodels.CarCreateRequest true "request body"
// @Success      204
// @Failure      400  {object}  httpmodels.CustomHttpError "[INCORRECT_REQUEST_BODY], [VALIDATION_FAILED]"
// @Failure      500  {object}  httpmodels.CustomHttpError "[ITERNAL_SERVER_ERROR]"
// @Router       /car/{carId} [put]
func (con Controller) UpdateCar(c echo.Context) error {
	carID := c.Param("carID")

	req := new(httpmodels.CarCreateRequest)

	err := c.Bind(req)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, httpmodels.UnprocessableEntity)
	}

	err = req.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, err = con.dbConnection.Execute("update cars set color = ?, price_in_cents = ?, max_speed_mph = ?, max_speed_kmp = ?, vendor_name = ?, model_name = ? where id = ?",
		req.Color,
		req.PriceInCents,
		req.MaxSpeedMPH,
		req.MaxSpeedKMP,
		req.VendorName,
		req.ModelName,
		carID,
	)

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, httpmodels.ServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

// DeleteCar godoc
// @Summary      Delete car
// @Description  Delete car from auction
// @Param        carId   path    string  false  "Id of the car"
// @Success      204
// @Failure      500  {object}  httpmodels.CustomHttpError "[ITERNAL_SERVER_ERROR]"
// @Router       /car/{carId} [delete]
func (con Controller) DeleteCar(c echo.Context) error {
	carID := c.Param("carID")

	_, err := con.dbConnection.Execute("delete from cars where id = ?", carID)

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, httpmodels.ServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

// GetCar godoc
// @Summary      Get car
// @Description  Get car from auction
// @Param        carId   path    string  false  "Id of the car"
// @Produce      json
// @Success      200  {object}  httpmodels.CarResponse
// @Failure      404  {object}  httpmodels.CustomHttpError "[NOT_FOUND]"
// @Failure      500  {object}  httpmodels.CustomHttpError "[ITERNAL_SERVER_ERROR]"
// @Router       /car/{carId} [get]
func (con Controller) GetCar(c echo.Context) error {
	carID := c.Param("carID")
	carRow := con.dbConnection.QueryRow("select * from cars where id = ?", carID)

	car := new(httpmodels.CarResponse)

	err := carRow.Scan(&car.Id, &car.Color, &car.PriceInCents, &car.MaxSpeedMPH, &car.MaxSpeedKMP, &car.VendorName, &car.ModelName, &car.DateCreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, httpmodels.NotFoundError)
		}
		return c.JSON(http.StatusInternalServerError, httpmodels.ServerError)
	}

	return c.JSON(http.StatusOK, car)
}

// GetCars godoc
// @Summary      Get car list
// @Description  Get car list from auction
// @Produce      json
// @Success      200  {object}  []httpmodels.CarResponse
// @Failure      500  {object}  httpmodels.CustomHttpError "[ITERNAL_SERVER_ERROR]"
// @Router       /car [get]
func (con Controller) ListCars(c echo.Context) error {
	carRows, err := con.dbConnection.Query("select * from cars")

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, httpmodels.ServerError)
	}

	cars := make([]httpmodels.CarResponse, 0)

	for carRows.Next() {
		car := new(httpmodels.CarResponse)

		err := carRows.Scan(&car.Id, &car.Color, &car.PriceInCents, &car.MaxSpeedMPH, &car.MaxSpeedKMP, &car.VendorName, &car.ModelName, &car.DateCreatedAt)

		if err != nil {
			// если произошла ошибка при парсинге строки из бд, тогда пропускаем такую entity
			continue
		}

		cars = append(cars, *car)
	}

	return c.JSON(http.StatusOK, cars)
}
