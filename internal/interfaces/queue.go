package interfaces

type MessageQueuer interface {
	ProduceMessages() error
	EmmitMessages() error
}
