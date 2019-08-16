
#include <stdio.h>

#include "ww.h"

int main( int argc, char * argv[] ) {
	
	if(ww_window_create(argc, argv, "Pixarray", 1024, 576)) {
		printf("Closing..\n");
		return 1;
	}
	
	while(!ww_window_received_quit_event()) {
		
		ww_window_update_events();
		ww_window_update_buffer();
	}
	
	ww_window_destroy();
	return 0;
}
