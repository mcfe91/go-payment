package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	authpb "github.com/mmcferren/go-micro/auth"
	mmpb "github.com/mmcferren/go-micro/money_movement"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var authClient authpb.AuthServiceClient
var mmClient mmpb.MoneyMovementServiceClient

func main() {
	authConn, err := grpc.Dial("auth:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := authConn.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	authClient = authpb.NewAuthServiceClient(authConn)
	
	mmConn, err := grpc.Dial("money-movement:7000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := mmConn.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	mmClient = mmpb.NewMoneyMovementServiceClient(mmConn)

	http.HandleFunc("/login", login)
	http.HandleFunc("/customer/payment/authorize", customerPaymentAuthorize)
	http.HandleFunc("/customer/payment/capture", customerPaymentCapture)

	fmt.Println("listening on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	// TODO: Don't store plain text pws!
	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	ctx := context.Background()
	token, err := authClient.GetToken(ctx, &authpb.Credentials{
		UserName: username,
		Password: password,
	})
	if err != nil {
		_, writeErr := w.Write(([]byte(err.Error())))
		if writeErr != nil {
			log.Println(writeErr)
		}
		return
	}
	_, err = w.Write([]byte(token.Jwt))  
	if err != nil {
		log.Println(err)
	}
	
}
func customerPaymentAuthorize(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get(("Authorization"))
	if authHeader == "" {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	ctx := context.Background()
	_, err := authClient.ValidateToken(ctx, &authpb.Token{Jwt: token})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	type authorizePayload struct {
		CustomerWalletUserId string `json:"customer_wallet_user_id"`
		MerchantWalletUserId string `json:"merchant_wallet_user_id"`
		Cents int64 `json:"cents"`
		Currency string `json:"currency"`
	}

	var payload authorizePayload
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusInternalServerError)
		return
	}

	ctx = context.Background()
	ar, err := mmClient.Authorize(ctx, &mmpb.AuthorizePayload{
		CustomerWalletUserId: payload.CustomerWalletUserId,
		MerchantWalletUserId: payload.MerchantWalletUserId,
		Cents: payload.Cents,
		Currency: payload.Currency,
	})
	if err != nil {
		_, writeErr := w.Write(([]byte(err.Error())))
		if writeErr != nil {
			log.Println(writeErr)
		}
		return
	}

	type response struct {
		Pid string `json:"pid"`
	}

	resp := response{
		Pid: ar.Pid,
	}

	responseJSON, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseJSON)
	if err != nil {
		log.Println(err)
	}
}
func customerPaymentCapture(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get(("Authorization"))
	if authHeader == "" {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	ctx := context.Background()
	_, err := authClient.ValidateToken(ctx, &authpb.Token{Jwt: token})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	type capturePayload struct {
		Pid string `json:"pid"`
	}

	var payload capturePayload
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusInternalServerError)
		return
	}
 
	ctx = context.Background()
	_, err = mmClient.Capture(ctx, &mmpb.CapturePayload{Pid: payload.Pid})
	if err != nil {
		_, writeErr := w.Write(([]byte(err.Error())))
		if writeErr != nil {
			log.Println(writeErr)
		}
		return
	}

	w.WriteHeader(http.StatusOK)	
}