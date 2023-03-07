package services

import (
	"context"
	"fmt"
	"os"

	"googlemaps.github.io/maps"
)

type ILocaliser interface {
	GetNearbyTowns(r int, town string) []string
}

type localiser struct {
}

func NewLocaliser() *localiser {
	return &localiser{}
}

func (*localiser) GetNearbyTowns(r int, town string) []string {
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("API_KEY")))
	result := []string{}
	if err != nil {
		return result
	}

	placesRequest := &maps.NearbySearchRequest{
		Location: &maps.LatLng{},
		Radius:   uint(r),
		Type:     maps.PlaceType("cities"),
	}

	geocodeRequest := &maps.GeocodingRequest{
		Address: town,
	}
	geocodeResult, err := c.Geocode(context.Background(), geocodeRequest)
	if err != nil {
		panic(err)
	}

	placesRequest.Location = &geocodeResult[0].Geometry.Location
	placesResult, err := c.NearbySearch(context.Background(), placesRequest)
	if err != nil {
		panic(err)
	}

	for _, place := range placesResult.Results {
		fmt.Println(place.Name)
		result = append(result, place.Name)
	}
	return result
}
