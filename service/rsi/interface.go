package rsi

type LoginServicer interface {
	EmailChecking(*EmailLoginBody) error
}
