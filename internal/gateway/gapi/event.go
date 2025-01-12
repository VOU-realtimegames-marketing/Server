package gapi

import (
	"VOU-Server/proto/gen"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateEvent(ctx context.Context, req *gen.CreateEventRequest) (*gen.CreateEventResponse, error) {
	res, err := server.AuthorizeUser(ctx, &gen.AuthorizeRequest{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "unauthorized action: %s", err)
	}

	if res.User.Role != counterpartRole {
		return nil, status.Errorf(codes.PermissionDenied, "unauthorized action: %s", err)
	}

	req.Owner = res.User.Username
	return server.eventClient.CreateEvent(ctx, req)
}

func (server *Server) GetAllEvents(ctx context.Context, req *gen.GetAllEventsRequest) (*gen.GetAllEventsResponse, error) {
	return server.eventClient.GetAllEvents(ctx, req)
}

func (server *Server) GetAllEventsOfOwner(ctx context.Context, req *gen.GetEventsOfOwnerRequest) (*gen.GetEventsOfOwnerResponse, error) {
	return server.eventClient.GetAllEventsOfOwner(ctx, req)
}

func (server *Server) GetEventById(ctx context.Context, req *gen.GetEventByIdRequest) (*gen.GetEventByIdResponse, error) {
	return server.eventClient.GetEventById(ctx, req)
}

func (server *Server) UpdateEventStatus(ctx context.Context, req *gen.UpdateEventStatusRequest) (*gen.UpdateEventStatusResponse, error) {
	res, err := server.AuthorizeUser(ctx, &gen.AuthorizeRequest{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "unauthorized action: %s", err)
	}

	if res.User.Role != adminRole {
		return nil, status.Errorf(codes.PermissionDenied, "unauthorized action: %s", err)
	}

	return server.eventClient.UpdateEventStatus(ctx, req)
}

func (server *Server) GetQuizzesByEventId(ctx context.Context, req *gen.GetQuizzesByEventIdRequest) (*gen.GetQuizzesByEventIdResponse, error) {
	res, err := server.AuthorizeUser(ctx, &gen.AuthorizeRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	if res.User.Role != counterpartRole && res.User.Role != adminRole {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	req.Owner = res.User.Username
	req.IsAdmin = res.User.Role == adminRole
	return server.eventClient.GetQuizzesByEventId(ctx, req)
}
