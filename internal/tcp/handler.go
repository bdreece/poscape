package tcp

import (
	"context"
	"errors"
	"net"
)

var ErrHandlerStopped = errors.New("tcp handler stopped")

type Handler struct {
	listener net.Listener
	context  context.Context
	cancel   context.CancelCauseFunc
}

func (handler *Handler) Start(ctx context.Context) {
	handler.context, handler.cancel = context.WithCancelCause(ctx)

	go handler.run()
}

func (handler *Handler) Stop(ctx context.Context) error {
	go func() {
		if err := handler.listener.Close(); err != nil {
			handler.cancel(err)
		}

		handler.cancel(ErrHandlerStopped)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-handler.context.Done():
		return handler.context.Err()
	}
}

func (handler *Handler) run() {
	ch := make(chan net.Conn, 1)
	go func() {
		for {
			conn, err := handler.listener.Accept()
			if err != nil {
				handler.cancel(err)
				return
			}

			ch <- conn
		}
	}()

	for {
		select {
		case <-handler.context.Done():
			return
		case conn := <-ch:
			go handler.accept(conn)
		}
	}
}

func (handler *Handler) accept(conn net.Conn) {
	for {

	}
}
