all: android ios tauri

gomobile:
	go get golang.org/x/mobile

android:  gomobile
	mkdir -p ~/workspace/filemingo/librsync_native/android/libs
	gomobile bind -target=android -o ~/workspace/filemingo/librsync_native/android/libs/librsyncbridge.aar github.com/filemingo/librsync-dart-wrapper/librsyncbridge

ios: gomobile
	gomobile bind -target=ios -o LibRsyncBridge.xcframework github.com/filemingo/librsync-dart-wrapper/librsyncbridge/ios

flutter:
	cd desktop && go build -o librsyncbridge.so -buildmode=c-shared main.go
	# Flutter-native
	cp desktop/librsyncbridge.so ~/workspace/filemingo/librsync_native/linux/
	cp desktop/librsyncbridge.h ~/workspace/filemingo/librsync_native/linux/include/gobridge/


tauri:
	cd desktop && go build -o librsyncbridge.a -buildmode=c-archive main.go
	mv desktop/librsyncbridge.a ~/workspace/filemingo/filemingo/src-tauri/librsyncbridge.a
	bindgen desktop/librsyncbridge.h -o ~/workspace/filemingo/filemingo/src-tauri/src/librsync_ffi.rs


