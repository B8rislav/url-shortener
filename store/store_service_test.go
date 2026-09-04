package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}
var ctx = context.Background()

func init() {
	InitStore(ctx)
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "naraimashou.com/user/fe-efjn-wdwdw"
	userUUId := "kfekf923i2-febfkje3432-dfnekfn342"
	shortUrl := "jneFef"

	SaveUrl(ctx, initialLink, shortUrl, userUUId)

	retrievedUrl, _ := GetUrl(ctx, shortUrl)

	assert.Equal(t, initialLink, retrievedUrl)
}
