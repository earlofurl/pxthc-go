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

func createRandomPackage(t *testing.T) Package {
	var lastRandomPackageTag PackageTag
	// create 10 random package tags
	for i := 0; i < 10; i++ {
		lastRandomPackageTag = createRandomPackageTag(t)
	}

	newRandomUom := createRandomUom(t)
	newRandomItem := createRandomItem(t)

	arg := CreatePackageParams{
		TagID:                             nulls.NewInt64(lastRandomPackageTag.ID),
		PackageType:                       util.RandomString(10),
		IsActive:                          util.RandomBool(),
		Quantity:                          decimal.NewFromFloatWithExponent(util.RandomDecimalTimes100(), -6),
		Notes:                             util.RandomString(10),
		PackagedDateTime:                  time.Now(),
		HarvestDateTime:                   nulls.NewTime(time.Now()),
		LabTestingState:                   util.RandomString(10),
		LabTestingStateDateTime:           nulls.NewTime(time.Now()),
		IsTradeSample:                     util.RandomBool(),
		IsTestingSample:                   util.RandomBool(),
		ProductRequiresRemediation:        util.RandomBool(),
		ContainsRemediatedProduct:         util.RandomBool(),
		RemediationDateTime:               nulls.NewTime(time.Now()),
		ReceivedDateTime:                  nulls.NewTime(time.Now()),
		ReceivedFromManifestNumber:        nulls.NewString(util.RandomString(10)),
		ReceivedFromFacilityLicenseNumber: nulls.NewString(util.RandomString(10)),
		ReceivedFromFacilityName:          nulls.NewString(util.RandomString(10)),
		IsOnHold:                          util.RandomBool(),
		ArchivedDate:                      nulls.NewTime(time.Now()),
		FinishedDate:                      nulls.NewTime(time.Now()),
		ItemID:                            nulls.NewInt64(newRandomItem.ID),
		ProvisionalLabel:                  nulls.NewString(util.RandomString(10)),
		IsProvisional:                     util.RandomBool(),
		IsSold:                            util.RandomBool(),
		PpuDefault:                        decimal.NewFromFloatWithExponent(util.RandomPercent(), -4),
		PpuOnOrder:                        decimal.NewFromFloatWithExponent(util.RandomPercent(), -4),
		TotalPackagePriceOnOrder:          decimal.NewFromFloatWithExponent(util.RandomPercent(), -4),
		PpuSoldPrice:                      decimal.NewFromFloatWithExponent(util.RandomPercent(), -4),
		TotalSoldPrice:                    decimal.NewFromFloatWithExponent(util.RandomPercent(), -4),
		PackagingSuppliesConsumed:         util.RandomBool(),
		IsLineItem:                        util.RandomBool(),
		FacilityLocationID:                1,
		//OrderID:                           nulls.NewInt64(0), // TODO: add back after Order model is created
		UomID: newRandomUom.ID,
		// TODO: add Lab Testing connection in random package creation
	}
	pckg, err := testQueries.CreatePackage(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, pckg)

	// TODO: fix disabled require tests below

	require.Equal(t, arg.TagID, pckg.TagID)
	require.Equal(t, arg.PackageType, pckg.PackageType)
	require.Equal(t, arg.Quantity, pckg.Quantity)
	require.Equal(t, arg.Notes, pckg.Notes)
	//require.Equal(t, arg.PackagedDateTime, pckg.PackagedDateTime)
	//require.Equal(t, arg.HarvestDateTime, pckg.HarvestDateTime)
	require.Equal(t, arg.LabTestingState, pckg.LabTestingState)
	//require.Equal(t, arg.LabTestingStateDateTime, pckg.LabTestingStateDateTime)
	require.Equal(t, arg.IsTradeSample, pckg.IsTradeSample)
	require.Equal(t, arg.IsTestingSample, pckg.IsTestingSample)
	require.Equal(t, arg.ProductRequiresRemediation, pckg.ProductRequiresRemediation)
	require.Equal(t, arg.ContainsRemediatedProduct, pckg.ContainsRemediatedProduct)
	//require.Equal(t, arg.RemediationDateTime, pckg.RemediationDateTime)
	//require.Equal(t, arg.ReceivedDateTime, pckg.ReceivedDateTime)
	require.Equal(t, arg.ReceivedFromManifestNumber, pckg.ReceivedFromManifestNumber)
	require.Equal(t, arg.ReceivedFromFacilityLicenseNumber, pckg.ReceivedFromFacilityLicenseNumber)
	require.Equal(t, arg.ReceivedFromFacilityName, pckg.ReceivedFromFacilityName)
	require.Equal(t, arg.IsOnHold, pckg.IsOnHold)
	//require.Equal(t, arg.ArchivedDate, pckg.ArchivedDate)
	//require.Equal(t, arg.FinishedDate, pckg.FinishedDate)
	require.Equal(t, arg.ItemID, pckg.ItemID)
	require.Equal(t, arg.ProvisionalLabel, pckg.ProvisionalLabel)
	require.Equal(t, arg.IsProvisional, pckg.IsProvisional)
	require.Equal(t, arg.IsSold, pckg.IsSold)
	require.Equal(t, arg.PpuDefault, pckg.PpuDefault)
	require.Equal(t, arg.PpuOnOrder, pckg.PpuOnOrder)
	require.Equal(t, arg.TotalPackagePriceOnOrder, pckg.TotalPackagePriceOnOrder)
	require.Equal(t, arg.PpuSoldPrice, pckg.PpuSoldPrice)
	require.Equal(t, arg.TotalSoldPrice, pckg.TotalSoldPrice)
	require.Equal(t, arg.PackagingSuppliesConsumed, pckg.PackagingSuppliesConsumed)
	require.Equal(t, arg.IsLineItem, pckg.IsLineItem)
	//require.Equal(t, arg.OrderID, pckg.OrderID)
	require.Equal(t, arg.UomID, pckg.UomID)

	return pckg
}

