package mysql

import (
	"context"
	"database/sql"
	"github.com/go-qbit/qerror"
	"github.com/go-qbit/timelog"
	"strconv"
	"sync"
	"unsafe"
)

const ctx_transaction_key = "MYSQL_TRANSACTION"

type transaction struct {
	tx           *sql.Tx
	savePoint    uint64
	savePointMtx sync.Mutex
}

func (s *MySQL) StartTransaction(ctx context.Context) (context.Context, error) {
	t := ctx.Value(s.transactionKey())

	if t == nil {
		ctx = timelog.Start(ctx, "BEGIN")
		tx, err := s.db.Begin()
		ctx = timelog.Finish(ctx)
		if err != nil {
			return nil, err
		}

		return context.WithValue(ctx, s.transactionKey(), &transaction{
			tx: tx,
		}), nil
	} else {
		t := t.(*transaction)
		t.savePointMtx.Lock()
		defer t.savePointMtx.Unlock()

		t.savePoint++

		_, err := t.tx.Exec("SAVEPOINT SP" + strconv.FormatUint(t.savePoint, 10))
		if err != nil {
			return nil, err
		}

		return ctx, nil
	}
}

func (s *MySQL) Commit(ctx context.Context) (context.Context, error) {
	ct := ctx.Value(s.transactionKey())

	if ct == nil {
		return nil, qerror.New("No started transaction")
	}

	t := ct.(*transaction)
	t.savePointMtx.Lock()
	defer t.savePointMtx.Unlock()

	if t.savePoint > 0 {
		_, err := t.tx.Exec("RELEASE SAVEPOINT SP" + strconv.FormatUint(t.savePoint, 10))
		if err != nil {
			return nil, err
		}

		t.savePoint--

		return ctx, nil
	}

	ctx = timelog.Start(ctx, "COMMIT")
	err := t.tx.Commit()
	ctx = timelog.Finish(ctx)
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, s.transactionKey(), nil), nil
}

func (s *MySQL) Rollback(ctx context.Context) (context.Context, error) {
	ct := ctx.Value(s.transactionKey())

	if ct == nil {
		return nil, qerror.New("No started transaction")
	}

	t := ct.(*transaction)
	t.savePointMtx.Lock()
	defer t.savePointMtx.Unlock()

	if t.savePoint > 0 {
		_, err := t.tx.Exec("ROLLBACK TO SAVEPOINT SP" + strconv.FormatUint(t.savePoint, 10))
		if err != nil {
			return nil, err
		}

		t.savePoint--

		return ctx, nil
	}

	ctx = timelog.Start(ctx, "ROLLBACK")
	err := t.tx.Rollback()
	ctx = timelog.Finish(ctx)
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, s.transactionKey(), nil), nil
}

func (s *MySQL) transactionKey() string {
	return ctx_transaction_key + strconv.FormatInt(int64(uintptr(unsafe.Pointer(s))), 10)
}
