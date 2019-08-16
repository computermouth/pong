
#include <stdio.h>

#include "ww.h"
#include "player_one.h"
#include "player_two.h"
#include "ball.h"

int screen_width = 1024;
int screen_height = 576;

int main( int argc, char * argv[] ) {
	
	if(ww_window_create(argc, argv, "Pixarray", screen_width, screen_height)) {
		printf("Closing..\n");
		return 1;
	}
	
	init_player_one();
	init_player_two();
	init_ball();
	
	int y_dir = -1;
	int x_dir = -1;
	
	player_two->pad_x = screen_width - player_two->width;
	ball->pad_x = (screen_width  / 2) - (ball->width  / 2);
	ball->pad_y = (screen_height / 2) - (ball->height / 2);
	
	while(!ww_window_received_quit_event()) {
		ww_window_update_events();
		
		// player 1 move up
		if (keystate.w == 1 && player_one->pad_y >= 0){
			player_one->pad_y = player_one->pad_y - 2;
		}
		// player 1 move down
		if (keystate.s == 1 && player_one->pad_y < screen_height - player_one->height) {
			player_one->pad_y = player_one->pad_y + 2;
		}
		// player 2 move up
		if (keystate.up == 1 && player_two->pad_y >= 0) {
			player_two->pad_y = player_two->pad_y - 2;
		}
		// player 2 move down
		if (keystate.dn == 1 && player_two->pad_y < screen_height - player_two->height) {
			player_two->pad_y = player_two->pad_y + 2;
		}
		
		ball->pad_y = ball->pad_y + y_dir;
		ball->pad_x = ball->pad_x + x_dir;
		
		// bounce off bottom
		if (ball->pad_y > screen_height - ball->height || ball->pad_y < 0){
			y_dir = y_dir * -1;
		}
		
		// move ball to center if it goes off right side or left side
		if (ball->pad_x > screen_width || ball->pad_x < 0 ) {
			ball->pad_x = (screen_width  / 2) - (ball->width  / 2);
			ball->pad_y = (screen_height / 2) - (ball->height / 2);
		}
		
		ww_draw_sprite(player_one);
		ww_draw_sprite(player_two);
		ww_draw_sprite(ball);
		
		ww_window_update_buffer();
	}
	
	ww_window_destroy();
	return 0;
}