func TestQueries_TestCreatePackage(t *testing.T) {
	createRandomPackage(t)
}

func TestQueries_TestGetPackage(t *testing.T) {
	pckg1 := createRandomPackage(t)

	pckg2, err := testQueries.GetPackage(context.Background(), pckg1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, pckg2)

	require.Equal(t, pckg1.ID, pckg2.ID)
	require.Equal(t, pckg1.TagID, pckg2.TagID)
	require.Equal(t, pckg1.PackageType, pckg2.PackageType)
	require.Equal(t, pckg1.Quantity, pckg2.Quantity)
	require.Equal(t, pckg1.Notes, pckg2.Notes)
	require.Equal(t, pckg1.PackagedDateTime, pckg2.PackagedDateTime)
	require.Equal(t, pckg1.HarvestDateTime, pckg2.HarvestDateTime)
	require.Equal(t, pckg1.LabTestingState, pckg2.LabTestingState)
	require.Equal(t, pckg1.LabTestingStateDateTime, pckg2.LabTestingStateDateTime)
	require.Equal(t, pckg1.IsTradeSample, pckg2.IsTradeSample)
	require.Equal(t, pckg1.IsTestingSample, pckg2.IsTestingSample)
	require.Equal(t, pckg1.ProductRequiresRemediation, pckg2.ProductRequiresRemediation)
	require.Equal(t, pckg1.ContainsRemediatedProduct, pckg2.ContainsRemediatedProduct)
	require.Equal(t, pckg1.RemediationDateTime, pckg2.RemediationDateTime)
	require.Equal(t, pckg1.ReceivedDateTime, pckg2.ReceivedDateTime)
	require.Equal(t, pckg1.ReceivedFromManifestNumber, pckg2.ReceivedFromManifestNumber)
	require.Equal(t, pckg1.ReceivedFromFacilityLicenseNumber, pckg2.ReceivedFromFacilityLicenseNumber)
	require.Equal(t, pckg1.ReceivedFromFacilityName, pckg2.ReceivedFromFacilityName)
	require.Equal(t, pckg1.IsOnHold, pckg2.IsOnHold)
	require.Equal(t, pckg1.ArchivedDate, pckg2.ArchivedDate)
	require.Equal(t, pckg1.FinishedDate, pckg2.FinishedDate)
	require.Equal(t, pckg1.ItemID, pckg2.ItemID)
	require.Equal(t, pckg1.ProvisionalLabel, pckg2.ProvisionalLabel)
	require.Equal(t, pckg1.IsProvisional, pckg2.IsProvisional)
	require.Equal(t, pckg1.IsSold, pckg2.IsSold)
	require.Equal(t, pckg1.PpuDefault, pckg2.PpuDefault)
	require.Equal(t, pckg1.PpuOnOrder, pckg2.PpuOnOrder)
	require.Equal(t, pckg1.TotalPackagePriceOnOrder, pckg2.TotalPackagePriceOnOrder)
	require.Equal(t, pckg1.PpuSoldPrice, pckg2.PpuSoldPrice)
	require.Equal(t, pckg1.TotalSoldPrice, pckg2.TotalSoldPrice)
	require.Equal(t, pckg1.PackagingSuppliesConsumed, pckg2.PackagingSuppliesConsumed)
	require.Equal(t, pckg1.IsLineItem, pckg2.IsLineItem)
	require.Equal(t, pckg1.OrderID, pckg2.OrderID)
	require.Equal(t, pckg1.UomID, pckg2.UomID)

}

