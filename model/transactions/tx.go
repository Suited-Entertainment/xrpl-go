package transactions

import (
	"encoding/json"
	"fmt"

	"github.com/Suited-Entertainment/xrpl-go/model/transactions/types"
)

type Tx interface {
	TxType() TxType
}

type BaseTx struct {
	Account            types.Address
	TransactionType    TxType
	Fee                types.XRPCurrencyAmount `json:",omitempty"`
	Sequence           uint32                  `json:",omitempty"`
	AccountTxnID       types.Hash256           `json:",omitempty"`
	Flags              uint32                  `json:","`
	LastLedgerSequence uint32                  `json:",omitempty"`
	Memos              []MemoWrapper           `json:",omitempty"`
	Signers            []Signer                `json:",omitempty"`
	SourceTag          *uint32                 `json:",omitempty"`
	SigningPubKey      string                  `json:",omitempty"`
	TicketSequence     uint                    `json:",omitempty"`
	TxnSignature       string                  `json:",omitempty"`
	Hash               string                  `json:",omitempty"`
}

func (tx *BaseTx) TxType() TxType {
	return tx.TransactionType
}

type AMMBid struct {
	BaseTx
}

func (*AMMBid) TxType() TxType {
	return AMMBidTx
}

type AMMClawback struct {
	BaseTx
}

func (*AMMClawback) TxType() TxType {
	return AMMClawbackTx
}

type AMMCreate struct {
	BaseTx
}

func (*AMMCreate) TxType() TxType {
	return AMMCreateTx
}

type AMMDelete struct {
	BaseTx
}

func (*AMMDelete) TxType() TxType {
	return AMMDeleteTx
}

type AMMDeposit struct {
	BaseTx
}

func (*AMMDeposit) TxType() TxType {
	return AMMDepositTx
}

type AMMVote struct {
	BaseTx
}

func (*AMMVote) TxType() TxType {
	return AMMVoteTx
}

type AMMWithdraw struct {
	BaseTx
}

func (*AMMWithdraw) TxType() TxType {
	return AMMWithdrawTx
}

type Batch struct {
	BaseTx
}

func (*Batch) TxType() TxType {
	return BatchTx
}

type CredentialAccept struct {
	BaseTx
}

func (*CredentialAccept) TxType() TxType {
	return CredentialAcceptTx
}

type CredentialCreate struct {
	BaseTx
}

func (*CredentialCreate) TxType() TxType {
	return CredentialCreateTx
}

type CredentialDelete struct {
	BaseTx
}

func (*CredentialDelete) TxType() TxType {
	return CredentialDeleteTx
}

type DelegateSet struct {
	BaseTx
}

func (*DelegateSet) TxType() TxType {
	return DelegateSetTx
}

type DIDDelete struct {
	BaseTx
}

func (*DIDDelete) TxType() TxType {
	return DIDDeleteTx
}

type DIDSet struct {
	BaseTx
}

func (*DIDSet) TxType() TxType {
	return DIDSetTx
}

type LedgerStateFix struct {
	BaseTx
}

func (*LedgerStateFix) TxType() TxType {
	return LedgerStateFixTx
}

type MPTokenAuthorize struct {
	BaseTx
}

func (*MPTokenAuthorize) TxType() TxType {
	return MPTokenAuthorizeTx
}

type MPTokenIssuanceCreate struct {
	BaseTx
}

func (*MPTokenIssuanceCreate) TxType() TxType {
	return MPTokenIssuanceCreateTx
}

type MPTokenIssuanceDestroy struct {
	BaseTx
}

func (*MPTokenIssuanceDestroy) TxType() TxType {
	return MPTokenIssuanceDestroyTx
}

type MPTokenIssuanceSet struct {
	BaseTx
}

func (*MPTokenIssuanceSet) TxType() TxType {
	return MPTokenIssuanceSetTx
}

type NFTokenModify struct {
	BaseTx
}

func (*NFTokenModify) TxType() TxType {
	return NFTokenModifyTx
}

type OracleDelete struct {
	BaseTx
}

func (*OracleDelete) TxType() TxType {
	return OracleDeleteTx
}

type OracleSet struct {
	BaseTx
}

func (*OracleSet) TxType() TxType {
	return OracleSetTx
}

type PermissionedDomainDelete struct {
	BaseTx
}

func (*PermissionedDomainDelete) TxType() TxType {
	return PermissionedDomainDeleteTx
}

type PermissionedDomainSet struct {
	BaseTx
}

func (*PermissionedDomainSet) TxType() TxType {
	return PermissionedDomainSetTx
}

type XChainAccountCreateCommit struct {
	BaseTx
}

func (*XChainAccountCreateCommit) TxType() TxType {
	return XChainAccountCreateCommitTx
}

type XChainAddAccountCreateAttestation struct {
	BaseTx
}

func (*XChainAddAccountCreateAttestation) TxType() TxType {
	return XChainAddAccountCreateAttestationTx
}

type XChainAddClaimAttestation struct {
	BaseTx
}

func (*XChainAddClaimAttestation) TxType() TxType {
	return XChainAddClaimAttestationTx
}

type XChainClaim struct {
	BaseTx
}

func (*XChainClaim) TxType() TxType {
	return XChainClaimTx
}

type XChainCommit struct {
	BaseTx
}

func (*XChainCommit) TxType() TxType {
	return XChainCommitTx
}

