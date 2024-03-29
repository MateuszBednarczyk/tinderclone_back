// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	domain "tinderclone_back/src/pkg/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ICityStore is an autogenerated mock type for the ICityStore type
type ICityStore struct {
	mock.Mock
}

// IsCityAlreadyAvailable provides a mock function with given fields: cityName
func (_m *ICityStore) IsCityAlreadyAvailable(cityName string) bool {
	ret := _m.Called(cityName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(cityName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsCityInCountryAlreadyAvailable provides a mock function with given fields: cityName, countryID
func (_m *ICityStore) IsCityInCountryAlreadyAvailable(cityName string, countryID uuid.UUID) bool {
	ret := _m.Called(cityName, countryID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, uuid.UUID) bool); ok {
		r0 = rf(cityName, countryID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SaveNewCity provides a mock function with given fields: entity
func (_m *ICityStore) SaveNewCity(entity *domain.City) error {
	ret := _m.Called(entity)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.City) error); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectAllCitiesWhereCountryIdEqual provides a mock function with given fields: countryID
func (_m *ICityStore) SelectAllCitiesWhereCountryIdEqual(countryID uuid.UUID) ([]domain.City, error) {
	ret := _m.Called(countryID)

	var r0 []domain.City
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) ([]domain.City, error)); ok {
		return rf(countryID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) []domain.City); ok {
		r0 = rf(countryID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.City)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(countryID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectCityByName provides a mock function with given fields: cityName
func (_m *ICityStore) SelectCityByName(cityName string) *domain.City {
	ret := _m.Called(cityName)

	var r0 *domain.City
	if rf, ok := ret.Get(0).(func(string) *domain.City); ok {
		r0 = rf(cityName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.City)
		}
	}

	return r0
}

type mockConstructorTestingTNewICityStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewICityStore creates a new instance of ICityStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewICityStore(t mockConstructorTestingTNewICityStore) *ICityStore {
	mock := &ICityStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
