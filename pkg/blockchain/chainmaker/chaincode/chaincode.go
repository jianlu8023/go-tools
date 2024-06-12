package chaincode

import (
	chainmakercommon "chainmaker.org/chainmaker/pb-go/v2/common"
	chainmakersdk "chainmaker.org/chainmaker/sdk-go/v2"
)

func CommonUp(sdk *chainmakersdk.ChainClient,
	contractName, method, txId string,
	kvs []*chainmakercommon.KeyValuePair,
	timeout int64, withSyncResult bool) (*chainmakercommon.TxResponse, error) {
	txResponse, err := sdk.InvokeContract(contractName, method, txId, kvs, timeout, withSyncResult)
	if err != nil {
		return nil, err
	}
	return txResponse, nil
}

func CommonQuery(sdk *chainmakersdk.ChainClient,
	contractName, method string,
	kvs []*chainmakercommon.KeyValuePair,
	timeout int64) (*chainmakercommon.TxResponse, error) {

	txResponse, err := sdk.QueryContract(contractName, method, kvs, timeout)

	if err != nil {
		return nil, err
	}
	return txResponse, nil
}

func CommonBlockInfo(sdk *chainmakersdk.ChainClient, txId string, rwSet bool) (*chainmakercommon.BlockInfo, error) {
	blockInfo, err := sdk.GetBlockByTxId(txId, rwSet)
	if err != nil {
		return nil, err
	}
	return blockInfo, nil
}
