package ports

type MockTestConnector interface {
	MockTest(val string) string
}
