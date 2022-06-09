package keeper

import (
	"encoding/binary"

	"github.com/nuonchanndarong/blog/x/blog/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	// Get the current number of posts in the store
	count := k.GetPostCount(ctx)
	// Assign post.Id = count
	post.Id = count
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostKey))
	// Convert post id into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, post.Id)
	// Marshal the post into bytes
	appendedValue := k.cdc.MustMarshal(&post)
	// Insert
	store.Set(byteKey, appendedValue)
	//update the post count
	k.SetPostCount(ctx, count+1)
	return count
}

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))
	// Convert PostCountKey to bytes
	byteKey := []byte(types.PostCountKey)
	// Get value of the count
	bz := store.Get(byteKey)
	//Return zero if the count value is not found
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey ("blog") and PostCountKey ("Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))
	// Convert PostCountKey to bytes
	byteKey := []byte(types.PostCountKey)
	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	// Set value of Post-count- to count
	store.Set(byteKey, bz)
}
