package dto

import (
	"otto/test/features/brands"
)

type BrandRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}

type BrandResponse struct {
	ID   int    `json:"id" binding:"require"`
	Name string `json:"name" binding:"require"`
}

func FromCore(userCore brands.Core) BrandResponse {
	userResponse := BrandResponse{
		ID:   int(userCore.ID),
		Name: userCore.Name,
	}
	return userResponse
}
