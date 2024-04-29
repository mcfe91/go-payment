package mm

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	pb "github.com/mmcferren/go-micro/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/status"
)

const (
	insertTransactionQuery = "INSERT INTO transaction (pid, src_user_id, dst_user_id, src_wallet_id, dst_wallet_id, src_account_id, dst_account_id, src_account_type, dst_account_type, final_dst_merchant_wallet_id, amount) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	selectTransactionQuery = "SELECT pid, src_user_id, dst_user_id, src_wallet_id, dst_wallet_id, src_account_id, dst_account_id, src_account_type, dst_account_type, final_dst_merchant_wallet_id, amount FROM transaction WHERE pid=?"
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

func fetchWallet(tx *sql.Tx, userID string) (wallet, error) {
	var w wallet

	stmt, err := tx.Prepare("SELECT id, user_id, wallet_type FROM wallet WHERE user_id=?")
	if err != nil {
		return w, status.Error(codes.Internal, err.Error())
	}

	err = stmt.QueryRow(userID).Scan(&w.ID, &w.userID, &w.walletType)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return w, status.Error(codes.InvalidArgument, err.Error())
		}
		return w, status.Error(codes.Internal, err.Error())
	}
	return w, nil
}

func fetchAccount(tx *sql.Tx, walletID int32, accountType string) (account, error) {
	var a account

	stmt, err := tx.Prepare("SELECT id, cents, account_type, wallet_id FROM account WHERE wallet_id=? AND account_type=?")
	if err != nil {
		return a, status.Error(codes.Internal, err.Error())
	}

	err = stmt.QueryRow(walletID, accountType).Scan(&a.ID, &a.cents, &a.accountType, &a.walletID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return a, status.Error(codes.InvalidArgument, err.Error())
		}
		return a, status.Error(codes.Internal, err.Error())
	}
	return a, nil
}

func transfer(tx *sql.Tx, srcAccount account, dstAccount account, amount int64) error {
	if srcAccount.cents < amount {
		return status.Error(codes.InvalidArgument, "not enough money")
	}

	// subtract money from src
	stmt, err := tx.Prepare("UPDATE account SET cents=? WHERE id=?")
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	_, err = stmt.Exec(srcAccount.cents - amount, srcAccount.ID)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	// add money to destination account
	stmt, err = tx.Prepare("UPDATE account SET cents=? WHERE id=?")
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	_, err = stmt.Exec(dstAccount.cents + amount, dstAccount.ID)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func createTransaction(tx *sql.Tx, srcAccount account, dstAccount account, srcWallet wallet, dstWallet wallet, merchantWallet wallet, amount int64) (string, error) {
	pid := uuid.NewString()

	stmt, err := tx.Prepare(insertTransactionQuery)
	if err != nil {
		return "", status.Error(codes.Internal, err.Error())
	}

	_, err = stmt.Exec(pid, srcAccount.walletID, dstAccount.walletID, srcAccount.ID, dstAccount.ID, srcWallet.ID, dstWallet.ID, srcAccount.accountType, dstAccount.accountType, merchantWallet.ID, amount)
	if err != nil {
		return "", status.Error(codes.Internal, err.Error())
	}

	return pid, nil
}