package sqlc

import (
	"context"
	"github.com/earlofurl/pxthc/util"
	"github.com/gobuffalo/nulls"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomItem(t *testing.T) Item {
	itemIsUsed := util.RandomBool()
	itemType1 := createRandomItemType(t)
	strain1 := createRandomStrain(t)

	arg := CreateItemParams{
		Description: util.RandomString(6),
		IsUsed:      itemIsUsed,
		ItemTypeID:  itemType1.ID,
		StrainID:    strain1.ID,
	}

	item, err := testQueries.CreateItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.Description, item.Description)
	require.Equal(t, itemIsUsed, item.IsUsed)
	require.Equal(t, itemType1.ID, item.ItemTypeID)
	require.Equal(t, strain1.ID, item.StrainID)

	require.NotZero(t, item.ID)
	require.NotZero(t, item.ItemTypeID)
	require.NotZero(t, item.StrainID)
	require.NotZero(t, item.CreatedAt)
	require.WithinDuration(t, item.CreatedAt, time.Now(), time.Second)

	return item
}

func TestQueries_CreateItem(t *testing.T) {
	createRandomItem(t)
}

func TestQueries_GetItem(t *testing.T) {
	item1 := createRandomItem(t)
	item2, err := testQueries.GetItem(context.Background(), item1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, item1.Description, item2.Description)
	require.Equal(t, item1.IsUsed, item2.IsUsed)
	require.Equal(t, item1.ItemTypeID, item2.ItemTypeID)
	require.Equal(t, item1.StrainID, item2.StrainID)
}

func TestQueries_ListItem(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomItem(t)
	}

	items, err := testQueries.ListItems(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, items)

	for _, item := range items {
		require.NotEmpty(t, item)
	}
}

func TestQueries_DeleteItem(t *testing.T) {
	item1 := createRandomItem(t)
	err := testQueries.DeleteItem(context.Background(), item1.ID)
	require.NoError(t, err)

	item2, err := testQueries.GetItem(context.Background(), item1.ID)
	require.Error(t, err)
	require.Empty(t, item2)
}

func TestQueries_UpdateItem(t *testing.T) {
	item1 := createRandomItem(t)
	itemType1 := createRandomItemType(t)
	strain1 := createRandomStrain(t)

	arg := UpdateItemParams{
		ID:          item1.ID,
		Description: util.RandomNullsString(10),
		IsUsed:      util.RandomNullsBool(),
		ItemTypeID:  nulls.NewInt64(itemType1.ID),
		StrainID:    nulls.NewInt64(strain1.ID),
	}

	item2, err := testQueries.UpdateItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, arg.Description.String, item2.Description)
	// require.NotEqual(t, arg.IsUsed, item2.IsUsed)
	require.Equal(t, arg.ItemTypeID.Int64, item2.ItemTypeID)
	require.Equal(t, arg.StrainID.Int64, item2.StrainID)
	require.NotEqual(t, item1.Description, item2.Description)
	require.NotEqual(t, item1.ItemTypeID, item2.ItemTypeID)
	require.NotEqual(t, item1.StrainID, item2.StrainID)
	require.WithinDuration(t, item1.CreatedAt, item2.CreatedAt, time.Second)
	require.WithinDuration(t, item2.UpdatedAt, time.Now(), time.Second)
}
