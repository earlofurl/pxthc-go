package sqlc

import (
	"context"
	"github.com/earlofurl/pxthc/util"
	"github.com/gobuffalo/nulls"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomOrder(t *testing.T) Order {
	arg := CreateOrderParams{
		ScheduledPackDateTime:     time.Now(),
		ScheduledShipDateTime:     time.Now(),
		ScheduledDeliveryDateTime: time.Now(),
		ActualPackDateTime:        nulls.NewTime(time.Now()),
		ActualShipDateTime:        nulls.NewTime(time.Now()),
		ActualDeliveryDateTime:    nulls.NewTime(time.Now()),
		OrderTotal:                decimal.NewFromFloat(util.RandomDecimalTimes100()),
		Notes:                     util.RandomString(10),
		Status:                    util.RandomString(10),
		CustomerName:              util.RandomString(10),
	}

	order, err := testQueries.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	// TODO: look into how to accurately test the time.Time fields
	//require.Equal(t, arg.ScheduledPackDateTime, order.ScheduledPackDateTime)
	//require.Equal(t, arg.ScheduledShipDateTime, order.ScheduledShipDateTime)
	//require.Equal(t, arg.ScheduledDeliveryDateTime, order.ScheduledDeliveryDateTime)
	//require.Equal(t, arg.ActualPackDateTime, order.ActualPackDateTime)
	//require.Equal(t, arg.ActualShipDateTime, order.ActualShipDateTime)
	//require.Equal(t, arg.ActualDeliveryDateTime, order.ActualDeliveryDateTime)
	//require.Equal(t, arg.OrderTotal, order.OrderTotal)
	//require.Equal(t, arg.Notes, order.Notes)
	//require.Equal(t, arg.Status, order.Status)
	//require.Equal(t, arg.CustomerName, order.CustomerName)
	require.NotZero(t, order.ID)
	require.NotZero(t, order.CreatedAt)

	return order
}

func TestQueries_TestCreateOrder(t *testing.T) {
	createRandomOrder(t)
}

func TestQueries_TestGetOrder(t *testing.T) {
	order1 := createRandomOrder(t)
	order2, err := testQueries.GetOrder(context.Background(), order1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order1.ID, order2.ID)
	require.Equal(t, order1.CreatedAt, order2.CreatedAt)
	require.Equal(t, order1.UpdatedAt, order2.UpdatedAt)
	require.Equal(t, order1.ScheduledPackDateTime, order2.ScheduledPackDateTime)
	require.Equal(t, order1.ScheduledShipDateTime, order2.ScheduledShipDateTime)
	require.Equal(t, order1.ScheduledDeliveryDateTime, order2.ScheduledDeliveryDateTime)
	require.Equal(t, order1.ActualPackDateTime, order2.ActualPackDateTime)
	require.Equal(t, order1.ActualShipDateTime, order2.ActualShipDateTime)
	require.Equal(t, order1.ActualDeliveryDateTime, order2.ActualDeliveryDateTime)
	require.Equal(t, order1.OrderTotal, order2.OrderTotal)
	require.Equal(t, order1.Notes, order2.Notes)
	require.Equal(t, order1.Status, order2.Status)
	require.Equal(t, order1.CustomerName, order2.CustomerName)
}

// TODO: Add the rest of the order tests
