package sqlc

import (
	"context"
	"github.com/earlofurl/pxthc/util"
	"github.com/gobuffalo/nulls"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomItemType(t *testing.T) ItemType {
	uom1 := createRandomUom(t)
	productCategory1 := createRandomProductCategory(t)

	productForm := util.RandomString(6)
	productModifier := util.RandomString(6)

	arg := CreateItemTypeParams{
		ProductForm:       productForm,
		ProductModifier:   productModifier,
		UomDefault:        uom1.ID,
		ProductCategoryID: productCategory1.ID,
	}

	itemType, err := testQueries.CreateItemType(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, itemType)

	require.Equal(t, productForm, itemType.ProductForm)
	require.Equal(t, productModifier, itemType.ProductModifier)
	require.Equal(t, uom1.ID, itemType.UomDefault)
	require.Equal(t, productCategory1.ID, itemType.ProductCategoryID)

	require.NotZero(t, itemType.ID)
	require.NotZero(t, itemType.UomDefault)
	require.NotZero(t, itemType.ProductCategoryID)
	require.NotZero(t, itemType.CreatedAt)

	return itemType
}

func TestQueries_CreateItemType(t *testing.T) {
	createRandomItemType(t)
}

func TestQueries_GetItemType(t *testing.T) {
	itemType1 := createRandomItemType(t)
	itemType2, err := testQueries.GetItemType(context.Background(), itemType1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, itemType2)

	require.Equal(t, itemType1.ID, itemType2.ID)
	require.Equal(t, itemType1.ProductForm, itemType2.ProductForm)
	require.Equal(t, itemType1.ProductModifier, itemType2.ProductModifier)
	require.Equal(t, itemType1.UomDefault, itemType2.UomDefault)
	require.Equal(t, itemType1.ProductCategoryID, itemType2.ProductCategoryID)
}

func TestQueries_ListItemTypes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomItemType(t)
	}

	itemTypes, err := testQueries.ListItemTypes(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, itemTypes)

	for _, itemType := range itemTypes {
		require.NotEmpty(t, itemType)
	}
}

func TestQueries_UpdateItemType(t *testing.T) {
	itemType1 := createRandomItemType(t)

	uom2 := createRandomUom(t)
	productCategory2 := createRandomProductCategory(t)

	productForm := util.RandomString(6)
	productModifier := util.RandomString(6)

	arg := UpdateItemTypeParams{
		ID:                itemType1.ID,
		ProductForm:       nulls.NewString(productForm),
		ProductModifier:   nulls.NewString(productModifier),
		UomDefault:        nulls.NewInt64(uom2.ID),
		ProductCategoryID: nulls.NewInt64(productCategory2.ID),
	}

	itemType2, err := testQueries.UpdateItemType(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, itemType2)

	require.Equal(t, itemType1.ID, itemType2.ID)
	require.Equal(t, productForm, itemType2.ProductForm)
	require.Equal(t, productModifier, itemType2.ProductModifier)
	require.Equal(t, uom2.ID, itemType2.UomDefault)
	require.Equal(t, productCategory2.ID, itemType2.ProductCategoryID)
}

func TestQueries_DeleteItemType(t *testing.T) {
	itemType1 := createRandomItemType(t)

	err := testQueries.DeleteItemType(context.Background(), itemType1.ID)
	require.NoError(t, err)

	itemType2, err := testQueries.GetItemType(context.Background(), itemType1.ID)
	require.Error(t, err)
	require.Empty(t, itemType2)
}
