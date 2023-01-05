package sqlc

import (
	"context"
	"database/sql"
	"github.com/earlofurl/pxthc/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomProductCategory(t *testing.T) ProductCategory {
	productCategoryName := util.RandomCategory()

	productCategory, err := testQueries.CreateProductCategory(context.Background(), productCategoryName)
	require.NoError(t, err)
	require.NotEmpty(t, productCategory)

	require.Equal(t, productCategoryName, productCategory.Name)

	require.NotZero(t, productCategory.ID)
	require.NotZero(t, productCategory.CreatedAt)

	return productCategory
}

func TestQueries_TestCreateProductCategory(t *testing.T) {
	createRandomProductCategory(t)
}

func TestQueries_TestGetProductCategory(t *testing.T) {
	productCategory1 := createRandomProductCategory(t)
	productCategory2, err := testQueries.GetProductCategory(context.Background(), productCategory1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, productCategory2)

	require.Equal(t, productCategory1.ID, productCategory2.ID)
	require.Equal(t, productCategory1.Name, productCategory2.Name)
	require.WithinDuration(t, productCategory1.CreatedAt, productCategory2.CreatedAt, time.Second)
}

func TestQueries_TestUpdateProductCategory(t *testing.T) {
	productCategory1 := createRandomProductCategory(t)

	arg := UpdateProductCategoryParams{
		ID:   productCategory1.ID,
		Name: util.RandomCategory(),
	}

	productCategory2, err := testQueries.UpdateProductCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, productCategory2)

	require.Equal(t, productCategory1.ID, productCategory2.ID)
	require.Equal(t, arg.Name, productCategory2.Name)
	require.WithinDuration(t, productCategory1.CreatedAt, productCategory2.CreatedAt, time.Second)
	require.WithinDuration(t, productCategory2.UpdatedAt, time.Now(), time.Second)
}

func TestQueries_TestDeleteProductCategory(t *testing.T) {
	productCategory1 := createRandomProductCategory(t)

	err := testQueries.DeleteProductCategory(context.Background(), productCategory1.ID)
	require.NoError(t, err)

	productCategory2, err := testQueries.GetProductCategory(context.Background(), productCategory1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, productCategory2)
}

func TestQueries_TestListProductCategories(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProductCategory(t)
	}

	productCategories, err := testQueries.ListProductCategories(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, productCategories)

}
