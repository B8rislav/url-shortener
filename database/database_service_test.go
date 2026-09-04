package database

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()

func init() {
	NewPool(ctx)
}

func TestPoolInit(t *testing.T) {
	assert.True(t, dbService.pool != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "naraimashou.com/user/fe-efjn-wdwdw"
	shortUrl := "jneFef"
	defer dbService.pool.Exec(ctx, `DELETE FROM urls WHERE shorturl = $1`, shortUrl)

	SaveUrl(ctx, shortUrl, initialLink)

	retrievedUrl, err := GetUrl(ctx, shortUrl)

	assert.NoError(t, err)
	assert.Equal(t, initialLink, retrievedUrl)
}

func TestRetrievalOfUnknownUrl(t *testing.T) {
	retrievedUrl, err := GetUrl(ctx, "notSaved")

	assert.ErrorIs(t, err, pgx.ErrNoRows)
	assert.Empty(t, retrievedUrl)
}
