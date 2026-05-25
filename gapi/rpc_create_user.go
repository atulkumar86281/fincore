package gapi

import (
	"context"
	db "github/atulkumar0001/Bank/db/sqlc"
	"github/atulkumar0001/Bank/pb"
	"github/atulkumar0001/Bank/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPass, err := util.HashedPassword(req.GetPassword())
	if err != nil{
		return	nil,status.Error(codes.Internal, "Something went wrong while hashing")
	}

	arg := db.CreateUserParams{
		Username: req.GetUsername(),
		HashedPassword: hashedPass,
		FullName: req.GetFullName(),
		Email: req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx,arg)

	if err != nil{
		if db.ErrorCode(err) == db.UniqueViolation {
			return	nil,status.Error(codes.AlreadyExists, "Username already exist")
		}
		return	nil,status.Error(codes.Internal, "Something went wrong while creating user")
	}

	rsp := &pb.CreateUserResponse{
		User: convertUser(user),
	}

	return rsp, nil
}
