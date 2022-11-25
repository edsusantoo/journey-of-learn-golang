package helper

import (
	"todos/model/domain"
	"todos/model/web/auth"
)

func ToAuthResponse(domain domain.User) auth.AuthResponse {
	return auth.AuthResponse{
		Id:        domain.Id,
		Username:  domain.Username,
		FirstName: domain.FirstName,
		LastName:  domain.LastName,
		Active:    domain.Active,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
