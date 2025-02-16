package main

import (
	"context"
	"errors"
	"grpc-api/data"
	pb "grpc-api/proto"
	"io"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type stockService struct {
	pb.UnimplementedStockServiceServer
}

func (s *stockService) ListStocks(ctx context.Context, _ *emptypb.Empty) (*pb.StockCodes, error) {
	configs := data.GetStockConfig()

	codes := []string{}
	for code := range configs {
		codes = append(codes, code)
	}

	response := &pb.StockCodes{
		StockCodes: codes,
	}

	return response, nil
}

func (s *stockService) ToggleStocks(stream pb.StockService_ToggleStocksServer) error {
	toggles := map[string]bool{}

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}

		code := req.StockCode
		isEnable := req.IsEnabled

		toggles[code] = isEnable
	}

	for code, isEnabled := range toggles {
		data.ToggleStock(code, isEnabled)
	}

	configs := data.GetStockConfig()

	subscribed := []string{}
	for code, isEnabled := range configs {
		if isEnabled {
			subscribed = append(subscribed, code)
		}
	}

	return stream.SendAndClose(&pb.StockCodes{
		StockCodes: subscribed,
	})

}

func (s *stockService) ListSubscriptions(_ *emptypb.Empty, stream pb.StockService_ListSubscriptionsServer) error {
	configs := data.GetStockConfig()

	for code, isEnabled := range configs {
		if isEnabled {
			stream.Send(&pb.StockCode{
				StockCode: code,
			})
		}
	}

	return nil

}

func (s *stockService) LiveStock(stream pb.StockService_LiveStockServer) error {
	go func() {
		for {
			req, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				return
			}

			code := req.StockCode
			isEnabled := req.IsEnabled

			data.ToggleStock(code, isEnabled)
		}
	}()

	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			time.Sleep(1 * time.Second)

			configs := data.GetStockConfig()
			for code, isEnabled := range configs {
				if isEnabled {
					history := data.GetStockPrice(code)
					if len(history) == 0 {
						continue
					}

					latestPrice := history[len(history)-1]

					stream.Send(&pb.StockPrices{
						StockPrices: map[string]*pb.StockPrice{
							code: {
								Price:     latestPrice.Price,
								Timestamp: timestamppb.New(latestPrice.Timestamp),
							},
						},
					})
				}
			}
		}
	}
}
