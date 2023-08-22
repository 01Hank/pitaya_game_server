package dbmodule

type BaseDBIn interface {
	Connect() error
	DisConnect()
}
