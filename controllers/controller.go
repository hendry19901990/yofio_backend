package controllers

import (
	"errors"
	"net/http"

	"github.com/hendry19901990/yofio_backend/repository"
	"github.com/jinzhu/gorm"
)

type Controller struct {
	DBType        string
	URLConnection string
	connection    *gorm.DB
}

func (cont *Controller) Init() error {
	conn, err := repository.GetConnection(cont.DBType, cont.URLConnection)
	if err != nil {
		return err
	}

	cont.connection = conn
	return nil
}

func (cont *Controller) GetConnection() (*gorm.DB, error) {
	if d := cont.connection.DB(); d != nil {
		if err := d.Ping(); err != nil {
			d.Close()
			if err := cont.Init(); err != nil {
				return nil, err
			}
		}
		return cont.connection, nil
	}

	return nil, errors.New("Error parsing sql.DB")
}

func (cont *Controller) WriteResponse(w http.ResponseWriter, content []byte, code int) {
	if code == 200 {
		w.Write(content)
	} else {
		http.Error(w, string(content), code)
	}
}
