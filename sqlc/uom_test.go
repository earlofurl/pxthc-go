package sqlc

import (
	"context"
	"github.com/earlofurl/pxthc/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomUom(t *testing.T) Uom {
	uomName := util.RandomString(6)
	uomAbbreviation := util.RandomString(3)

	arg := CreateUomParams{
		Name:         uomName,
		Abbreviation: uomAbbreviation,
	}

	uom, err := testQueries.CreateUom(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, uom)

	require.Equal(t, uomName, uom.Name)
	require.Equal(t, uomAbbreviation, uom.Abbreviation)

	require.NotZero(t, uom.ID)
	require.NotZero(t, uom.CreatedAt)

	return uom
}

func TestQueries_TestCreateUom(t *testing.T) {
	createRandomUom(t)
}

func TestQueries_TestGetUom(t *testing.T) {
	uom1 := createRandomUom(t)
	uom2, err := testQueries.GetUom(context.Background(), uom1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, uom2)

	require.Equal(t, uom1.ID, uom2.ID)
	require.Equal(t, uom1.Name, uom2.Name)
	require.Equal(t, uom1.Abbreviation, uom2.Abbreviation)
}

func TestQueries_TestListUoms(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUom(t)
	}

	uoms, err := testQueries.ListUoms(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, uoms)

	for _, uom := range uoms {
		require.NotEmpty(t, uom)
	}
}
