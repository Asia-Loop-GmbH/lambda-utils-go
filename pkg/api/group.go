package api

import "github.com/asia-loop-gmbh/lambda-utils-go/v4/pkg/dbadmin"

type GroupDetails struct {
	Group  dbadmin.Group  `json:"group"`
	Orders []OrderDetails `json:"orders"`
}

type CreateGroupRequest struct {
	Store string `json:"store"`
}
