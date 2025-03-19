package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/t57r/gator/internal/config"
	"github.com/t57r/gator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error opening DB connection: %v", err)
	}
	dbQueries := database.New(db)

	s := state{
		db:     dbQueries,
		config: &cfg,
	}
	cmds := commands{
		hanlders: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerFetchFeed)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", handlerFollow)
	cmds.register("following", handlerListFeedFollows)

	if len(os.Args) < 2 {
		fmt.Println("Please add args to the command")
		os.Exit(1)
	}
	commandName := os.Args[1]

	err = cmds.run(&s, command{
		name: commandName,
		args: os.Args[1:],
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
