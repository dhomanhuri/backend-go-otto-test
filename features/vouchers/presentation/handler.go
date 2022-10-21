package presentation

import (
	"fmt"
	"otto/test/features/vouchers"
	"otto/test/features/vouchers/presentation/dto"
	"otto/test/helper"

	"github.com/gin-gonic/gin"
)

type VoucherPres struct {
	Pres vouchers.Bussines
}

func VoucherHandler(pres vouchers.Bussines) VoucherPres {
	return VoucherPres{
		Pres: pres,
	}
}

func (Bb VoucherPres) VoucherInsert(c *gin.Context) {
	voucherReq := dto.VoucherRequest{}
	err := c.Bind(&voucherReq)
	fmt.Println(err)
	if err != nil {
		c.JSON(helper.BadRequest("error bidn"))
		c.Abort()
		return
	}
	core := vouchers.Core{
		Name:    voucherReq.Name,
		Price:   voucherReq.Price,
		BrandID: voucherReq.BrandID,
	}

	result := Bb.Pres.BussInsert(core)
	if result != nil {
		c.JSON(helper.BadRequest(result.Error()))
		c.Abort()
		return
	}
	c.JSON(helper.SuccessInsert())
}

func (Bb VoucherPres) GetVoucher(c *gin.Context) {
	id := c.Query("id")
	result, err := Bb.Pres.BussGet(id)
	if err != nil {
		c.JSON(helper.FailedBadRequestWithMSG(err.Error()))
		return
	}

	c.JSON(helper.SuccessGetData(dto.FromCore(result)))
}

func (Bb VoucherPres) GetAllVouchersByBrandsId(c *gin.Context) {
	id := c.Query("id")

	result, err := Bb.Pres.BussGetByBrandID(id)
	if err != nil {
		c.JSON(helper.FailedBadRequestWithMSG("error handler"))
	}

	c.JSON(helper.SuccessGetData(dto.FromCoreList(result)))
}
