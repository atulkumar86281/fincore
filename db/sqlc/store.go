package db

import (
	"context"
	"database/sql"
	"fmt"
)


type Store struct{
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store{

	return &Store{
		db : db,
		Queries: New(db),

	}
}


func (store *Store) execTx(ctx context.Context, fn func(*Queries) error ) error{
	tx, err := store.db.BeginTx(ctx,nil)
	if err != nil{
		return err
	}

	q := New(tx)

	err = fn(q)

	if err != nil{
		if rbErr := tx.Rollback(); rbErr != nil{
			return fmt.Errorf("tx error: %v, rb error: %v",err,rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct{
	FromAccountId int64 `json:"from_account_id"`
	ToAccountId int64 `json:"to_account_id"`
	Amount int64 `json:"amount"`
}

type TrasnferTxResult struct{
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"from_account_id"`
	ToAccount Account	`json:"to_account_id"`
	FromEntry Entry `json:"from_entry"`
	ToEntry Entry `json:"to_entry"`
}

// var txnKey = struct{}{}

func (store *Store) TransferTx(ctx context.Context, transferTxArg TransferTxParams) (TrasnferTxResult,error){
	var result TrasnferTxResult

	err := store.execTx(ctx,func(q *Queries)error{
		var err error


		// Creating a transfer
		result.Transfer, err = q.CreateTransfer(ctx,CreateTransferParams{
			FromAccountID: transferTxArg.FromAccountId,
			ToAccountID: transferTxArg.ToAccountId,
			Amount: transferTxArg.Amount,
		})
		if err != nil{
			return err
		}

		// creating a deduct entry in from account
		result.FromEntry,err = q.CreateEntry(ctx,CreateEntryParams{
			AccountID: transferTxArg.FromAccountId,
			Amount: -transferTxArg.Amount,
		})

		if err != nil{
			return err
		}

		// creating a credit entry in To account
		result.ToEntry,err = q.CreateEntry(ctx,CreateEntryParams{
			AccountID: transferTxArg.ToAccountId,
			Amount: transferTxArg.Amount,
		})

		if err != nil{
			return err
		}

		//Get account -> Update its balance
		account1 ,err := q.GetAccountForUpdate(ctx,transferTxArg.FromAccountId)

		if err != nil{
			return err
		}

		result.FromAccount, err = q.UpdateAccount(ctx,UpdateAccountParams{
			ID: transferTxArg.FromAccountId,
			Balance: account1.Balance - transferTxArg.Amount,
		})

		if err != nil{
			return err
		}

		account2 ,err := q.GetAccountForUpdate(ctx,transferTxArg.ToAccountId)

		if err != nil{
			return err
		}

		result.ToAccount, err = q.UpdateAccount(ctx,UpdateAccountParams{
			ID: transferTxArg.ToAccountId,
			Balance: account2.Balance + transferTxArg.Amount,
		})

		if err != nil{
			return err
		}

		return nil
	})

	return result,err
}