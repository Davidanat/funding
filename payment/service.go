package payment

import (
	"funding/user"
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	// "github.com/veritrans/go-midtrans"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (srv *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	// midclient := midtrans.NewClient()
	// midclient.ServerKey = ""
	// midclient.ClientKey = ""
	// midclient.APIEnvType = midtrans.Sandbox

	// snapGateway := midtrans.SnapGateway{
	// 	Client: midclient,
	// }

	// snapReq := &midtrans.SnapReq{
	// 	TransactionDetails: midtrans.TransactionDetails{
	// 		OrderID:  strconv.Itoa(transaction.ID) + transaction.Code,
	// 		GrossAmt: int64(transaction.Amount),
	// 	},
	// 	CustomerDetail: &midtrans.CustDetail{
	// 		FName: user.Name,
	// 		Email: user.Email,
	// 	},
	// }

	// snapResp, err := snapGateway.GetToken(snapReq)
	// if err != nil {
	// 	return "", err
	// }

	// 1. Initiate Snap client
	var s = snap.Client{}
	server := os.Getenv("SERVER_KEY")
	s.New(server, midtrans.Sandbox)

	// 2. Initiate Snap request
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.Code,
			GrossAmt: int64(transaction.Amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
	}

	// 3. Request create Snap transaction to Midtrans
	snapResp, err := s.CreateTransaction(req)
	if err != nil {
		return "", err
	}

	return snapResp.RedirectURL, nil
}
