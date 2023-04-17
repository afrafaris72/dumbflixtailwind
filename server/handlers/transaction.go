package handlers

import (
	dto "dumbflix/dto/result"
	transactionDto "dumbflix/dto/transaction"
	"dumbflix/models"
	"dumbflix/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) FindTransactions(c echo.Context) error {
	transactions, err := h.TransactionRepository.FindTransactions()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transactions})
}

func (h *handlerTransaction) GetTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transaction})
}

func (h *handlerTransaction) CreateTransaction(c echo.Context) error {
	request := new(transactionDto.CreateTransactionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	// request.BuyerID = int(userId)
	// request.Status = "pending"

	transaction := models.Transaction{
		StartDate: request.StartDate,
		DueDate:   request.DueDate,
		UserID:    int(userId),
		Status:    request.Status,
	}

	data, err := h.TransactionRepository.CreateTransaction(transaction)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	// var s = snap.Client{}
	// s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	// req := &snap.Request{
	// 	TransactionDetails: midtrans.TransactionDetails{
	// 		OrderID:  strconv.Itoa(dataTransactions.ID),
	// 		GrossAmt: int64(dataTransactions.Price),
	// 	},
	// 	CreditCard: &snap.CreditCardDetails{
	// 		Secure: true,
	// 	},
	// 	CustomerDetail: &midtrans.CustomerDetails{
	// 		FName: dataTransactions.Buyer.Name,
	// 		Email: dataTransactions.Buyer.Email,
	// 	},
	// }

	// snapResp, _ := s.CreateTransaction(req)

	// return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: snapResp})

	transaction, _ = h.TransactionRepository.GetTransaction(transaction.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)})
}

func (h *handlerTransaction) UpdateTransaction(c echo.Context) error {
	request := new(transactionDto.UpdateTransactionRequest)

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.StartDate != 0 {
		transaction.StartDate = request.StartDate
	}

	if request.DueDate != 0 {
		transaction.DueDate = request.DueDate
	}

	if request.UserID != 0 {
		transaction.UserID = request.UserID
	}

	if request.Status != "" {
		transaction.Status = request.Status
	}

	fmt.Println(transaction)

	data, err := h.TransactionRepository.UpdateTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)})
}

func (h *handlerTransaction) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.TransactionRepository.DeleteTransaction(transaction, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)})
}

// func (h *handlerTransaction) Notification(c echo.Context) error {
// 	var notificationPayload map[string]interface{}
// 	if err := c.Bind(&notificationPayload); err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	transactionStatus := notificationPayload["transaction_status"].(string)
// 	fraudStatus := notificationPayload["fraud_status"].(string)
// 	orderId := notificationPayload["order_id"].(string)
// 	order_id, _ := strconv.Atoi(orderId)
// 	fmt.Print("ini payloadnya", notificationPayload)
// 	if transactionStatus == "capture" {
// 		if fraudStatus == "challenge" {
// 			// TODO set transaction status on your database to 'challenge'
// 			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
// 			h.TransactionRepository.UpdateTransaction("pending", order_id)
// 		} else if fraudStatus == "accept" {
// 			// TODO set transaction status on your database to 'success'
// 			h.TransactionRepository.UpdateTransaction("success", order_id)
// 		}
// 	} else if transactionStatus == "settlement" {
// 		// TODO set transaction status on your databaase to 'success'
// 		h.TransactionRepository.UpdateTransaction("success", order_id)
// 	} else if transactionStatus == "deny" {
// 		// TODO you can ignore 'deny', because most of the time it allows payment retries
// 		// and later can become success
// 		h.TransactionRepository.UpdateTransaction("failed", order_id)
// 	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
// 		// TODO set transaction status on your databaase to 'failure'
// 		h.TransactionRepository.UpdateTransaction("failed", order_id)
// 	} else if transactionStatus == "pending" {
// 		// TODO set transaction status on your databaase to 'pending' / waiting payment
// 		h.TransactionRepository.UpdateTransaction("pending", order_id)
// 	}
// 	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: notificationPayload})
// }

func convertResponseTransaction(u models.Transaction) models.TransactionResponse {
	return models.TransactionResponse{
		ID:        u.ID,
		StartDate: u.StartDate,
		DueDate:   u.DueDate,
		User:      u.User,
		Status:    u.Status,
	}
}
