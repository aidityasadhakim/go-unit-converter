package services

import "fmt"

type UnitConverterService struct{}

func NewUnitConverterService() *UnitConverterService {
	return &UnitConverterService{}
}

func (s *UnitConverterService) ConvertLength(value float64, from string, to string) (float64, error) {
	meters, err := toMeters(value, from)
	if err != nil {
		return 0, err
	}
	result, err := fromMeters(meters, to)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *UnitConverterService) ConvertTemperature(value float64, from string, to string) (float64, error) {
	// First convert to Celsius as the base unit
	celsius, err := toCelsius(value, from)
	if err != nil {
		return 0, err
	}
	// Then convert from Celsius to the target unit
	result, err := fromCelsius(celsius, to)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *UnitConverterService) ConvertWeight(value float64, from string, to string) (float64, error) {
	// First convert to grams as the base unit
	grams, err := toGrams(value, from)
	if err != nil {
		return 0, err
	}
	// Then convert from grams to the target unit
	result, err := fromGrams(grams, to)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func toMeters(value float64, from string) (float64, error) {
	switch from {
	case "mm":
		return value / 1000, nil
	case "cm":
		return value / 100, nil
	case "m":
		return value, nil
	case "km":
		return value * 1000, nil
	case "in":
		return value * 0.0254, nil
	case "ft":
		return value * 0.3048, nil
	case "yd":
		return value * 0.9144, nil
	case "mi":
		return value * 1609.34, nil
	default:
		return 0, nil
	}
}

func fromMeters(value float64, to string) (float64, error) {
	switch to {
	case "mm":
		return value * 1000, nil
	case "cm":
		return value * 100, nil
	case "m":
		return value, nil
	case "km":
		return value / 1000, nil
	case "in":
		return value / 0.0254, nil
	case "ft":
		return value / 0.3048, nil
	case "yd":
		return value / 0.9144, nil
	case "mi":
		return value / 1609.34, nil
	default:
		return 0, fmt.Errorf("unsupported length unit: %s", to)
	}
}

// Helper functions for temperature conversion
func toCelsius(value float64, from string) (float64, error) {
	switch from {
	case "c":
		return value, nil
	case "f":
		return (value - 32) * 5 / 9, nil
	case "k":
		return value - 273.15, nil
	default:
		return 0, fmt.Errorf("unsupported temperature unit: %s", from)
	}
}

func fromCelsius(value float64, to string) (float64, error) {
	switch to {
	case "c":
		return value, nil
	case "f":
		return (value * 9 / 5) + 32, nil
	case "k":
		return value + 273.15, nil
	default:
		return 0, fmt.Errorf("unsupported temperature unit: %s", to)
	}
}

// Helper functions for weight conversion
func toGrams(value float64, from string) (float64, error) {
	switch from {
	case "g":
		return value, nil
	case "kg":
		return value * 1000, nil
	case "mg":
		return value / 1000, nil
	case "oz":
		return value * 28.3495, nil
	case "lb":
		return value * 453.592, nil
	default:
		return 0, fmt.Errorf("unsupported weight unit: %s", from)
	}
}

func fromGrams(value float64, to string) (float64, error) {
	switch to {
	case "g":
		return value, nil
	case "kg":
		return value / 1000, nil
	case "mg":
		return value * 1000, nil
	case "oz":
		return value / 28.3495, nil
	case "lb":
		return value / 453.592, nil
	default:
		return 0, fmt.Errorf("unsupported weight unit: %s", to)
	}
}