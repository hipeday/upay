package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron/v2"
	"github.com/hipeday/upay/internal/middleware"
	"github.com/hipeday/upay/internal/web3/trc20/trongrid"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Order struct {
	Id       int64     `json:"id,omitempty"`
	CreateAt time.Time `json:"create_at"`
	// Status create waiting_payment success close
	Status          string     `json:"status,omitempty"`
	Amount          int64      `json:"amount,omitempty"`
	FromAddress     *string    `json:"from_address,omitempty"`
	ToAddress       string     `json:"to_address,omitempty"`
	ContractAddress string     `json:"contract_address,omitempty"`
	ConfirmAt       *time.Time `json:"confirm_at,omitempty"`
	TransactionId   *string    `json:"transaction_id,omitempty"`
}

type OrderRoute struct {
	orders map[int64]Order
}

func setupOrder(db *sqlx.DB, engine *gin.Engine) {
	route := OrderRoute{}
	route.orders = make(map[int64]Order)
	// register auth api
	route.Register(engine, middleware.RequestLoggingMiddleware(), middleware.BearerAuthorizationMiddleware())
}

func (o OrderRoute) Register(engine *gin.Engine, middlewares ...gin.HandlerFunc) {
	routerGroup := engine.Group(order)
	routes := routerGroup.Use(middlewares...)
	routes.POST("", o.create)
	routes.GET("", o.get)
}

// create 创建支付订单
func (o OrderRoute) create(context *gin.Context) {
	value := context.Query("amount")
	var amount float64
	if value != "" {
		valueFloat, err := strconv.ParseFloat(value, 10)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			context.Abort()
		}
		amount = valueFloat * 1000000
	}
	current := Order{
		Id:              time.Now().UnixNano(),
		CreateAt:        time.Now(),
		Status:          "create",
		Amount:          int64(amount),
		FromAddress:     nil,
		ToAddress:       "TYYi4mhbkUwcezbnXP5UyKUbcS15Kx1qt8",
		ContractAddress: "TXLAQ63Xg1NAzckPwKHvzw7CSEmLMEqcdj",
	}
	o.orders[current.Id] = current
	context.JSON(http.StatusOK, current)
	o.confirmOrder(current.Id)
}

// get 获取支付订单
func (o OrderRoute) get(context *gin.Context) {
	context.JSON(http.StatusOK, o.orders)
}

// sync 订单状态同步器 定时从 trongrid 获取订单状态
func (o OrderRoute) sync(orderId int64) {
	log.Printf("订单状态同步器开始执行 orderId: %d", orderId)
	current := o.orders[orderId]
	if current.Status != "create" {
		return
	}
	client := trongrid.NewTronGridClient("https://nile.trongrid.io")
	onlyTo := true
	onlyConfirmed := true
	payload := trongrid.GetContractTransactionsPayload{
		Address:         &current.ToAddress,
		OnlyConfirmed:   &onlyConfirmed,
		OnlyTo:          &onlyTo,
		OnlyFrom:        nil,
		Limit:           nil,
		Fingerprint:     nil,
		OrderBy:         nil,
		MinTimestamp:    nil,
		MaxTimestamp:    nil,
		ContractAddress: &current.ContractAddress,
	}
	transactions, err := client.GetContractTransactions(payload)
	if err != nil {
		log.Fatal(err)
	}

	response := transactions.(trongrid.Response[[]trongrid.ContractTransaction])

	success := response.Success
	if success {
		if response.Data != nil {
			data := *response.Data
			for i := range data {
				transaction := data[i]
				value := transaction.Value
				amount, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				if transaction.Type == "Transfer" && transaction.To == current.ToAddress && amount == current.Amount && transaction.BlockTimestamp > current.CreateAt.UnixMilli() {
					current.Status = "confirmed"
					current.FromAddress = &transaction.From
					unix := time.UnixMilli(transaction.BlockTimestamp)
					current.ConfirmAt = &unix
					current.TransactionId = &transaction.TransactionId
					o.orders[current.Id] = current
				}
			}
		}
	}
}

func (o OrderRoute) confirmOrder(orderId int64) {
	// 创建一个新的调度器
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}

	j, err := s.NewJob(
		gocron.DurationJob(
			1*time.Second,
		),
		gocron.NewTask(
			o.sync,
			orderId,
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	s.Start()

	// each job has a unique id
	log.Print("订单id: ", orderId, " 已经开始执行。 任务id为: ", j.ID())

}
