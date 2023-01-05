package sqlc

import (
	"context"
	"fmt"
	"github.com/earlofurl/pxthc/util"
	"github.com/gobuffalo/nulls"
	"github.com/shopspring/decimal"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQueries_PackageToPackageTx(t *testing.T) {
	store := NewStore(testDB)

	package1 := createRandomPackage(t)
	package2 := createRandomPackage(t)
	labTest1 := createRandomLabTest(t)
	fmt.Println(">> before:", "Pckg 1:", package1.Quantity, "Pckg 2:", package2.Quantity)

	n := 5
	amount := decimal.NewFromFloatWithExponent(10.0, -6)

	errs := make(chan error)
	results := make(chan CreatePckgToPckgTxResult)

	// run n concurrent package to package transfer transaction
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.CreatePckgToPckgTx(context.Background(), CreatePckgToPckgTxParams{
				FromPackageID: package1.ID,
				ToPackageID:   package2.ID,
				Amount:        amount,
				UomID:         package1.UomID,
				LabTestID:     labTest1.ID.Int64,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check package adjustment
		pckgAdjustment := result.PackageAdjustment
		require.NotEmpty(t, pckgAdjustment)
		require.Equal(t, package1.ID, pckgAdjustment.FromPackageID)
		require.Equal(t, package2.ID, pckgAdjustment.ToPackageID)
		// TODO: Fix the decimal types here.
		//require.Equal(t, amount, pckgAdjustment.Amount)
		require.NotZero(t, pckgAdjustment.ID)
		require.NotZero(t, pckgAdjustment.CreatedAt)

		_, err = store.GetPackageAdjustment(context.Background(), pckgAdjustment.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := result.FromPackageAdjEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, package1.ID, fromEntry.PackageID)
		//require.Equal(t, amount.Neg(), fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetPackageAdjEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToPackageAdjEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, package2.ID, toEntry.PackageID)
		//require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetPackageAdjEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check packages
		fromPackage := result.FromPackage
		require.NotEmpty(t, fromPackage)
		require.Equal(t, package1.ID, fromPackage.ID)

		toPackage := result.ToPackage
		require.NotEmpty(t, toPackage)
		require.Equal(t, package2.ID, toPackage.ID)

		// check the diff of balance
		diff1 := package1.Quantity.Sub(fromPackage.Quantity)
		diff2 := toPackage.Quantity.Sub(package2.Quantity)
		require.Equal(t, diff1, diff2)
		//require.Equal(t, amount, diff1)

		k := diff1.Div(amount)
		require.True(t, k.GreaterThanOrEqual(decimal.NewFromInt(1)) && k.LessThanOrEqual(decimal.NewFromInt(int64(n))))
		require.NotContains(t, existed, k)
		existed[int(k.IntPart())] = true
	}

	// check the final updated balance
	updatedPackage1, err := store.GetPackage(context.Background(), package1.ID)
	require.NoError(t, err)

	updatedPackage2, err := store.GetPackage(context.Background(), package2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedPackage1.Quantity, updatedPackage2.Quantity)
	require.Equal(t, package1.Quantity.Sub(amount.Mul(decimal.NewFromInt(int64(n)))), updatedPackage1.Quantity)
	require.Equal(t, package2.Quantity.Add(amount.Mul(decimal.NewFromInt(int64(n)))), updatedPackage2.Quantity)
}

func TestQueries_CreatePackageTx(t *testing.T) {
	store := NewStore(testDB)

	n := 5
	amount := decimal.NewFromFloatWithExponent(util.RandomFloat(1, 20), -6)

	errs := make(chan error)
	results := make(chan CreatePackageTxResult)

	// run n concurrent package creation transactions
	for i := 0; i < n; i++ {
		i := i
		go func() {
			randomSourcePackage := createRandomPackage(t)

			fmt.Println("Source Package ID:", randomSourcePackage.ID)
			fmt.Println("Source Package Qty >> before:", randomSourcePackage.Quantity)

			createPackageParams := CreatePackageParams{
				TagID:                             nulls.NewInt64(int64(i) + util.RandomInt(1, 1000)),
				PackageType:                       "test",
				IsActive:                          true,
				Quantity:                          amount,
				Notes:                             "test",
				PackagedDateTime:                  time.Now(),
				HarvestDateTime:                   randomSourcePackage.HarvestDateTime,
				LabTestingState:                   "test",
				LabTestingStateDateTime:           nulls.NewTime(time.Now()),
				IsTradeSample:                     false,
				IsTestingSample:                   false,
				ProductRequiresRemediation:        false,
				ContainsRemediatedProduct:         false,
				RemediationDateTime:               nulls.NewTime(time.Now()),
				ReceivedDateTime:                  nulls.NewTime(time.Now()),
				ReceivedFromManifestNumber:        nulls.NewString("test"),
				ReceivedFromFacilityLicenseNumber: nulls.NewString("test"),
				ReceivedFromFacilityName:          nulls.NewString("test"),
				IsOnHold:                          false,
				ArchivedDate:                      nulls.NewTime(time.Now()),
				FinishedDate:                      nulls.NewTime(time.Now()),
				ItemID:                            nulls.NewInt64(util.RandomInt(1, 5)),
				ProvisionalLabel:                  nulls.NewString("test"),
				IsProvisional:                     false,
				IsSold:                            false,
				PpuDefault:                        decimal.NewFromFloat(1.0),
				PpuOnOrder:                        decimal.NewFromFloat(1.0),
				TotalPackagePriceOnOrder:          decimal.NewFromFloat(0.0),
				PpuSoldPrice:                      decimal.NewFromFloat(0.0),
				TotalSoldPrice:                    decimal.NewFromFloat(0.0),
				PackagingSuppliesConsumed:         false,
				IsLineItem:                        false,
				OrderID:                           nulls.NewInt64(1),
				UomID:                             1,
				FacilityLocationID:                1,
			}

			result, err := store.CreatePackageTx(context.Background(), CreatePackageTxParams{
				SourcePackageID:     nulls.NewInt64(randomSourcePackage.ID),
				CreatePackageParams: createPackageParams,
			})

			results <- result
			errs <- err
		}()
	}

	// check the results
	//existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		result := <-results
		err := <-errs
		require.NoError(t, err)
		require.NotEmpty(t, result)

		fmt.Println("Created Package:", result.CreatedPackage)
		fmt.Println("Created Package Adjustment:", result.PackageAdjustment)
		fmt.Println("From Package Adj Entry:", result.FromPackageAdjEntry)
		fmt.Println("To Package Adj Entry:", result.ToPackageAdjEntry)
		fmt.Println("Source Package Child Package Entry:", result.SourcePackageChildPackageEntry)

		// check the diff of balance
		//diff1 := result.CreatedPackage.Quantity.Sub(package1.Quantity)
		//require.Equal(t, amount, diff1)
		//k := diff1.Div(amount)
		//require.True(t, k.GreaterThanOrEqual(decimal.NewFromInt(1)) && k.LessThanOrEqual(decimal.NewFromInt(int64(n))))
		//require.NotContains(t, existed, k)
		//existed[int(k.IntPart())] = true
	}
}
