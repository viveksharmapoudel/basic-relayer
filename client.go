package main

import (
	"errors"
	"fmt"

	"github.com/icon-project/goloop/client"
	"github.com/icon-project/goloop/common"
	"github.com/icon-project/goloop/common/codec"
	"github.com/icon-project/goloop/server/jsonrpc"
	v3 "github.com/icon-project/goloop/server/v3"
	"github.com/icon-project/icon-bridge/cmd/iconbridge/chain/icon/types"
	"github.com/icon-project/icon-bridge/common/log"
)

type Client struct {
	*client.ClientV3
}

func (c *Client) GetValidatorsByHash(hash common.HexHash) ([]common.Address, error) {
	data, err := c.GetDataByHash(&v3.DataHashParam{Hash: jsonrpc.HexBytes(common.HexBytes(hash).String())})
	if err != nil {
		return nil, errors.New(fmt.Sprintf(err.Error(), "GetDataByHash; %v", err))
	}

	var validators []common.Address
	_, err = codec.BC.UnmarshalFromBytes(data, &validators)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(err.Error(), "Unmarshal Validators: %v", err))
	}
	return validators, nil
}

func NewClient(endpoint string) *Client {
	return &Client{
		client.NewClientV3(endpoint),
	}
}

func (cl *Client) btpSource() {
	op, err := cl.ClientV3.GetBTPSourceInformation()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(op)
}

func (cl *Client) GetBtpHeader(height, networkID types.HexInt) (*BtpBlockHeaderFormat, error) {
	headerEncoded, err := cl.ClientV3.GetBTPHeader(&v3.BTPMessagesParam{
		Height:    jsonrpc.HexInt(height),
		NetworkId: jsonrpc.HexInt(networkID),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var header BtpBlockHeaderFormat
	_, err = Base64ToData(headerEncoded, &header)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &header, nil
}

func (cl *Client) GetBTPProof(height, networkID types.HexInt) (*Secp256k1Proof, error) {
	// btp_proof
	proofsEncoded, err := cl.ClientV3.GetBTPProof(&v3.BTPMessagesParam{
		Height:    jsonrpc.HexInt(height),
		NetworkId: jsonrpc.HexInt(networkID),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var btpproof Secp256k1Proof
	_, err = Base64ToData(proofsEncoded, &btpproof)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &btpproof, nil

}

func (cl *Client) GetValidatorSet(height types.HexInt) ([]common.Address, error) {

	blockHeaderBytes, err := cl.ClientV3.GetBlockHeaderByHeight(&v3.BlockHeightParam{
		Height: jsonrpc.HexInt(height),
	})

	var blockHeader types.BlockHeader
	_, err = codec.BC.UnmarshalFromBytes(blockHeaderBytes, &blockHeader)
	if err != nil {
		log.Println(err)
		return []common.Address{}, err
	}

	validatorsList, err := cl.GetValidatorsByHash(blockHeader.NextValidatorsHash)
	if err != nil {
		log.Println(err)
		return []common.Address{}, err
	}

	return validatorsList, nil
}
