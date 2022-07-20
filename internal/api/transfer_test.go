package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/oramaz/tx-system/internal/db/mock"
	db "github.com/oramaz/tx-system/internal/db/sqlc"
	"github.com/stretchr/testify/require"
)

func TestTransferAPI(t *testing.T) {
	account1 := randomAccount()
	account2 := randomAccount()

	amount := int64(10)

	testCases := []struct {
		name          string
		body          transferRequest
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, resp *http.Response)
	}{
		{
			name: "OK",
			body: transferRequest{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.TransferTxParams{
					FromAccountID: account1.ID,
					ToAccountID:   account2.ID,
					Amount:        amount,
				}

				store.EXPECT().
					TransferTx(gomock.Any(), gomock.Eq(arg)).
					Times(1)
			},
			checkResponse: func(t *testing.T, resp *http.Response) {
				require.Equal(t, http.StatusOK, resp.StatusCode)
			},
		},
		{
			name: "NegativeAmount",
			body: transferRequest{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        -amount,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					TransferTx(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, resp *http.Response) {
				require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			},
		},
		{
			name: "TransferTxError",
			body: transferRequest{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					TransferTx(gomock.Any(), gomock.Any()).
					Times(1).Return(db.TransferTxResult{}, sql.ErrTxDone)
			},
			checkResponse: func(t *testing.T, resp *http.Response) {
				require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/transfers"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			request.Header.Set("Content-Type", "application/json")
			require.NoError(t, err)

			resp, err := server.router.Test(request)
			require.NoError(t, err)

			tc.checkResponse(t, resp)
		})
	}
}
