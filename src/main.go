package main

import (
	"fmt"
	"io/ioutil"
	"github.com/gen2brain/raylib-go/raylib"
)

// Window
const WIDTH			int32 = 448
const HEIGHT		int32 = 448
// Character textures
const MAX_TEXTURE	string = "../assets/max.png"
const GAN_TEXTURE	string = "../assets/gans.png"
const WIL_TEXTURE	string = "../assets/will.png"
// Environment textures
const FLR_TEXTURE	string = "../assets/env/floor.png"
const TBL_TEXTURE	string = "../assets/env/table.png"
const VNT_TEXTURE	string = "../assets/env/fvent.png"
const TCH_TEXTURE	string = "../assets/env/t_chr.png"
const MCH_TEXTURE	string = "../assets/env/m_chr.png"
// Map files
const LYR1			string = "../maps/lyr1.txt"
const LYR2			string = "../maps/lyr2.txt"
const LYR3			string = "../maps/lyr3.txt"

// Generic NPC/Player entity.
type Entity struct {
	Dialogue, Name	string;
	Sprite			rl.Texture2D;
	x_pos, y_pos	int8
}

// Game state struct.
type Reality struct {
	Will, Fmax, Gans	Entity;
	lyr1, lyr2, lyr3	string;
	floor				rl.Texture2D
}

func read_map_data(file string) (string, error) {
	f_contents, err_v := ioutil.ReadFile(file);
	if (err_v != nil) {
		fmt.Println("ERROR: %s", err_v);
		return "", err_v;
	}
	return string(f_contents), err_v;
}

func load_entity(texture, title string) Entity {
	var new_entity	Entity;

	new_entity.Name = title;
	new_entity.Sprite = rl.LoadTexture(texture);
	return (new_entity);
}

func floor_tex(lyr1 string) rl.Texture2D {
	var curr_x	float32 = 0.0;
	var curr_y	float32 = 0.0;
	var Width	= int(WIDTH)
	var Height	= int(HEIGHT)
	floor := rl.GenImageColor(Width, Height, rl.White);
	f_tile := rl.LoadImage(FLR_TEXTURE);
	f_tile_rec := rl.Rectangle { curr_x, curr_y, float32(f_tile.Width), float32(f_tile.Height) };
	floor_rec := rl.Rectangle {0, 0, float32(floor.Width), float32(floor.Height) };

	for ()
}

func main() {
	var game	Reality;
	var ret		error;

	rl.InitWindow(WIDTH, HEIGHT, "Larproom simulator");
	game.Gans = load_entity(GAN_TEXTURE, "Gansmeneer");
	game.Will = load_entity(WIL_TEXTURE, "Willem");
	game.Fmax = load_entity(MAX_TEXTURE, "Max")
	game.lyr1, ret = read_map_data(LYR1);
	if (ret != nil) { return };
	game.lyr2, ret = read_map_data(LYR2);
	if (ret != nil) { return };
	game.lyr3, ret = read_map_data(LYR3);
	if (ret != nil) { return };
	for (!rl.WindowShouldClose()) {
		rl.BeginDrawing()
			rl.ClearBackground(rl.White)
		rl.EndDrawing()
	}
}
