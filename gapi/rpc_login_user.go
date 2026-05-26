package gapi

import (
	"context"
	"errors"
	db "github/atulkumar0001/Bank/db/sqlc"
	"github/atulkumar0001/Bank/pb"
	"github/atulkumar0001/Bank/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)


func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest)(*pb.LoginUserResponse, error){

	user, err := server.store.GetUser(ctx,req.GetUsername())

	if err != nil{
		if errors.Is(err, db.ErrRecordNotFound) {
			return	nil,status.Error(codes.NotFound, "Record Not Found")
		}
		return	nil,status.Error(codes.Internal, "Something went wrong")
	}

	err = util.CheckPassword(req.Password,user.HashedPassword)

	if err != nil{
		return	nil,status.Error(codes.NotFound, "Incorrect password")
	}

	accessToken, accessTokenPayload, err := server.tokenMaker.CreateToken(req.Username,server.config.AcccessTokenDuration)

	if err != nil{
		return	nil,status.Error(codes.Internal, "Something went wrong")
	}

	refreshToken, refreshTokenPayload, err := server.tokenMaker.CreateToken(req.Username,server.config.RefreshTokenDuration)

	if err != nil{
		return	nil,status.Error(codes.Internal, "Something went wrong")		
	}

	mtdt := server.extractMetadata(ctx)
	session, err := server.store.CreateSession(ctx,db.CreateSessionParams{
		ID: refreshTokenPayload.ID,
		Username: req.Username,
		RefreshToken: refreshToken,
		UserAgent: mtdt.UserAgent,
		ClientIp: mtdt.ClientIp,
		IsBlocked: false,
		ExpiresAt: refreshTokenPayload.ExpiredAt,
	})

	if err != nil{
		return	nil,status.Error(codes.Internal, "Something went wrong")
	}

	res := &pb.LoginUserResponse{
		User: convertUser(user),
		SessionId: session.ID.String(),
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		AccessTokenExpiresAt: timestamppb.New(accessTokenPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshTokenPayload.ExpiredAt),
	}
	return res,nil;
}