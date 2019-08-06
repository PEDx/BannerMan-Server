package user

import (
	"BannerMan-Server/pkg/errno"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type CreateResponse struct {
	Username string `json:"username"`
}

func (r *CreateRequest) checkParam() error {
	if r.Username == "" {
		return errno.New(errno.ErrValidation, nil).Add("username is empty.")
	}
	if r.Password == "" {
		return errno.New(errno.ErrValidation, nil).Add("password is empty.")
	}
	return nil
}