type XChainCreateBridge struct {
	BaseTx
}

func (*XChainCreateBridge) TxType() TxType {
	return XChainCreateBridgeTx
}

type XChainCreateClaimID struct {
	BaseTx
}

func (*XChainCreateClaimID) TxType() TxType {
	return XChainCreateClaimIDTx
}

type XChainModifyBridge struct {
	BaseTx
}

func (*XChainModifyBridge) TxType() TxType {
	return XChainModifyBridgeTx
}

func UnmarshalTx(data json.RawMessage) (Tx, error) {
	if len(data) == 0 {
		return nil, nil
	}
	if data[0] != '{' {
		return nil, fmt.Errorf("unexpected tx format; must be tx object")
	}
	type txTypeParser struct {
		TransactionType TxType
	}
	var txType txTypeParser
	if err := json.Unmarshal(data, &txType); err != nil {
		return nil, err
	}
	var tx Tx
	switch txType.TransactionType {
	case AccountSetTx:
		tx = &AccountSet{}
	case AccountDeleteTx:
		tx = &AccountDelete{}
	case AMMBidTx:
		tx = &AMMBid{}
	case AMMClawbackTx:
		tx = &AMMClawback{}
	case AMMCreateTx:
		tx = &AMMCreate{}
	case AMMDeleteTx:
		tx = &AMMDelete{}
	case AMMDepositTx:
		tx = &AMMDeposit{}
	case AMMVoteTx:
		tx = &AMMVote{}
	case AMMWithdrawTx:
		tx = &AMMWithdraw{}
	case BatchTx:
		tx = &Batch{}
	case CredentialAcceptTx:
		tx = &CredentialAccept{}
	case CredentialCreateTx:
		tx = &CredentialCreate{}
	case CredentialDeleteTx:
		tx = &CredentialDelete{}
	case CheckCancelTx:
		tx = &CheckCancel{}
	case CheckCashTx:
		tx = &CheckCash{}
	case CheckCreateTx:
		tx = &CheckCreate{}
	case DepositPreauthTx:
		tx = &DepositPreauth{}
	case DelegateSetTx:
		tx = &DelegateSet{}
	case DIDDeleteTx:
		tx = &DIDDelete{}
	case DIDSetTx:
		tx = &DIDSet{}
	case EscrowCancelTx:
		tx = &EscrowCancel{}
	case EscrowCreateTx:
		tx = &EscrowCreate{}
	case EscrowFinishTx:
		tx = &EscrowFinish{}
	case LedgerStateFixTx:
		tx = &LedgerStateFix{}
	case MPTokenAuthorizeTx:
		tx = &MPTokenAuthorize{}
	case MPTokenIssuanceCreateTx:
		tx = &MPTokenIssuanceCreate{}
	case MPTokenIssuanceDestroyTx:
		tx = &MPTokenIssuanceDestroy{}
	case MPTokenIssuanceSetTx:
		tx = &MPTokenIssuanceSet{}
	case NFTokenAcceptOfferTx:
		tx = &NFTokenAcceptOffer{}
	case NFTokenBurnTx:
		tx = &NFTokenBurn{}
	case NFTokenCancelOfferTx:
		tx = &NFTokenCancelOffer{}
	case NFTokenCreateOfferTx:
		tx = &NFTokenCreateOffer{}
	case NFTokenMintTx:
		tx = &NFTokenMint{}
	case NFTokenModifyTx:
		tx = &NFTokenModify{}
	case OfferCreateTx:
		tx = &OfferCreate{}
	case OfferCancelTx:
		tx = &OfferCancel{}
	case OracleDeleteTx:
		tx = &OracleDelete{}
	case OracleSetTx:
		tx = &OracleSet{}
	case PaymentTx:
		tx = &Payment{}
	case PaymentChannelClaimTx:
		tx = &PaymentChannelClaim{}
	case PaymentChannelCreateTx:
		tx = &PaymentChannelCreate{}
	case PaymentChannelFundTx:
		tx = &PaymentChannelFund{}
	case PermissionedDomainDeleteTx:
		tx = &PermissionedDomainDelete{}
	case PermissionedDomainSetTx:
		tx = &PermissionedDomainSet{}
	case SetRegularKeyTx:
		tx = &SetRegularKey{}
	case SignerListSetTx:
		tx = &SignerListSet{}
	case TicketCreateTx:
		tx = &TicketCreate{}
	case TrustSetTx:
		tx = &TrustSet{}
	case XChainAccountCreateCommitTx:
		tx = &XChainAccountCreateCommit{}
	case XChainAddAccountCreateAttestationTx:
		tx = &XChainAddAccountCreateAttestation{}
	case XChainAddClaimAttestationTx:
		tx = &XChainAddClaimAttestation{}
	case XChainClaimTx:
		tx = &XChainClaim{}
	case XChainCommitTx:
		tx = &XChainCommit{}
	case XChainCreateBridgeTx:
		tx = &XChainCreateBridge{}
	case XChainCreateClaimIDTx:
		tx = &XChainCreateClaimID{}
	case XChainModifyBridgeTx:
		tx = &XChainModifyBridge{}
	default:
		fmt.Printf("[xrpl-go] unsupported transaction type %s\n", txType.TransactionType)
		tx = &BaseTx{}
	}
	if err := json.Unmarshal(data, tx); err != nil {
		return nil, err
	}
	return tx, nil
}
