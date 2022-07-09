package ports

// KGS is a obstract to the Key Generator Service
type KGS interface {
	GetKey() string
}
