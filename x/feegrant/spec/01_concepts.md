<!--
order: 1
-->

# Concepts

## FeeAllowanceGrant

`FeeAllowanceGrant` is stored in the KVStore to record a grant with full context. Every grant will contain `granter`, `grantee` and what kind of `allowance` is granted. `granter` is an account address who is giving permission to `grantee` (the beneficiary account address) to pay for some or all of `grantee`'s transaction fees. `allowance` defines what kind of fee allowance (`BasicFeeAllowance` or `PeriodicFeeAllowance`, see below) is granted to `grantee`. `allowance` accepts an interface which implements `FeeAllowanceI`, encoded as `Any` type. There can be only one existing fee grant allowed for a `grantee` and `granter`, self grants are not allowed.

+++ https://github.com/cosmos/cosmos-sdk/blob/d97e7907f176777ed8a464006d360bb3e1a223e4/proto/cosmos/feegrant/v1beta1/feegrant.proto#L75-L81

`FeeAllowanceI` looks like: 

+++ https://github.com/cosmos/cosmos-sdk/blob/d97e7907f176777ed8a464006d360bb3e1a223e4/x/feegrant/types/fees.go#L9-L32

## Fee Allowance types
There are two types of fee allowances present at the moment:
- `BasicFeeAllowance`
- `PeriodicFeeAllowance`

## BasicFeeAllowance

`BasicFeeAllowance` is permission for `grantee` to use fee from a `granter`'s account. If any of the `spend_limit` or `expiration` reaches its limit, the grant will be removed from the state.
 

+++ https://github.com/cosmos/cosmos-sdk/blob/d97e7907f176777ed8a464006d360bb3e1a223e4/proto/cosmos/feegrant/v1beta1/feegrant.proto#L13-L26

- `spend_limit` is the limit of coins that are allowed to be used from the `granter` account. If it is empty, it assumes there's no spend limit, `grantee` can use any number of available tokens from `granter` account address before the expiration.

- `expiration` specifies an optional time when this allowance expires. If the value is left empty, there is no expiry for the grant.

- When a grant is created with empty values for `spend_limit` and `expiration`, it is still a valid grant. It won't restrict the `grantee` to use any number of tokens from `granter` and it won't have any expiration. The only way to restrict the `grantee` is by revoking the grant.

## PeriodicFeeAllowance

`PeriodicFeeAllowance` is a repeating fee allowance for the mentioned period, we can mention when the grant can expire as well as when a period can reset. We can also define the maximum number of coins that can be used in a mentioned period of time.

+++ https://github.com/cosmos/cosmos-sdk/blob/d97e7907f176777ed8a464006d360bb3e1a223e4/proto/cosmos/feegrant/v1beta1/feegrant.proto#L28-L73

- `basic` is the instance of `BasicFeeAllowance` which is optional for periodic fee allowance. If empty, the grant will have no `expiration` and no `spend_limit`.

- `period` is the specific period of time or blocks, after each period passes, `period_spend_limit` will be reset. 

- `period_spend_limit` specifies the maximum number of coins that can be spent in the period.

- `period_can_spend` is the number of coins left to be spent before the period_reset time.

- `period_reset` keeps track of when a next period reset should happen.

## FeeAccount flag

`feegrant` module introduces a `FeeAccount` flag for CLI for the sake of executing transactions with fee granter. When this flag is set, `clientCtx` will append the granter account address for transactions generated through CLI.

+++ https://github.com/cosmos/cosmos-sdk/blob/d97e7907f176777ed8a464006d360bb3e1a223e4/client/cmd.go#L224-L235

+++ https://github.com/cosmos/cosmos-sdk/blob/d97e7907f176777ed8a464006d360bb3e1a223e4/client/tx/tx.go#L120

+++ https://github.com/cosmos/cosmos-sdk/blob/d97e7907f176777ed8a464006d360bb3e1a223e4/x/auth/tx/builder.go#L268-L277

```go 
message Fee {
  repeated cosmos.base.v1beta1.Coin amount = 1;
  uint64 gas_limit = 2;
  string payer = 3;
  string granter = 4;
}
```

Example cmd:
```go
./simd tx gov submit-proposal --title="Test Proposal" --description="My awesome proposal" --type="Text" --from validator-key --fee-account=cosmos1xh44hxt7spr67hqaa7nyx5gnutrz5fraw6grxn --chain-id=testnet --fees="10stake"
```

## DeductGrantedFeeDecorator

`feegrant` module also adds a `DeductGrantedFeeDecorator` ante handler. Whenever a transaction is being executed with `granter` field set, then this ante handler will check whether `payer` and `granter` have proper fee allowance grant in state. If it exists the fees will be deducted from the `granter`'s account address. If the `granter` field isn't set then this ante handler works as normal fee deductor.