package services

import "tinderclone_back/src/pkg/stores"

type IPartnerSuggester interface {
}

type partnerSuggester struct {
	userStore stores.IUserStore
}

func NewPartnerSuggester(userStore stores.IUserStore) *partnerSuggester {
	return &partnerSuggester{
		userStore: userStore,
	}
}

