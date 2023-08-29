// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package server

import (
	"context"

	"go-service/internal/pb"
	"go-service/service/order/internal/logic/order_service"
	"go-service/service/order/internal/svc"
)

type OrderServiceServer struct {
	svcCtx *svc.ServiceContext
	__.UnimplementedOrderServiceServer
}

func NewOrderServiceServer(svcCtx *svc.ServiceContext) *OrderServiceServer {
	return &OrderServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServiceServer) PostOrderMes(ctx context.Context, in *__.OrderReq) (*__.WholeOrder, error) {
	l := orderservicelogic.NewPostOrderMesLogic(ctx, s.svcCtx)
	return l.PostOrderMes(in)
}

func (s *OrderServiceServer) GetOrder(ctx context.Context, in *__.Id) (*__.WholeOrder, error) {
	l := orderservicelogic.NewGetOrderLogic(ctx, s.svcCtx)
	return l.GetOrder(in)
}

func (s *OrderServiceServer) GetUserOrder(ctx context.Context, in *__.Id) (*__.UserOrderAll, error) {
	l := orderservicelogic.NewGetUserOrderLogic(ctx, s.svcCtx)
	return l.GetUserOrder(in)
}
