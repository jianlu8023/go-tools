package chaincode

import (
	chainmakercommon "chainmaker.org/chainmaker/pb-go/v2/common"
	chainmakersdk "chainmaker.org/chainmaker/sdk-go/v2"
)

// CommonUp 通用调用合约
// @param sdk: 链客户端
// @param contractName: 合约名称
// @param method: 合约方法
// @param txId: 交易ID
// @param kvs: 合约参数
// @param timeout: 超时时间
// @param withSyncResult: 是否同步返回结果
// @return: *chainmakercommon.TxResponse
// @return: error
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

// CommonQuery 通用查询合约
// @param sdk: 链客户端
// @param contractName: 合约名称
// @param method: 合约方法
// @param kvs: 合约参数
// @param timeout: 超时时间
// @return: *chainmakercommon.TxResponse
// @return: error
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

// CommonBlockInfo 通用查询区块信息
// @param sdk: 链客户端
// @param txId: 交易ID
// @param rwSet: 是否返回读写集
// @return: *chainmakercommon.BlockInfo
// @return: error
func CommonBlockInfo(sdk *chainmakersdk.ChainClient,
	txId string, rwSet bool) (*chainmakercommon.BlockInfo, error) {
	blockInfo, err := sdk.GetBlockByTxId(txId, rwSet)
	if err != nil {
		return nil, err
	}
	return blockInfo, nil
}
