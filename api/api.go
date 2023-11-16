package api

import (
	"context"
	"fmt"
	"github.com/s-l33/train-ticket-app/api/pb"
	"math/rand"
	"strings"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type Receipt struct {
	ID      string  `json:"id"`
	From    string  `json:"from"`
	To      string  `json:"to"`
	User    User    `json:"user"`
	Price   float32 `json:"price"`
	Section string  `json:"section"`
}

var (
	sections = strings.Split("A,B", ",")
	price    = 20.0
)

type trainTicketAppServer struct {
	pb.UnimplementedTrainTicketAppServer
	receipts map[string]*Receipt
	users    map[string]*User
}

func NewTrainTicketAppServer() pb.TrainTicketAppServer {
	return &trainTicketAppServer{
		receipts: make(map[string]*Receipt),
		users:    make(map[string]*User),
	}
}

func (s *trainTicketAppServer) PurchaseTicket(ctx context.Context, req *pb.PurchaseTicketRequest) (*pb.PurchaseTicketResponse, error) {
	from := req.From
	to := req.To
	user := User{
		ID:        fmt.Sprintf("U-%d", len(s.users)+1),
		FirstName: req.User.FirstName,
		LastName:  req.User.LastName,
		Email:     req.User.Email,
	}

	// Select randomly from sections
	section := sections[rand.Intn(len(sections))]

	receipt := Receipt{
		ID:      fmt.Sprintf("RT-%d", len(s.receipts)+1),
		From:    from,
		To:      to,
		User:    user,
		Price:   float32(price),
		Section: section,
	}

	s.receipts[receipt.ID] = &receipt
	s.users[user.ID] = &user

	resp := pb.PurchaseTicketResponse{
		Receipt: &pb.Receipt{
			Id:      receipt.ID,
			From:    receipt.From,
			To:      receipt.To,
			Price:   receipt.Price,
			Section: receipt.Section,
			User: &pb.User{
				Id:        &receipt.User.ID,
				FirstName: receipt.User.FirstName,
				LastName:  receipt.User.LastName,
				Email:     receipt.User.Email,
			},
		},
	}

	return &resp, nil
}

func (s *trainTicketAppServer) GetReceiptDetails(ctx context.Context, req *pb.GetReceiptDetailsRequest) (*pb.GetReceiptDetailsResponse, error) {
	receiptID := req.ReceiptID

	receipt, ok := s.receipts[receiptID]
	if !ok {
		return nil, fmt.Errorf("receipt with ID %s not found", receiptID)
	}

	resp := pb.GetReceiptDetailsResponse{
		Receipt: &pb.Receipt{
			Id:      receipt.ID,
			From:    receipt.From,
			To:      receipt.To,
			Section: receipt.Section,
			Price:   receipt.Price,
			User: &pb.User{
				Id:        &receipt.User.ID,
				FirstName: receipt.User.FirstName,
				LastName:  receipt.User.LastName,
				Email:     receipt.User.Email,
			},
		},
	}

	return &resp, nil
}

func (s *trainTicketAppServer) GetUsersBySection(ctx context.Context, req *pb.GetUsersBySectionRequest) (*pb.GetUsersBySectionResponse, error) {
	section := req.Section

	users := make([]*pb.User, 0)
	for _, r := range s.receipts {
		if r.Section == section {
			users = append(
				users,
				&pb.User{Id: &r.User.ID, FirstName: r.User.FirstName, LastName: r.User.LastName, Email: r.User.Email},
			)
		}
	}

	resp := pb.GetUsersBySectionResponse{
		Users: users,
	}

	return &resp, nil
}

func (s *trainTicketAppServer) RemoveUser(ctx context.Context, req *pb.RemoveUserRequest) (*pb.RemoveUserResponse, error) {
	userID := req.UserID

	_, ok := s.users[userID]
	if !ok {
		return nil, fmt.Errorf("user with ID %s not found", userID)
	}

	delete(s.users, userID)

	resp := pb.RemoveUserResponse{}

	return &resp, nil
}

func (s *trainTicketAppServer) ModifyUserSeat(ctx context.Context, req *pb.ModifyUserSeatRequest) (*pb.ModifyUserSeatResponse, error) {
	receiptID := req.ReceiptID
	newSection := req.NewSection

	receipt, ok := s.receipts[receiptID]
	if !ok {
		return nil, fmt.Errorf("receipt with ID %s not found", receiptID)
	}

	oldSection := receipt.Section
	receipt.Section = newSection

	resp := pb.ModifyUserSeatResponse{
		OldSection: oldSection,
		NewSection: newSection,
	}

	return &resp, nil
}
