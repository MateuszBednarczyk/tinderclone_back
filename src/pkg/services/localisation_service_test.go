package services

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNearbyTowns(t *testing.T) {
	//given
	town := "Poland, Warsaw"
	r := 50000

	//when
	result := NewLocaliser().GetNearbyTowns(r, town)
	fmt.Println(result)

	//then
	assert.True(t, len(result) > 0, len(result))
}
