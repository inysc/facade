package hog

import "context"

func Op(ctx context.Context, op string) Event    { return DefaultLogger.Op(ctx, op) }
func Trace(ctx context.Context, op string) Event { return DefaultLogger.Trace(ctx, op) }
func Debug(ctx context.Context, op string) Event { return DefaultLogger.Debug(ctx, op) }
func Info(ctx context.Context, op string) Event  { return DefaultLogger.Info(ctx, op) }
func Warn(ctx context.Context, op string) Event  { return DefaultLogger.Warn(ctx, op) }
func Error(ctx context.Context, op string) Event { return DefaultLogger.Error(ctx, op) }
func Fatal(ctx context.Context, op string) Event { return DefaultLogger.Fatal(ctx, op) }
func Panic(ctx context.Context, op string) Event { return DefaultLogger.Panic(ctx, op) }
