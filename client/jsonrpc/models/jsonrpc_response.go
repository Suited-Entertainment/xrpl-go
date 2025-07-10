package jsonrpcmodels

import (
	"encoding/json"
	"fmt"

	"github.com/Suited-Entertainment/xrpl-go/client"
	"github.com/mitchellh/mapstructure"
)

type JsonRpcResponse struct {
	Result    AnyJson                      `json:"result"`
	Warning   string                       `json:"warning,omitempty"`
	Warnings  []client.XRPLResponseWarning `json:"warnings,omitempty"`
	Forwarded bool                         `json:"forwarded,omitempty"`
}

type AnyJson map[string]interface{}

type ApiWarning struct {
	Id      int         `json:"id"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (r JsonRpcResponse) GetResult(v any) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json",
		Result: &v, DecodeHook: mapstructure.TextUnmarshallerHookFunc()})

	if err != nil {
		return err
	}

	err = dec.Decode(r.Result)
	if err != nil {
		// Use json decoding if mapstructure decoding fails
		data, jsonErr := json.Marshal(r.Result)
		if jsonErr != nil {
			return fmt.Errorf("%w; %w", err, jsonErr)
		}

		jsonErr = json.Unmarshal(data, &v)
		if jsonErr != nil {
			return fmt.Errorf("%w; %w", err, jsonErr)
		}
		return nil
	}
	return nil
}
