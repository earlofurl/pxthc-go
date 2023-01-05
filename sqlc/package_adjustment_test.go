package sqlc

import (
	"context"
	"github.com/earlofurl/pxthc/util"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomPackageAdjustment(t *testing.T, fromPackageID int64, toPackageID int64) PackageAdjustment {
	arg := CreatePackageAdjustmentParams{
		FromPackageID: fromPackageID,
		ToPackageID:   toPackageID,
		Amount:        decimal.NewFromFloatWithExponent(util.RandomDecimalTimes100(), -6),
		UomID:         util.RandomInt(1, 6),
	}

	packageAdjustment, err := testQueries.CreatePackageAdjustment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, packageAdjustment)

	require.Equal(t, arg.FromPackageID, packageAdjustment.FromPackageID)
	require.Equal(t, arg.ToPackageID, packageAdjustment.ToPackageID)
	require.Equal(t, arg.Amount, packageAdjustment.Amount)
	require.Equal(t, arg.UomID, packageAdjustment.UomID)

	require.NotZero(t, packageAdjustment.ID)
	require.NotZero(t, packageAdjustment.CreatedAt)

	return packageAdjustment
}

func TestQueries_TestCreatePackageAdjustment(t *testing.T) {
	fromPackage := createRandomPackage(t)
	toPackage := createRandomPackage(t)
	createRandomPackageAdjustment(t, fromPackage.ID, toPackage.ID)
}

func TestQueries_TestGetPackageAdjustment(t *testing.T) {
	fromPackage := createRandomPackage(t)
	toPackage := createRandomPackage(t)
	packageAdjustment1 := createRandomPackageAdjustment(t, fromPackage.ID, toPackage.ID)

	packageAdjustment2, err := testQueries.GetPackageAdjustment(context.Background(), packageAdjustment1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, packageAdjustment2)

	require.Equal(t, packageAdjustment1.ID, packageAdjustment2.ID)
	require.Equal(t, packageAdjustment1.FromPackageID, packageAdjustment2.FromPackageID)
	require.Equal(t, packageAdjustment1.ToPackageID, packageAdjustment2.ToPackageID)
	require.Equal(t, packageAdjustment1.Amount, packageAdjustment2.Amount)
	require.Equal(t, packageAdjustment1.UomID, packageAdjustment2.UomID)
}

func TestQueries_TestListPackageAdjustment(t *testing.T) {
	fromPackage := createRandomPackage(t)
	toPackage := createRandomPackage(t)

	for i := 0; i < 5; i++ {
		createRandomPackageAdjustment(t, fromPackage.ID, toPackage.ID)
	}

	arg := ListPackageAdjustmentsParams{
		FromPackageID: fromPackage.ID,
		ToPackageID:   toPackage.ID,
		Limit:         5,
		Offset:        2,
	}

	packageAdjustments, err := testQueries.ListPackageAdjustments(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, packageAdjustments)

	for _, packageAdjustment := range packageAdjustments {
		require.NotEmpty(t, packageAdjustment)
	}
}
