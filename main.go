package main

import (
	"fmt"
	"container/list"
	"math/rand"
	"./dxlib"
)

type Actor struct {
	image int
	x     float32
	y     float32
	vx    float32
	vy    float32
	width int
}

func (actor *Actor) update() {
	actor.x += actor.vx
	actor.y += actor.vy
}
func (actor *Actor) draw() {
	x1 := int(actor.x) - actor.width
	y1 := int(actor.y) - actor.width
	x2 := int(actor.x) + actor.width
	y2 := int(actor.y) + actor.width
	dxlib.DrawExtendGraph(x1, y1, x2, y2, int(actor.image), 0)
}

func (actor *Actor) isOffScreen() bool {
	return actor.x > 690 || actor.x < -50 || actor.y > 520 || actor.y < -50
}

func main() {
	actors := list.New()

	dxlib.ChangeWindowMode(1)
	dxlib.SetWindowText("GO Example")

	dxlib.DxLib_Init()
	dxlib.SetDrawScreen(-2)

	var images map[string]int = make(map[string]int)
	for i := 0; i < 10 ; i++ {
		fname := fmt.Sprintf("image/image%02d.png", i)
		images[fname] = dxlib.LoadGraph(fname)
	}

	for i := 0; ; i++ {
		if dxlib.ProcessMessage() != 0 {
			break
		}
		dxlib.ClearDrawScreen()

		if i%2 == 0 {
			vx := 20.0 * (rand.Float32() - 0.5)
			vy := 20.0 * (rand.Float32() - 0.5)
			width := int(200.0 * (rand.Float32() - 0.5))
			fname := fmt.Sprintf("image/image%02d.png", rand.Int31n(9))
			image := images[fname]
			newActor := &Actor{x: 320, y: 240, vx: vx, vy: vy, width: width, image: image}
			actors.PushBack(newActor)
		}
		for e := actors.Front(); e != nil; {
			var actor = e.Value.(*Actor)

			if actor.isOffScreen() {
				prevE := e
				e = e.Next()
				actors.Remove(prevE)
			} else {
				actor.update()
				actor.draw()
				e = e.Next()
			}
		}

		dxlib.ScreenFlip()		
	}

	dxlib.DxLib_End()
}
