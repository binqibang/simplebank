package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store defines all functions to execute db queries and transactions.
type Store struct {
	db *sql.DB
	// Embedded struct, like Inheritance in OOP.
	*Queries
}

// NewStore creates a `Store`.
func NewStore(db *sql.DB) *Store {
	return &Store{
		db,
		New(db),
	}
}

// ExecTx executes a function within a database transaction.
func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// Starts a transaction.
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	// Execute the query.
	err = fn(q)
	if err != nil {
		// If rollback failed.
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	// Commit the transaction.
	return tx.Commit()
}

// TransferTxParams contains the input parameters of the transfer transaction.
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction.
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to the other.
// It creates the transfer, add account entries, and update accounts'
// balance within a database transaction.
func (s *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	// Execute the transaction by `execTx()`.
	err := s.execTx(ctx, func(q *Queries) error {
		var err error
		// Create a transfer.
		result.Transfer, err = s.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// Create the from entry.
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		// Create the to entry.
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// Change balance.
		// Make sure we first update the smaller id account,
		// to avoid the deadlock.
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID,
				-arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID,
				arg.Amount, arg.FromAccountID, -arg.Amount)
		}

		return err
	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (account1 Account, account2 Account, err error) {
	account1, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})
	return
}
