package controller

import "app/db"

type Controller struct {
	dbConnection *db.DBConnection
}

func NewController(dbConnection *db.DBConnection) *Controller {
	return &Controller{
		dbConnection: dbConnection,
	}
}
