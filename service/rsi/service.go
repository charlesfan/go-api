package rsi

var (
	// === Repository ===
	// === Service ===
	LoginService LoginServicer
)

func Init() {
	LoginService = NewLoginService()
}
