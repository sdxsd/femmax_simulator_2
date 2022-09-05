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
const MAX_TEXTURE	string = "assets/max.png"
const GAN_TEXTURE	string = "assets/gans.png"
const WIL_TEXTURE	string = "assets/will.png"
// Environment textures
const FLR_TEXTURE	string = "assets/env/floor.png"
const TBL_TEXTURE	string = "assets/env/table.png"
const VNT_TEXTURE	string = "assets/env/fvent.png"
const TCH_TEXTURE	string = "assets/env/t_chr.png"
const MCH_TEXTURE	string = "assets/env/m_chr.png"
// Map files
const LYR1			string = "maps/lyr1.txt"
const LYR2			string = "maps/lyr2.txt"
const LYR3			string = "maps/lyr3.txt"

// Generic NPC/Player entity.
type Entity struct {
	Dialogue, Name	string;
	Sprite			rl.Texture2D;
	x_pos, y_pos	int32;
}

type Prop struct {
	Sprite rl.Texture2D;
	x_pos, y_pos int32;
}

// Game state struct.
type Reality struct {
	Will, Fmax, Gans	Entity;
	lyr1, lyr2, lyr3	string;
	floor				rl.Texture2D;
	table, chair		Prop;
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

func parse_lyr1(lyr1 string) rl.Texture2D {
	var Width = int(WIDTH)
	var Height = int(HEIGHT)
	floor := rl.GenImageColor(Width, Height, rl.White);
	f_tile := rl.LoadImage(FLR_TEXTURE);
	v_tile := rl.LoadImage(VNT_TEXTURE);
	tile_rec := rl.Rectangle { 0, 0, float32(f_tile.Width), float32(f_tile.Height) };
	floor_rec := rl.Rectangle { 0, 0, float32(floor.Width), float32(floor.Height) };
	chars := []rune(lyr1);
	for i := 0; i < len(chars); i++ {
		if (chars[i] == 10) {
			tile_rec.Y += 64;
			tile_rec.X = 0;
		}
		if (chars[i] == '#') {
			rl.ImageDraw(floor, f_tile, floor_rec, tile_rec, rl.White);
			tile_rec.X += 64;
		}
		if (chars[i] == 'V') {
			rl.ImageDraw(floor, v_tile, floor_rec, tile_rec, rl.White);
			tile_rec.X += 64;
		}
	}
	rl.UnloadImage(f_tile);
	rl.UnloadImage(v_tile);
	floor_tex := rl.LoadTextureFromImage(floor);
	rl.UnloadImage(floor);
	return (floor_tex);
}

func parse_lyr2(table, chair *Prop, lyr2 string) {
	chars := []rune(lyr2);
	var x, y int32;

	for i := 0; i < len(chars); i++ {
		if (chars[i] == '\n') {
			y++;
			x = 0;
		}
		if (chars[i] == 'T') {
			table.x_pos = 64 * x;
			table.y_pos = 64 * y;
		}
		if (chars[i] == 'M') {
			chair.x_pos = 64 * x;
			chair.y_pos = 64 * y;
		}
		x++;
	}
}

func parse_lyr3(gans, will, femmax *Entity, lyr3 string) {
	chars := []rune(lyr3);
	var x, y int32;

	for i := 0; i < len(chars); i++ {
		if (chars[i] == '\n') {
			y++;
			x = 0;
		}
		if (chars[i] == 'M') {
			femmax.x_pos = 64 * x;
			femmax.y_pos = 64 * y;
		}
		if (chars[i] == 'G') {
			gans.x_pos = 64 * x;
			gans.y_pos = 64 * y;
		}
		if (chars[i] == 'W') {
			will.x_pos = 64 * x;
			will.y_pos = 64 * y;
		}
		x++;
	}
}

func cntr_pos(pos int32, character Entity, h_w bool) int32 {
	if (h_w == true) {
		return (pos - character.Sprite.Height / 3)
	}
	return (pos - character.Sprite.Width)
}

func main() {
	var game	Reality;
	var ret		error;

	rl.InitWindow(WIDTH, HEIGHT, "Larproom simulator");
	rl.SetTargetFPS(60);
	game.Gans = load_entity(GAN_TEXTURE, "Gansmeneer");
	game.Will = load_entity(WIL_TEXTURE, "Willem");
	game.Fmax = load_entity(MAX_TEXTURE, "Max")
	game.lyr1, ret = read_map_data(LYR1);
	if (ret != nil) { return };
	game.lyr2, ret = read_map_data(LYR2);
	if (ret != nil) { return };
	game.lyr3, ret = read_map_data(LYR3);
	if (ret != nil) { return };
	game.floor = parse_lyr1(game.lyr1);
	parse_lyr2(&game.table, &game.chair, game.lyr2);
	fmt.Println(game.table.x_pos, game.table.y_pos);
	game.table.Sprite = rl.LoadTexture(TBL_TEXTURE);
	game.chair.Sprite = rl.LoadTexture(MCH_TEXTURE);
	parse_lyr3(&game.Gans, &game.Will, &game.Fmax, game.lyr3)
	for (!rl.WindowShouldClose()) {
		rl.BeginDrawing()
			rl.ClearBackground(rl.White)
			rl.DrawTexture(game.floor, 0, 0, rl.White);
			rl.DrawTexture(game.table.Sprite, game.table.x_pos - game.table.Sprite.Width, game.table.y_pos - game.table.Sprite.Height / 2, rl.White);
			rl.DrawTexture(game.chair.Sprite, game.chair.x_pos, game.chair.y_pos, rl.White);
			rl.DrawTexture(game.Gans.Sprite, cntr_pos(game.Gans.x_pos, game.Gans, false), cntr_pos(game.Gans.y_pos, game.Gans, true), rl.White);
			rl.DrawTexture(game.Fmax.Sprite, cntr_pos(game.Fmax.x_pos, game.Fmax, false), cntr_pos(game.Fmax.y_pos, game.Fmax, true), rl.White);
			rl.DrawTexture(game.Will.Sprite, cntr_pos(game.Will.x_pos, game.Will, false), cntr_pos(game.Will.y_pos, game.Will, true), rl.White);
		rl.EndDrawing()
	}
}
