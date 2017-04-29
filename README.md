# dxlib-go
A minimal example of a DxLib application written in Go

## Build
```sh
go get golang.org/x/text/encoding/japanese
go build
```

## Example of main loop
```go
dxlib.DxLib_Init()
dxlib.SetDrawScreen(-2)
for {
	if dxlib.ProcessMessage() != 0 {
		break
	}
	dxlib.ClearDrawScreen()

	// something...
	
	dxlib.ScreenFlip()
}

dxlib.DxLib_End()
```
