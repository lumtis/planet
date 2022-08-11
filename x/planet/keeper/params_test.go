package keeper_test

import (
	"testing"

	testkeeper "github.com/lubtd/planet/testutil/keeper"
	"github.com/lubtd/planet/x/planet/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PlanetKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
