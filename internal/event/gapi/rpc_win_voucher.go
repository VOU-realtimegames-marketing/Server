package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/proto/gen"
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) WinVoucher(ctx context.Context, req *gen.WinVoucherRequest) (*gen.WinVoucherResponse, error) {
	event, err := server.store.GetEventById(ctx, req.GetEventId())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "event not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get event: %s", err)
	}

	voucher, err := server.store.CreateVoucher(ctx, db.CreateVoucherParams{
		EventID:   event.ID,
		ExpiresAt: event.EndTime,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create voucher: %s", err)
	}

	_, err = server.store.CreateVoucherOwner(ctx, db.CreateVoucherOwnerParams{
		VoucherID: voucher.ID,
		Username:  req.GetUsername(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create voucher owner: %s", err)
	}

	rsp := &gen.WinVoucherResponse{
		Voucher: &gen.Voucher{
			Id:        voucher.ID,
			EventId:   voucher.EventID,
			QrCode:    voucher.QrCode.String,
			Type:      voucher.Type,
			Status:    voucher.Status,
			Discount:  voucher.Discount,
			ExpiresAt: timestamppb.New(voucher.ExpiresAt),
		},
	}

	return rsp, nil
}
