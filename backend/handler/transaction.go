package handler

import (
	"time"

	"github.com/aliirsyaadn/neobank-technical-test/entity"
	"github.com/aliirsyaadn/neobank-technical-test/service"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"go.uber.org/zap"
)

type transactionHandler struct {
	log                *zap.SugaredLogger
	transactionService service.TransactionService
}

func NewTransactionHandler(
	log *zap.SugaredLogger,
	transactionService service.TransactionService,
) TransactionHandler {
	return &transactionHandler{
		log:                log,
		transactionService: transactionService,
	}
}

func (h *transactionHandler) Get(c *gin.Context) {
	referenceNo := c.Param("reference_no")

	res, err := h.transactionService.Get(c, referenceNo)
	if err != nil {
		c.JSONP(400, entity.Response{
			Message: "Failed to get",
			Error:   err.Error(),
		})
		return
	}

	c.JSONP(200, entity.Response{
		Message: "success",
		Data:    res,
	})
}

func (h *transactionHandler) GetList(c *gin.Context) {
	var req entity.GetListTransactionRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSONP(400, entity.Response{
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	if req.EndDateStr != "" {
		req.EndDate, err = time.Parse("2006-01-02", req.EndDateStr)
		if err != nil {
			c.JSONP(400, entity.Response{
				Message: "Invalid request",
				Error:   err.Error(),
			})
			return
		}
	}

	if req.StartDateStr != "" {
		req.StartDate, err = time.Parse("2006-01-02", req.StartDateStr)
		if err != nil {
			c.JSONP(400, entity.Response{
				Message: "Invalid request",
				Error:   err.Error(),
			})
			return
		}
	}

	res, err := h.transactionService.GetList(c, req)
	if err != nil {
		c.JSONP(400, entity.Response{
			Message: "Failed to get list",
			Error:   err.Error(),
		})
		return
	}

	c.JSONP(200, entity.Response{
		PaginationData: &res.Pagination,
		Message:        "success",
		Data:           res.Transactions,
	})
}

func (h *transactionHandler) Create(c *gin.Context) {
	var req entity.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSONP(400, entity.Response{
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	req.MakerUserID = c.Keys["id"].(string)

	res, err := h.transactionService.Create(c, req)
	if err != nil {
		c.JSONP(400, entity.Response{
			Message: "Failed to create",
			Error:   err.Error(),
		})
		return
	}

	c.JSONP(200, entity.Response{
		Message: "success",
		Data:    res,
	})
}

func (h *transactionHandler) Update(c *gin.Context) {
	referenceNo := c.Param("reference_no")
	if referenceNo == "" {
		c.JSONP(400, entity.Response{
			Message: "Invalid request",
			Error:   "Reference no is required",
		})
		return
	}

	var req entity.UpdateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSONP(400, entity.Response{
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	req.ReferenceNo = referenceNo

	res, err := h.transactionService.Update(c, req)
	if err != nil {
		c.JSONP(400, entity.Response{
			Message: "Failed to update",
			Error:   err.Error(),
		})
		return
	}

	c.JSONP(200, entity.Response{
		Message: "success",
		Data:    res,
	})
}

func (h *transactionHandler) ValidateTransferRecordCSV(c *gin.Context) {
	_, file, err := c.Request.FormFile("file")
	if err != nil {
		c.JSONP(400, entity.Response{
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSONP(400, entity.Response{
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}
	defer f.Close()

	var records []entity.TransferRecordData

	err = gocsv.Unmarshal(f, &records)
	if err != nil {
		c.JSONP(400, entity.Response{
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	res, err := h.transactionService.ValidateTransferRecordCSV(c, records)
	if err != nil {
		c.JSONP(400, entity.Response{
			Message: "Failed to validate",
			Error:   err.Error(),
		})
		return
	}

	if len(res.Errors) > 0 {
		c.JSONP(400, entity.Response{
			Message: "success with errors",
			Data:    res,
		})
		return
	}

	c.JSONP(200, entity.Response{
		Message: "success",
		Data:    res,
	})
}
