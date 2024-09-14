package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	g := newGame(1, "god of war", 50, "action adventure")

	assert.Equal(t, 1,      g.id, "id")
	assert.Equal(t, "god of war",  g.name, "name")
	assert.Equal(t, 50,   g.price, "price")
	assert.Equal(t, "action adventure", g.genre, "genre")
}

func TestString(t *testing.T) {
	g := newGame(2, "x-com 2", 30, "strategy")

	assert.Equal(t, "2: x-com 2 costs 30", g.item.String(), "item string")
	assert.Equal(t, "Game 2: x-com 2 costs 30 of genre strategy", g.String(), "game string")
}

func TestList(t *testing.T) {
	games := newGameList()

	assert.Len(t, games, 3, "gamelist len")

 	assert.Equal(t, 1,      games[0].id, "game 0 id")
	assert.Equal(t, "god of war",  games[0].name, "game 0 name")
	assert.Equal(t, 50,   games[0].price, "game 0 price")
	assert.Equal(t, "action adventure", games[0].genre, "game 0 genre")

 	assert.Equal(t, 3,      games[1].id, "game 1 id")
	assert.Equal(t, "minecraft",  games[1].name, "game 1 name")
	assert.Equal(t, 20,   games[1].price, "game 1 price")
	assert.Equal(t, "sandbox", games[1].genre, "game 1 genre")

 	assert.Equal(t, 4,      games[2].id, "game 2 id")
	assert.Equal(t, "warcraft",  games[2].name, "game 2 name")
	assert.Equal(t, 40,   games[2].price, "game 2 price")
	assert.Equal(t, "strategy", games[2].genre, "game 2 genre")
}

func TestById(t *testing.T) {
	g, err := queryById(newGameList(), 3)
	assert.NoError(t, err, "no error")
 	assert.Equal(t, 3,      g.id, "game id")
	assert.Equal(t, "minecraft",  g.name, "game name")
	assert.Equal(t, 20,   g.price, "game price")
	assert.Equal(t, "sandbox", g.genre, "game genre")

	g, err = queryById(newGameList(), 11)
	assert.EqualError(t, err, "No such game", "error")
}

func TestNameByPrice(t *testing.T) {
	names := listNameByPrice(newGameList(), 21)

	assert.Len(t, names, 1, "namelist len")
	if 1 > 0 {
	 	assert.Equal(t, "minecraft", names[0], "names 1")
	}
	if 1 > 1 {
	 	assert.Equal(t, "N/A", names[1], "names 2")
	}
	if 1 > 2 {
	 	assert.Equal(t, "N/A", names[2], "names 3")
	}
	if 1 > 3 {
	 	assert.Equal(t, "N/A", names[3], "names 4")
	}
}

