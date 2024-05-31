package utils

import (
	"errors"
	"exoplanetservice/models"
)

func ValidateExoplanet(exoplanet models.Exoplanet) error {
	if exoplanet.Name == "" || exoplanet.Description == "" {
		return errors.New("name and description are required")
	}
	if exoplanet.Distance < 10 || exoplanet.Distance > 1000 {
		return errors.New("distance must be between 10 and 1000 light years")
	}
	if exoplanet.Radius < 0.1 || exoplanet.Radius > 10 {
		return errors.New("radius must be between 0.1 and 10 earth-radius units")
	}
	if exoplanet.Type == "Terrestrial" {
		if exoplanet.Mass == nil || *exoplanet.Mass < 0.1 || *exoplanet.Mass > 10 {
			return errors.New("terrestrial planets must have mass between 0.1 and 10 earth-mass units")
		}
	}
	return nil
}
