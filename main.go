package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/zepetriko/go_aggregator/internal/cli"
	"github.com/zepetriko/go_aggregator/internal/config"
	"github.com/zepetriko/go_aggregator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	dbQueries := database.New(db)

	programState := &cli.State{
		Db:     dbQueries,
		Config: &cfg,
	}

	commands := cli.NewCommands()
	commands.Register("login", cli.HandlerLogin)
	commands.Register("register", cli.HandlerRegister)
	commands.Register("reset", cli.HandlerReset)
	commands.Register("users", cli.HandlerUsers)
	commands.Register("agg", cli.HandlerAgg)
	commands.Register("addfeed", cli.HandlerAddFeed)
	commands.Register("feeds", cli.HandlerFeeds)

	if len(os.Args) < 2 {
		fmt.Println("error: not enough arguments")
		os.Exit(1)
	}

	cmd := cli.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := commands.Run(programState, cmd); err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

}
