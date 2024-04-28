package mm

import (
	"context"
	"database/sql"

	pb "github.com/mmcferren/go-micro/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/status"
)

type Implementation struct {
	db * sql.DB
	pb.UnimplementedMoneyMovementServiceServer
}

func NewMoneyMovementImplementation(db *sql.DB) *Implementation {
	return &Implementation{db: db}
}

func (this *Implementation) Authorize(ctx context.Context, authorizePayload *pb.AuthorizePayload) (*pb.AuthorizeResponse, error) {
	if authorizePayload.GetCurrency() != "USD" {
		return nil, status.Error(codes.InvalidArgument, "only accepts USD")
	}

	// Begin transaction
	tx, err := this.db.Begin()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	merchantWallet, err := fetchWallet(tx, authorizePayload.MerchantWalletUserId)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, err
	}

	customerWallet, err := fetchWallet(tx, authorizePayload.CustomerWalletUserId)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, err
	}

	srcAccount, err := fetchAccount(tx, customerWallet.ID, "DEFAULT")
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, err
	}

	dstAccount, err := fetchAccount(tx, customerWallet.ID, "PAYMENT")
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, err
	}

	err = transfer(tx, srcAccount, dstAccount, authorizePayload.Cents)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, err
	}

	pid, err := createTransaction(tx, srcAccount, dstAccount, customerWallet, customerWallet, merchantWallet, authorizePayload.Cents)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, err
	}

	// End transaction
	err = tx.Commit()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.AuthorizeResponse{Pid: pid}, nil
}