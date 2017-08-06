package main

import (
    "github.com/evandroflores/claimr/database"
    "github.com/evandroflores/claimr/cmd"
    "github.com/shomali11/slacker"
    "log"
    "fmt"
)

func init() {
    database.InitDB()
}

func main() {
    bot := slacker.NewClient("xoxb-221107798822-MhoNS4UseJkvo5azVFKRjpud")

    log.Println("Loading commands...")
    for _, command := range cmd.CommandList() {
        log.Println(fmt.Sprintf("%s - %s", command.Usage, command.Description))
        bot.Command(command.Usage, command.Description, command.Handler)
    }
    log.Println(fmt.Sprintf("%d Commands loaded.", len(cmd.CommandList())))

    err := bot.Listen()
    if err != nil {
        log.Fatal(err)
    }
}
