package background

import (
	"context"
	"encoding/binary"
	"errors"
	"io"
	"log/slog"
	"net"
	"syscall"

	"github.com/Nyarum/diho_gpkov2/internal/actor"
	"github.com/Nyarum/diho_gpkov2/internal/actorhandler"
	"github.com/Nyarum/diho_gpkov2/internal/packets"
)

func NewTCP(ctx context.Context, addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		slog.Error("Error creating listener", "error", err)
		return err
	}
	defer listener.Close()

	slog.Info("Server is listening on port", "addr", addr)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			conn, err := listener.Accept()
			if err != nil {
				slog.Error("Error accepting connection", "error", err)
				return err
			}

			slog.Info("Accepted connect", "addr", conn.RemoteAddr())

			pktBuf, err := packets.EncodeWithHeader(ctx, packets.NewFirstTime(), binary.BigEndian)
			if err != nil {
				return err
			}

			_, err = conn.Write(pktBuf)
			if err != nil {
				return err
			}

			go func() {
				err := connection(ctx, conn)
				if err != nil {
					if isNetConnClosedErr(err) {
						slog.Info("Connection closed", "addr", conn.RemoteAddr())
						return
					}

					slog.Error("Connection error", "error", err)
				}
			}()
		}
	}
}

func connection(ctx context.Context, conn net.Conn) error {
	actorClient := actor.NewActor("client", actorhandler.NewClient(conn)).Start(ctx)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			buf := make([]byte, 2048)

			ln, err := conn.Read(buf)
			if err != nil {
				return err
			}

			if ln == 2 {
				conn.Write([]byte{0, 2})
				continue
			}

			header, err := packets.DecodeHeader(buf)
			if err != nil {
				return err
			}

			buf = buf[8:]

			slog.Info("Header", "header", header)

			if ln < int(header.Len) {
				moreData := make([]byte, int(header.Len)-ln)
				_, err := conn.Read(buf)
				if err != nil {
					return err
				}

				buf = append(buf, moreData...)
			}

			actorClient.Send(actorhandler.IncomePacket{
				Opcode: header.Opcode,
				Data:   buf,
			})
		}
	}
}

func isNetConnClosedErr(err error) bool {
	switch {
	case
		errors.Is(err, net.ErrClosed),
		errors.Is(err, io.EOF),
		errors.Is(err, syscall.EPIPE):
		return true
	default:
		return false
	}
}
