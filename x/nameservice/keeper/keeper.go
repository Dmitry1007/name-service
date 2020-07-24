package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/sdk-tutorials/nameservice/x/nameservice/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	CoinKeeper types.BankKeeper
	storeKey   sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc 		   *codec.Codec // The wire codec for binary encoding/decoding.
}

// Sets the entire Whois metadata struct for a name
func (k Keeper) SetWhois(ctx sdk.Context, name string, whois types.Whois) {
	if whois.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(whois))
}





// Default Scaffold of Keeper
// package keeper
//
// import (
// 	"fmt"
//
// 	"github.com/tendermint/tendermint/crypto"
// 	"github.com/tendermint/tendermint/libs/log"
//
// 	"github.com/cosmos/cosmos-sdk/codec"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/Dmitry1007/name-service/x/nameservice/types"
// )
//
// // Keeper of the nameservice store
// type Keeper struct {
// 	storeKey   sdk.StoreKey
// 	cdc        *codec.Codec
// 	paramspace types.ParamSubspace
// }
//
//
// // Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
// type Keeper struct {
// 	storeKey  sdk.StoreKey // Unexposed key to access store from sdk.Context
// 	cdc *codec.Codec // The wire codec for binary encoding/decoding.
// 	CoinKeeper types.BankKeeper
// }
//
//
// // NewKeeper creates a nameservice keeper
// func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramspace types.ParamSubspace) Keeper {
// 	keeper := Keeper{
// 		storeKey:   key,
// 		cdc:        cdc,
// 		paramspace: paramspace.WithKeyTable(types.ParamKeyTable()),
// 	}
// 	return keeper
// }
//
// // Logger returns a module-specific logger.
// func (k Keeper) Logger(ctx sdk.Context) log.Logger {
// 	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
// }
//
// // Get returns the pubkey from the adddress-pubkey relation
// func (k Keeper) Get(ctx sdk.Context, key string) (/* TODO: Fill out this type */, error) {
// 	store := ctx.KVStore(k.storeKey)
// 	var item /* TODO: Fill out this type */
// 	byteKey := []byte(key)
// 	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &item)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return item, nil
// }
//
// func (k Keeper) set(ctx sdk.Context, key string, value /* TODO: fill out this type */ ) {
// 	store := ctx.KVStore(k.storeKey)
// 	bz := k.cdc.MustMarshalBinaryLengthPrefixed(value)
// 	store.Set([]byte(key), bz)
// }
//
// func (k Keeper) delete(ctx sdk.Context, key string) {
// 	store := ctx.KVStore(k.storeKey)
// 	store.Delete([]byte(key))
// }
