package presentation

import (
	"fmt"
	"otto/test/features/redemptions"
	"otto/test/features/redemptions/presentation/dto"
	"otto/test/helper"

	"github.com/gin-gonic/gin"
)

type RedemptionsPres struct {
	Pres redemptions.Bussines
}

func RedemptionHandler(pres redemptions.Bussines) RedemptionsPres {
	return RedemptionsPres{
		Pres: pres,
	}
}

func (Bb RedemptionsPres) RedemptionInsert(c *gin.Context) {
	redempReq := dto.RedemptionRequest{}
	err := c.Bind(&redempReq)
	fmt.Println("asas", err)
	if err != nil {
		c.JSON(helper.BadRequest("error bidn"))
		c.Abort()
		return
	}
	core := redemptions.Core{
		VoucherID:     redempReq.VoucherID,
		TransactionID: redempReq.TransactionID,
		Quantity:      redempReq.Quantity,
	}

	result := Bb.Pres.BussInsert(core)
	if result != nil {
		c.JSON(helper.BadRequest(result.Error()))
		c.Abort()
		return
	}
	c.JSON(helper.SuccessInsert())
}

func (Bb RedemptionsPres) GetAllRedemptionsByBrandsId(c *gin.Context) {
	transactionId := c.Query("transactionId")

	total, result, err := Bb.Pres.BussGetByTransactionID(transactionId)
	if err != nil {
		c.JSON(helper.FailedBadRequestWithMSG("error handler"))
		return
	}
	data := dto.ParrentData{}
	data.Vouchers = dto.FromCoreList(result)
	data.Total = total
	c.JSON(helper.SuccessGetData(data))
}
