package adapter

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/rgoncalvesrr/fullcycle-labs-cloudrun/application"
	"github.com/rgoncalvesrr/fullcycle-labs-cloudrun/configs"
	"net/http"
	"regexp"
)

type coordinateServiceAdapter struct {
	cfg *configs.Config
}

type Output struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

func NewCoordinateServiceAdapter(config *configs.Config) application.ICoordinateService {
	return &coordinateServiceAdapter{cfg: config}
}

func (c *coordinateServiceAdapter) GetByCep(ctx context.Context, cep string) (*application.Coordinate, error) {
	url := fmt.Sprintf("%s/{cep}", c.cfg.CepApiUrl)

	match, _ := regexp.MatchString("[0-9]{8}", cep)

	if !match {
		return nil, application.ErrCepMalformed
	}

	client := resty.New()
	r, e := client.R().
		SetContext(ctx).
		SetHeader("Accept", "application/json").
		SetPathParam("cep", cep).
		SetResult(&Output{}).
		Get(url)

	if e != nil {
		return nil, e
	}

	switch r.StatusCode() {
	case http.StatusNotFound:
		return nil, application.ErrCepNotFound
	case http.StatusBadRequest:
		return nil, application.ErrCepInvalid
	}

	cord := application.NewCoordinate(r.Result().(*Output).Lat, r.Result().(*Output).Lng)
	return cord, nil
}
