module github.com/filemingo/librsync-dart-wrapper

go 1.20

require (
	github.com/balena-os/librsync-go v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.3.0
	github.com/sirupsen/logrus v1.9.0
)

require (
	github.com/balena-os/circbuf v0.1.3 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
)

replace github.com/balena-os/librsync-go => github.com/gurupras/librsync-go v0.9.0-gurupras-1
