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

func convertVouchers(vouchers []db.Voucher) []*gen.Voucher {
	genVouchers := make([]*gen.Voucher, len(vouchers))
	for i, voucher := range vouchers {
		genVouchers[i] = &gen.Voucher{
			Id:        voucher.ID,
			EventId:   voucher.EventID,
			QrCode:    voucher.QrCode.String,
			Type:      voucher.Type,
			Status:    voucher.Status,
			Discount:  voucher.Discount,
			ExpiresAt: timestamppb.New(voucher.ExpiresAt),
		}
	}
	return genVouchers
}

func (server *Server) GetMyVouchers(ctx context.Context, req *gen.GetMyVouchersRequest) (*gen.GetMyVouchersResponse, error) {
	vouchers, err := server.store.ListMyVouchers(ctx, req.GetUsername())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "no voucher found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get vouchers: %s", err)
	}

	genVouchers := convertVouchers(vouchers)
	rsp := &gen.GetMyVouchersResponse{
		Vouchers: genVouchers,
	}

	return rsp, nil
}
