package data

import "github.com/Suited-Entertainment/xrpl-go/model/client/common"

type DownloadShardRequest struct {
	Shards []ShardDescriptor `json:"shards"`
}

type ShardDescriptor struct {
	Index common.LedgerIndex `json:"index"`
	URL   string             `json:"url"`
}

func (*DownloadShardRequest) Method() string {
	return "download_shard"
}