//func TestQueries_ListPackagesWithTestResults(t *testing.T) {
//	var lastPackage Package
//	for i := 0; i < 10; i++ {
//		lastPackage = createRandomPackage(t)
//	}
//
//	var lastTestResult TestResult
//	for i := 0; i < 10; i++ {
//		lastTestResult = createRandomTestResult(t)
//	}
//
//
//	pckg2, err := testQueries.GetPackageWithTestResults(context.Background())
//	require.NoError(t, err)
//	require.NotEmpty(t, pckg2)
//
//	require.Equal(t, lastPackage.ID, pckg2.ID)
//	require.Equal(t, lastPackage.TagID, pckg2.TagID)
//	require.Equal(t, lastPackage.PackageType, pckg2.PackageType)
//	require.Equal(t, lastPackage.Quantity, pckg2.Quantity)
//	require.Equal(t, lastPackage.Notes, pckg2.Notes)
//	require.Equal(t, lastPackage.PackagedDateTime, pckg2.PackagedDateTime)
//	require.Equal(t, lastPackage.HarvestDateTime, pckg2.HarvestDateTime)
//	require.Equal(t, lastPackage.LabTestingState, pckg2.LabTestingState)
//	require.Equal(t, lastPackage.LabTestingStateDateTime, pckg2.LabTestingStateDateTime)
//	require.Equal(t, lastPackage.IsTradeSample, pckg2.IsTradeSample)
//	require.Equal(t, lastPackage.IsTestingSample, pckg2.IsTestingSample)
//	require.Equal(t, lastPackage.ProductRequiresRemediation, pckg2.ProductRequiresRemediation)
//	require.Equal(t, lastPackage.ContainsRemediatedProduct, pckg2.ContainsRemediatedProduct)
//	require.Equal(t, lastPackage.RemediationDateTime, pckg2.RemediationDateTime)
//	require.Equal(t, lastPackage.ReceivedDateTime, pckg2.ReceivedDateTime)
//	require.Equal(t, lastPackage.ReceivedFromManifestNumber, pckg2.ReceivedFromManifestNumber)
//	require.Equal(t, lastPackage.ReceivedFromFacilityLicenseNumber, pckg2.ReceivedFromFacilityLicenseNumber)
//	require.Equal(t, lastPackage.ReceivedFromFacilityName, pckg2.ReceivedFromFacilityName)
//	require.Equal(t, lastPackage.IsOnHold, pckg2.IsOnHold)
//	require.Equal(t, lastPackage.ArchivedDate, pckg2.ArchivedDate)
//	require.Equal(t, lastPackage.FinishedDate, pckg2.FinishedDate)
//	require.Equal(t, lastPackage.ItemID, pckg2.ItemID)
//	require.Equal(t, lastPackage.ProvisionalLabel, pckg2.ProvisionalLabel)
//	require.Equal(t, lastPackage.IsProvisional, pckg2.IsProvisional)
//	require.Equal(t, lastPackage.IsSold, pckg2.IsSold)
//	require.Equal(t, lastPackage.PpuDefault, pckg2.PpuDefault)
//	require.Equal(t, lastPackage.PpuOnOrder, pckg2.PpuOnOrder)
//	require.Equal(t, lastPackage.TotalPackagePriceOnOrder, pckg2.TotalPackagePriceOnOrder)
//	require.Equal(t, lastPackage.PpuSoldPrice, pckg2.PpuSoldPrice)
//	require.Equal(t, lastPackage.TotalSoldPrice, pckg2.TotalSoldPrice)
//	require.Equal(t, lastPackage.PackagingSuppliesConsumed, pckg2.PackagingSuppliesConsumed)
//	require.Equal(t, lastPackage.IsLineItem, pckg2.IsLineItem)
//	require.Equal(t, lastPackage.OrderID, pckg2.OrderID)
//	require.Equal(t, lastPackage.UomID, pckg2.UomID)
//
//}

// TODO: Add the rest of the package tests
