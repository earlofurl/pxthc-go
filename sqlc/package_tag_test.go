package sqlc

import (
	"context"
	"github.com/earlofurl/pxthc/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomPackageTag(t *testing.T) PackageTag {
	arg := CreatePackageTagParams{
		TagNumber:     util.RandomTagNumber(),
		IsAssigned:    util.RandomBool(),
		IsProvisional: util.RandomBool(),
		IsActive:      util.RandomBool(),
		//AssignedPackageID: nulls.NewInt64(newRandomPackage.ID),
	}

	packagetag, err := testQueries.CreatePackageTag(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, packagetag)

	require.Equal(t, arg.TagNumber, packagetag.TagNumber)
	require.Equal(t, arg.IsAssigned, packagetag.IsAssigned)
	require.Equal(t, arg.IsProvisional, packagetag.IsProvisional)
	require.Equal(t, arg.IsActive, packagetag.IsActive)
	require.NotZero(t, packagetag.CreatedAt)

	return packagetag
}

func TestQueries_TestCreatePackageTag(t *testing.T) {
	createRandomPackageTag(t)
}

func TestQueries_Create1000PackageTags(t *testing.T) {
	for i := 0; i < 1000; i++ {
		createRandomPackageTag(t)
	}
}

func TestQueries_TestGetPackageTag(t *testing.T) {
	packagetag1 := createRandomPackageTag(t)
	packagetag2, err := testQueries.GetPackageTag(context.Background(), packagetag1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, packagetag2)

	require.Equal(t, packagetag1.ID, packagetag2.ID)
	require.Equal(t, packagetag1.TagNumber, packagetag2.TagNumber)
	require.Equal(t, packagetag1.IsAssigned, packagetag2.IsAssigned)
	require.Equal(t, packagetag1.IsProvisional, packagetag2.IsProvisional)
	require.Equal(t, packagetag1.IsActive, packagetag2.IsActive)
	require.WithinDuration(t, packagetag1.CreatedAt, packagetag2.CreatedAt, time.Second)
}

func TestQueries_TestListPackageTags(t *testing.T) {
	arg := ListPackageTagsParams{
		IsAssigned: false,
		Limit:      20,
		Offset:     0,
	}

	for i := 0; i < 10; i++ {
		createRandomPackageTag(t)
	}

	packagetags, err := testQueries.ListPackageTags(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, packagetags)

	for _, packagetag := range packagetags {
		require.NotEmpty(t, packagetag)
	}
}

func TestQueries_GetPackageTagByTagNumber(t *testing.T) {
	packagetag1 := createRandomPackageTag(t)
	packagetag2, err := testQueries.GetPackageTagByTagNumber(context.Background(), packagetag1.TagNumber)
	require.NoError(t, err)
	require.NotEmpty(t, packagetag2)

	require.Equal(t, packagetag1.ID, packagetag2.ID)
	require.Equal(t, packagetag1.TagNumber, packagetag2.TagNumber)
	require.Equal(t, packagetag1.IsAssigned, packagetag2.IsAssigned)
	require.Equal(t, packagetag1.IsProvisional, packagetag2.IsProvisional)
	require.Equal(t, packagetag1.IsActive, packagetag2.IsActive)
	require.WithinDuration(t, packagetag1.CreatedAt, packagetag2.CreatedAt, time.Second)

}

func TestQueries_TestUpdatePackageTag(t *testing.T) {
	packagetag1 := createRandomPackageTag(t)

	arg := UpdatePackageTagParams{
		ID:            packagetag1.ID,
		IsAssigned:    util.RandomNullsBool(),
		IsProvisional: util.RandomNullsBool(),
		IsActive:      util.RandomNullsBool(),
	}

	packagetag2, err := testQueries.UpdatePackageTag(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, packagetag2)

	require.Equal(t, packagetag1.ID, packagetag2.ID)
	require.Equal(t, arg.IsAssigned.Bool, packagetag2.IsAssigned)
	require.Equal(t, arg.IsProvisional.Bool, packagetag2.IsProvisional)
	require.Equal(t, arg.IsActive.Bool, packagetag2.IsActive)
	require.Equal(t, packagetag1.AssignedPackageID, packagetag2.AssignedPackageID)
	require.WithinDuration(t, packagetag1.CreatedAt, packagetag2.CreatedAt, time.Second)
}
