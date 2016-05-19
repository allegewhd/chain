package core

import (
	"golang.org/x/net/context"

	"chain/core/explorer"
	"chain/cos/bc"
	"chain/errors"
	"chain/net/http/httpjson"
)

func (a *api) getBlockSummary(ctx context.Context, hash string) (*explorer.BlockSummary, error) {
	return explorer.GetBlockSummary(ctx, a.store, hash)
}

func (a *api) getTx(ctx context.Context, txHashStr string) (*explorer.Tx, error) {
	return explorer.GetTx(ctx, a.store, txHashStr)
}

func (a *api) getAsset(ctx context.Context, assetID string) (*explorer.Asset, error) {
	return explorer.GetAsset(ctx, a.store, assetID)
}

func (a *api) listBlocks(ctx context.Context) (interface{}, error) {
	prev, limit, err := getPageData(ctx, 50)
	if err != nil {
		return nil, err
	}

	list, last, err := explorer.ListBlocks(ctx, a.store, prev, limit)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"blocks": httpjson.Array(list),
		"last":   last,
	}, nil
}

// EXPERIMENTAL(jeffomatic), implemented for R3 demo. Before baking it into the
// public API, we should decide whether this style of API querying is desirable.
func (a *api) getExplorerAssets(ctx context.Context, req struct {
	AssetIDs []string `json:"asset_ids"`
}) (interface{}, error) {
	assets, err := explorer.GetAssets(ctx, a.store, req.AssetIDs)
	if err != nil {
		return nil, err
	}

	var res []*explorer.Asset
	for _, a := range assets {
		res = append(res, a)
	}

	return res, nil
}

func (a *api) listExplorerUTXOsByAsset(ctx context.Context, assetID string) (interface{}, error) {
	prev, limit, err := getPageData(ctx, 50)
	if err != nil {
		return nil, err
	}

	h, err := bc.ParseHash(assetID)
	if err != nil {
		return nil, errors.WithDetailf(httpjson.ErrBadRequest, "invalid asset ID: %q", assetID)
	}

	list, last, err := explorer.ListUTXOsByAsset(ctx, a.store, bc.AssetID(h), prev, limit)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"utxos": httpjson.Array(list),
		"last":  last,
	}, nil
}