package sqlc

import (
	"context"
	"github.com/earlofurl/pxthc/util"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomPackageAdjEntry(t *testing.T, productPackage Package) PackageAdjEntry {
	arg := CreatePackageAdjEntryParams{
		PackageID: productPackage.ID,
		Amount:    decimal.NewFromFloatWithExponent(util.RandomDecimalTimes100(), -6),
		UomID:     util.RandomInt(1, 6),
	}

	packageAdjEntry, err := testQueries.CreatePackageAdjEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, packageAdjEntry)

	require.Equal(t, arg.PackageID, packageAdjEntry.PackageID)
	require.Equal(t, arg.Amount, packageAdjEntry.Amount)
	require.Equal(t, arg.UomID, packageAdjEntry.UomID)

	require.NotZero(t, packageAdjEntry.ID)
	require.NotZero(t, packageAdjEntry.CreatedAt)

	return packageAdjEntry
}

func TestQueries_TestCreatePackageAdjEntry(t *testing.T) {
	productPackage := createRandomPackage(t)
	createRandomPackageAdjEntry(t, productPackage)
}

func TestQueries_TestGetPackageAdjEntry(t *testing.T) {
	productPackage := createRandomPackage(t)
	packageAdjEntry1 := createRandomPackageAdjEntry(t, productPackage)

	packageAdjEntry2, err := testQueries.GetPackageAdjEntry(context.Background(), packageAdjEntry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, packageAdjEntry2)

	require.Equal(t, packageAdjEntry1.ID, packageAdjEntry2.ID)
	require.Equal(t, packageAdjEntry1.PackageID, packageAdjEntry2.PackageID)
	require.Equal(t, packageAdjEntry1.Amount, packageAdjEntry2.Amount)
	require.Equal(t, packageAdjEntry1.UomID, packageAdjEntry2.UomID)
}

func TestQueries_ListPackageAdjEntries(t *testing.T) {
	productPackage := createRandomPackage(t)

	for i := 0; i < 5; i++ {
		createRandomPackageAdjEntry(t, productPackage)
	}

	arg := ListPackageAdjEntriesParams{
		PackageID: productPackage.ID,
		Limit:     5,
		Offset:    2,
	}

	packageAdjEntries, err := testQueries.ListPackageAdjEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, packageAdjEntries)

	for _, packageAdjEntry := range packageAdjEntries {
		require.NotEmpty(t, packageAdjEntry)
	}
}
