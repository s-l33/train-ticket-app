package api

import (
	"context"
	"testing"

	"github.com/s-l33/train-ticket-app/api/pb"
)

func TestPurchaseTicket(t *testing.T) {
	s := NewTrainTicketAppServer()

	req := &pb.PurchaseTicketRequest{
		From: "New York",
		To:   "Los Angeles",
		User: &pb.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@example.com",
		},
	}

	resp, err := s.PurchaseTicket(context.Background(), req)
	if err != nil {
		t.Errorf("failed to purchase ticket: %v", err)
	}

	if resp.Receipt == nil {
		t.Errorf("receipt is nil")
	}

	if resp.Receipt.From != req.From {
		t.Errorf("unexpected value for receipt.From: %v", resp.Receipt.From)
	}

	if resp.Receipt.To != req.To {
		t.Errorf("unexpected value for receipt.To: %v", resp.Receipt.To)
	}

	if resp.Receipt.User.FirstName != req.User.FirstName {
		t.Errorf("unexpected value for receipt.User.FirstName: %v", resp.Receipt.User.FirstName)
	}

	if resp.Receipt.User.LastName != req.User.LastName {
		t.Errorf("unexpected value for receipt.User.LastName: %v", resp.Receipt.User.LastName)
	}

	if resp.Receipt.User.Email != req.User.Email {
		t.Errorf("unexpected value for receipt.User.Email: %v", resp.Receipt.User.Email)
	}
}

func TestGetReceiptDetails(t *testing.T) {
	s := NewTrainTicketAppServer()

	req := &pb.PurchaseTicketRequest{
		From: "New York",
		To:   "Los Angeles",
		User: &pb.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@example.com",
		},
	}

	resp, err := s.PurchaseTicket(context.Background(), req)
	if err != nil {
		t.Errorf("failed to purchase ticket: %v", err)
	}

	receiptID := resp.Receipt.Id

	getReceiptDetailsReq := &pb.GetReceiptDetailsRequest{
		ReceiptID: receiptID,
	}

	getReceiptDetailsResp, err := s.GetReceiptDetails(context.Background(), getReceiptDetailsReq)
	if err != nil {
		t.Errorf("failed to get receipt details: %v", err)
	}

	if getReceiptDetailsResp.Receipt == nil {
		t.Errorf("receipt is nil")
	}

	if getReceiptDetailsResp.Receipt.Id != receiptID {
		t.Errorf("unexpected value for receipt.Id: %v", getReceiptDetailsResp.Receipt.Id)
	}

	if getReceiptDetailsResp.Receipt.From != req.From {
		t.Errorf("unexpected value for receipt.From: %v", getReceiptDetailsResp.Receipt.From)
	}

	if getReceiptDetailsResp.Receipt.To != req.To {
		t.Errorf("unexpected value for receipt.To: %v", getReceiptDetailsResp.Receipt.To)
	}

	if getReceiptDetailsResp.Receipt.User.FirstName != req.User.FirstName {
		t.Errorf("unexpected value for receipt.User.FirstName: %v", getReceiptDetailsResp.Receipt.User.FirstName)
	}

	if getReceiptDetailsResp.Receipt.User.LastName != req.User.LastName {
		t.Errorf("unexpected value for receipt.User.LastName: %v", getReceiptDetailsResp.Receipt.User.LastName)
	}

	if getReceiptDetailsResp.Receipt.User.Email != req.User.Email {
		t.Errorf("unexpected value for receipt.User.Email: %v", getReceiptDetailsResp.Receipt.User.Email)
	}
}
