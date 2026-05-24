package gapi

import (
	db "github/atulkumar0001/Bank/db/sqlc"
	"github/atulkumar0001/Bank/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.Users) (*pb.User){
	return &pb.User{
		Username: user.Username,
		FullName: user.FullName,
		Email: user.Email,
		PasswordLastUpdated: timestamppb.New(user.PasswordLastUpdated),
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}