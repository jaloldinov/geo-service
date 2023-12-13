package handler

import (
	"geo/pkg/logger"
	"geo/storage"
)

type Handler struct {
	storage storage.StorageI
	hub     *Hub
	log     logger.LoggerI
}

func NewHandler(strg storage.StorageI, hub *Hub, loger logger.LoggerI) *Handler {
	return &Handler{storage: strg, hub: hub, log: loger}
}

//func NewHandler(strg storage.StorageI, loger logger.LoggerI) *Handler {
//	return &Handler{storage: strg, log: loger}
//}
