package logger

import (
	"context"
	"log/slog"

	ctxpkg "github.com/AidanThomas/iris/internal/core/context"
)

var _ slog.Handler = (*Handler)(nil)

type Handler struct {
	handler slog.Handler
}

func newHandler(handler slog.Handler) *Handler {
	return &Handler{handler: handler}
}

func (h *Handler) Handle(ctx context.Context, record slog.Record) error {
	correlationID := ctxpkg.GetCorrelationID(ctx)
	if correlationID != "" {
		record = record.Clone()
		record.AddAttrs(slog.String("correlationID", correlationID))
	}
	return h.handler.Handle(ctx, record)
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return newHandler(h.handler.WithAttrs(attrs))
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return newHandler(h.handler.WithGroup(name))
}

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}
