package network 


var (
	ErrInvalidPayload = SandboxResponse{
		StatusCode: 400,
		Message: "Invalid payload",
	}
)