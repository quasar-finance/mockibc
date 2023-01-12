package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "mockibc/testutil/keeper"
	"mockibc/x/mockibc/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MockibcKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
