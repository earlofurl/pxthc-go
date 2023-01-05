package sqlc

import (
	"context"
	"github.com/earlofurl/pxthc/util"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomLabTest(t *testing.T) LabTest {
	arg := CreateLabTestParams{
		TestName:                util.RandomString(10),
		BatchCode:               util.RandomString(10),
		TestIDCode:              util.RandomString(10),
		LabFacilityName:         util.RandomString(24),
		TestPerformedDateTime:   time.Now(),
		OverallPassed:           util.RandomBool(),
		TestTypeName:            util.RandomString(10),
		TestPassed:              util.RandomBool(),
		TestComment:             util.RandomString(10),
		ThcTotalPercent:         decimal.NewFromFloatWithExponent(util.RandomPercent(), -6),
		ThcTotalValue:           decimal.NewFromFloatWithExponent(util.RandomDecimalTimes100(), -6),
		CbdPercent:              decimal.NewFromFloatWithExponent(util.RandomPercent(), -6),
		CbdValue:                decimal.NewFromFloatWithExponent(util.RandomDecimalTimes100(), -6),
		TerpeneTotalPercent:     decimal.NewFromFloatWithExponent(util.RandomPercent(), -6),
		TerpeneTotalValue:       decimal.NewFromFloatWithExponent(util.RandomDecimalTimes100(), -6),
		ThcAPercent:             decimal.NewFromFloatWithExponent(util.RandomPercent(), -6),
		ThcAValue:               decimal.NewFromFloatWithExponent(util.RandomDecimalTimes100(), -6),
		Delta9ThcPercent:        decimal.NewFromFloatWithExponent(util.RandomPercent(), -6),
		Delta9ThcValue:          decimal.NewFromFloatWithExponent(util.RandomDecimalTimes100(), -6),
		Delta8ThcPercent:        decimal.NewFromFloatWithExponent(util.RandomPercent(), -6),
		Delta8ThcValue:          decimal.NewFromFloatWithExponent(util.RandomDecimalTimes100(), -6),
		ThcVPercent:             decimal.NewFromFloatWithExponent(util.RandomPercent(), -6),
		ThcVValue:               decimal.NewFromFloatWithExponent(util.RandomDecimalTimes100(), -6),
		CbdAPercent:             decimal.NewFromFloatWithExponent(util.RandomPercent(), -6),
		CbdAValue:               decimal.NewFromFloatWithExponent(util.RandomDecimalTimes100(), -6),
		CbnPercent:              decimal.NewFromFloatWithExponent(util.RandomPercent(), -6),
		CbnValue:                decimal.NewFromFloatWithExponent(util.RandomDecimalTimes100(), -6),
		CbgAPercent:             decimal.NewFromFloatWithExponent(util.RandomPercent(), -6),
		CbgAValue:               decimal.NewFromFloatWithExponent(util.RandomDecimalTimes100(), -6),
		CbgPercent:              decimal.NewFromFloatWithExponent(util.RandomPercent(), -6),
		CbgValue:                decimal.NewFromFloatWithExponent(util.RandomDecimalTimes100(), -6),
		CbcPercent:              decimal.NewFromFloatWithExponent(util.RandomPercent(), -6),
		CbcValue:                decimal.NewFromFloatWithExponent(util.RandomDecimalTimes100(), -6),
		TotalCannabinoidPercent: decimal.NewFromFloatWithExponent(util.RandomPercent(), -6),
		TotalCannabinoidValue:   decimal.NewFromFloatWithExponent(util.RandomDecimalTimes100(), -6),
	}

	labTest, err := testQueries.CreateLabTest(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, labTest)

	require.Equal(t, arg.TestName, labTest.TestName)
	require.Equal(t, arg.BatchCode, labTest.BatchCode)
	require.Equal(t, arg.TestIDCode, labTest.TestIDCode)
	require.Equal(t, arg.LabFacilityName, labTest.LabFacilityName)
	require.Equal(t, arg.OverallPassed, labTest.OverallPassed)
	require.Equal(t, arg.TestTypeName, labTest.TestTypeName)
	require.Equal(t, arg.TestPassed, labTest.TestPassed)
	require.Equal(t, arg.TestComment, labTest.TestComment)
	require.Equal(t, arg.ThcTotalPercent, labTest.ThcTotalPercent)
	require.Equal(t, arg.ThcTotalValue, labTest.ThcTotalValue)
	require.Equal(t, arg.CbdPercent, labTest.CbdPercent)
	require.Equal(t, arg.CbdValue, labTest.CbdValue)
	require.Equal(t, arg.TerpeneTotalPercent, labTest.TerpeneTotalPercent)
	require.Equal(t, arg.TerpeneTotalValue, labTest.TerpeneTotalValue)
	require.Equal(t, arg.ThcAPercent, labTest.ThcAPercent)
	require.Equal(t, arg.ThcAValue, labTest.ThcAValue)
	require.Equal(t, arg.Delta9ThcPercent, labTest.Delta9ThcPercent)
	require.Equal(t, arg.Delta9ThcValue, labTest.Delta9ThcValue)
	require.Equal(t, arg.Delta8ThcPercent, labTest.Delta8ThcPercent)
	require.Equal(t, arg.Delta8ThcValue, labTest.Delta8ThcValue)
	require.Equal(t, arg.ThcVPercent, labTest.ThcVPercent)
	require.Equal(t, arg.ThcVValue, labTest.ThcVValue)
	require.Equal(t, arg.CbdAPercent, labTest.CbdAPercent)
	require.Equal(t, arg.CbdAValue, labTest.CbdAValue)
	require.Equal(t, arg.CbnPercent, labTest.CbnPercent)
	require.Equal(t, arg.CbnValue, labTest.CbnValue)
	require.Equal(t, arg.CbgAPercent, labTest.CbgAPercent)
	require.Equal(t, arg.CbgAValue, labTest.CbgAValue)
	require.Equal(t, arg.CbgPercent, labTest.CbgPercent)
	require.Equal(t, arg.CbgValue, labTest.CbgValue)
	require.Equal(t, arg.CbcPercent, labTest.CbcPercent)
	require.Equal(t, arg.CbcValue, labTest.CbcValue)
	require.Equal(t, arg.TotalCannabinoidPercent, labTest.TotalCannabinoidPercent)
	require.Equal(t, arg.TotalCannabinoidValue, labTest.TotalCannabinoidValue)

	require.NotZero(t, labTest.ID)
	require.NotZero(t, labTest.CreatedAt)
	require.WithinDuration(t, time.Now(), labTest.CreatedAt, time.Second)

	return labTest
}

func TestQueries_CreateLabTest(t *testing.T) {
	createRandomLabTest(t)
}

func TestQueries_AssignLabTestToPackage(t *testing.T) {
	labTest := createRandomLabTest(t)
	productPackage := createRandomPackage(t)

	arg := AssignLabTestToPackageParams{
		LabTestID: labTest.ID,
		PackageID: productPackage.ID,
	}

	err := testQueries.AssignLabTestToPackage(context.Background(), arg)

	// check if the lab test is assigned to the package
	labTest, err = testQueries.GetLabTest(context.Background(), labTest.ID)
	productPackage, err = testQueries.GetPackage(context.Background(), productPackage.ID)

	require.NoError(t, err)
}

// TODO: Add the rest of the lab test tests
