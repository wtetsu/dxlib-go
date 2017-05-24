# dxlib-go
A minimal example of a DxLib application written in Go

## Build
```sh
go get golang.org/x/text/encoding/japanese
go build
```

![ss01](https://raw.githubusercontent.com/wtetsu/dxlib-go/images/ss01.png)

animation gif
![ss02](https://raw.githubusercontent.com/wtetsu/dxlib-go/images/ss02.gif)

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
