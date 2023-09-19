package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const (
	mercenairesRoleID = 1153682851292778646
	barmanRoleID      = 1153682995706859641
	guildID           = 1148626765577723914
)

var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	discord.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Retrieve a list of members for the specified guild.
	members, err := discordgo.GuildMembers(guildID, "", 1000) // You can adjust the limit as needed.
	if err != nil {
		log.Println("Error fetching members:", err)
	}

	// Iterate through the list of members and access their roles.
	for _, member := range members {
		roles := getMemberRoles(member)

		// Print or process member roles as needed.
		log.Printf("Member %s#%s Roles: %v", member.User.Username, member.User.Discriminator, roles)
	}

	// Cleanly close down the Discord session.
	discord.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "/lol" {
		s.ChannelMessageSend(m.ChannelID, "mdr!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "/mdr" {
		s.ChannelMessageSend(m.ChannelID, "lol!")
	}
}

// getMemberRoles retrieves the roles for a given member.
func getMemberRoles(member *discordgo.Member) []string {
	roles := make([]string, len(member.Roles))

	for i, roleID := range member.Roles {
		roles[i] = roleID
	}

	return roles
}
