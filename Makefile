
WIDTH=1280
HEIGHT=720

all: images
	gcc src/*.c -O2 -Wall -pedantic -std=gnu11 -Isrc/ -Iimg/ -o main -lSDL2

images: tools
	./concoord/concoord img/*.yaml

tools:
	make -C concoord

memtest:
	valgrind --track-origins=yes --leak-check=yes ./main -S 1/8
