all: clean game

clean:
	rm game.run 2>/dev/null || true

game.run:
	@go build -o game.run ./game

play: clean game.run
	./game.run

