package presentation

import (
	"otto/test/features/transactions"
	"otto/test/features/transactions/presentation/dto"
	"otto/test/helper"

	"github.com/gin-gonic/gin"
)

type TransactionsPres struct {
	Pres transactions.Bussines
}

func TransactionHandler(pres transactions.Bussines) TransactionsPres {
	return TransactionsPres{
		Pres: pres,
	}
}

func (Bb TransactionsPres) TransactionInsert(c *gin.Context) {
	transReq := dto.TransactionRequest{}
	err := c.Bind(&transReq)
	if err != nil {
		c.JSON(helper.BadRequest("error bidn"))
		c.Abort()
		return
	}
	core := transactions.Core{
		Desc: transReq.Desc,
	}

	result := Bb.Pres.BussInsert(core)
	if result != nil {
		c.JSON(helper.BadRequest(result.Error()))
		c.Abort()
		return
	}
	c.JSON(helper.SuccessInsert())
}
