package main

var (
	CTemplate = `

{{$GSpritename := .Spritename}}

// #ifdef {{.Spritename}}_SPRITE_H
// #error "{{.Spritename}} already exists"
// #else
#ifndef {{.Spritename}}_SPRITE_H
#define {{.Spritename}}_SPRITE_H

#include "ww.h"

// declare sprite
extern ww_sprite_t * {{.Spritename}};
ww_sprite_t * {{.Spritename}} = NULL;

{{range .Animations}}
	
	// declare animation {{$GAnimation := .Animation}}
	extern ww_animation_t * {{$GSpritename}}_{{.Animation}};
	ww_animation_t * {{$GSpritename}}_{{.Animation}} = NULL;
	
	{{range .Frames}}
		
		//declare frame {{$GFrame := .Frame}}
		extern ww_frame_t * {{$GSpritename}}_{{$GAnimation}}_{{.Frame}};
		ww_frame_t * {{$GSpritename}}_{{$GAnimation}}_{{.Frame}} = NULL;
		
		
		{{range .Layers}}
		extern ww_polygon_t * {{$GSpritename}}_{{$GAnimation}}_{{$GFrame}}_{{.Layer}};
		ww_polygon_t * {{$GSpritename}}_{{$GAnimation}}_{{$GFrame}}_{{.Layer}} = NULL;
		{{end}}
		
		void init_{{$GSpritename}}_{{$GAnimation}}_{{$GFrame}}(){
		{{range .Layers}}
		
			{{if .Color}}
			//declare layer
			ww_rgba_t {{$GSpritename}}_{{$GAnimation}}_{{$GFrame}}_{{.Layer}}_color = { {{index .TrueColor 0}}, {{index .TrueColor 1}}, {{index .TrueColor 2}} };
			short {{$GSpritename}}_{{$GAnimation}}_{{$GFrame}}_{{.Layer}}_x[{{len .X}}] = { {{$lenX := dec (len .X)}}{{range $i, $e := .X}}{{$e}}{{if ne $i $lenX}}, {{end}}{{end}} };
			short {{$GSpritename}}_{{$GAnimation}}_{{$GFrame}}_{{.Layer}}_y[{{len .Y}}] = { {{$lenY := dec (len .Y)}}{{range $i, $e := .Y}}{{$e}}{{if ne $i $lenY}}, {{end}}{{end}} };
			{{$GSpritename}}_{{$GAnimation}}_{{$GFrame}}_{{.Layer}} = ww_new_polygon({{$GSpritename}}_{{$GAnimation}}_{{$GFrame}}_{{.Layer}}_color, {{$GSpritename}}_{{$GAnimation}}_{{$GFrame}}_{{.Layer}}_x, {{$GSpritename}}_{{$GAnimation}}_{{$GFrame}}_{{.Layer}}_y, {{len .X}});
			{{end}}
			
		{{end}}
		
			{{$GSpritename}}_{{$GAnimation}}_{{.Frame}} = ww_new_frame(
				{{range .Layers}}{{if .Color}}{{$GSpritename}}_{{$GAnimation}}_{{$GFrame}}_{{.Layer}},
				{{end}}{{end}}NULL 
			);
		}
		
	{{end}}
	
	void init_{{$GSpritename}}_{{$GAnimation}}(){
		{{range .Frames}}init_{{$GSpritename}}_{{$GAnimation}}_{{.Frame}}();
		{{end}}
		
		int delay[] = { {{$lenFrames := dec (len .Frames)}}{{range $i, $e := .Frames}}{{.Delay}}{{if ne $i $lenFrames}}, {{end}}{{end}} };
	
		{{$GSpritename}}_{{$GAnimation}} = ww_new_animation(
			delay,
			{{range .Frames}}{{$GSpritename}}_{{$GAnimation}}_{{.Frame}},
			{{end}}NULL 
		);
	}
  
{{end}}

void init_{{.Spritename}}(){
	
	{{range .Animations}}init_{{$GSpritename}}_{{.Animation}}();
	{{end}}

	{{.Spritename}} = ww_new_sprite( 0,
		{{range .Animations}}{{$GSpritename}}_{{.Animation}},
		{{end}}NULL
	);
	
}

#endif
`
)
