package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Bot parameters
var (
	GuildID        = "1148626765577723914" //flag.String("guild", "", "Test guild ID")
	BotToken       = flag.String("token", "", "Bot access token")
	AppID          = "1148617230129573991" //flag.String("app", "", "Application ID")
	Cleanup        = flag.Bool("cleanup", true, "Cleanup of commands")
	ResultsChannel = "général" //flag.String("results", "", "Channel where send survey results to")
)

var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

// Important note: call every command in order it's placed in the example.
var lastMessage = "Vous maintenant accès au reste du Discord ! Bienvenue à vous et have fun ;-)";
var isFormFinished = false

var (
	componentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"profession": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var response *discordgo.InteractionResponse

			data := i.MessageComponentData()
			switch data.Values[0] {
			case "freelance":
				isFormFinished = true;
				response = &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Tu es un mercenaire.",
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				}
			case "recruteur":
				isFormFinished = true;
				response = &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Tu es un barman.",
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				}
			default:
				response = &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Quel est votre niveau d'XP ?",
						Flags:   discordgo.MessageFlagsEphemeral,
						Components: []discordgo.MessageComponent{
							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									discordgo.SelectMenu{
										// Select menu, as other components, must have a customID, so we set it to this value.
										CustomID:    "xp",
										Placeholder: "Tu te situes où là concrètement ?",
										Options: []discordgo.SelectMenuOption{
											{
												Label: "Etudiant (ou en formation)",
												// As with components, this things must have their own unique "id" to identify which is which.
												// In this case such id is Value field.
												Value: "etudiant",
												Emoji: discordgo.ComponentEmoji{
													Name: "🦦",
												},
												// You can also make it a default option, but in this case we won't.
												Default:     false,
												Description: "Tu es en plein apprentissage des arcanes du développement.",
											},
											{
												Label: "Junior",
												Value: "junior",
												Emoji: discordgo.ComponentEmoji{
													Name: "🟨",
												},
												Description: "Tu viens de commencer dans le chemin de développement et tu as bien raison !",
											},
											{
												Label: "Confirmé",
												Value: "confirme",
												Emoji: discordgo.ComponentEmoji{
													Name: "🐍",
												},
												Description: "Tu commences à te sentir bien dans tes baskets.",
											},
											{
												Label: "Senior",
												Value: "senior",
												Emoji: discordgo.ComponentEmoji{
													Name: "🐍",
												},
												Description: "Tu es le sensei de toute ta feature team !",
											},
										},
									},
								},
							},
						},
					},
				}
			}
			err := s.InteractionRespond(i.Interaction, response)
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Second) // Doing that so user won't see instant response.
			if isFormFinished {
				_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Content: lastMessage,
					Flags: discordgo.MessageFlagsEphemeral,	
				})
			}

			if err != nil {
				panic(err)
			}
		},
		"xp": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var response *discordgo.InteractionResponse

			data := i.MessageComponentData()
			switch data.Values[0] {
			case "etudiant":
				isFormFinished = true
				response = &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Tu es un vagabond.",
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				}
			case "junior":
				isFormFinished = true
				response = &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Tu es un client.",
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				}
			case "confirme":
				isFormFinished = true
				response = &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Tu es un client.",
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				}
			case "senior":
				response = &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Avez-vous le gout de la transmission du savoir et de la connaissance ?",
						Flags:   discordgo.MessageFlagsEphemeral,
						Components: []discordgo.MessageComponent{
							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									discordgo.SelectMenu{
										CustomID:    "transmission",
										Placeholder: "",
										Options: []discordgo.SelectMenuOption{
											{
												Label: "Oui",
												Value: "yes",
												Emoji: discordgo.ComponentEmoji{
													Name: "🦦",
												},
												Description: "Tu veux partager tes connaissances.",
											},
											{
												Label: "Non",
												Value: "no",
												Emoji: discordgo.ComponentEmoji{
													Name: "🟨",
												},
												Description: "",
											},
										},
									},
								},
							},
						},
					},
				}
			}
			err := s.InteractionRespond(i.Interaction, response)
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Second) // Doing that so user won't see instant response.
			if isFormFinished {
				_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Content: lastMessage,
					Flags: discordgo.MessageFlagsEphemeral,
					
				})
			}

			if err != nil {
				panic(err)
			}
		},
		"transmission": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var response *discordgo.InteractionResponse
			data := i.MessageComponentData()
			switch data.Values[0] {
			case "yes":
				response = &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Tu es un menestrel.",
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				}
			case "no":
				response = &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Tu es un client.",
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				}
			}
			err := s.InteractionRespond(i.Interaction, response)
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Second) // Doing that so user won't see instant response.
			_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
				Content: lastMessage,
				Flags: discordgo.MessageFlagsEphemeral,
				
			})
			if err != nil {
				panic(err)
			}
		},
	}
	commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"formula": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var response *discordgo.InteractionResponse
			switch i.ApplicationCommandData().Options[0].Name {
			case "begin":
				if (isFormFinished) {	
					response = &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Hop, hop, hop, tu as déjà fait le formulaire !",
							Flags:   discordgo.MessageFlagsEphemeral,
						},
					}
				} else {
					response = &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Bienvenue dans notre auberge des devs. Nous allons vous poser quelques questions afin de vous attribuer un rôle au sein de cette auberge. Première chose : ",
							Flags:   discordgo.MessageFlagsEphemeral,
							Components: []discordgo.MessageComponent{
								discordgo.ActionsRow{
									Components: []discordgo.MessageComponent{
										discordgo.SelectMenu{
											// Select menu, as other components, must have a customID, so we set it to this value.
											CustomID:    "profession",
											Placeholder: "Dans quelle position professionnelle es-tu ?",
											Options: []discordgo.SelectMenuOption{
												{
													Label: "Freelance",
													// As with components, this things must have their own unique "id" to identify which is which.
													// In this case such id is Value field.
													Value: "freelance",
													Emoji: discordgo.ComponentEmoji{
														Name: "🦦",
													},
													// You can also make it a default option, but in this case we won't.
													Default:     false,
													Description: "Tu es un developpeur à la recherche de missions en solitaire",
												},
												{
													Label: "Recruteur",
													Value: "recruteur",
													Emoji: discordgo.ComponentEmoji{
														Name: "🟨",
													},
													Description: "Tu es un chasseur de têtes à la recherche de la perle rare",
												},
												{
													Label: "Ni l'un ni l'autre",
													Value: "default",
													Emoji: discordgo.ComponentEmoji{
														Name: "🐍",
													},
													Description: "Salarié/Consultant",
												},
											},
										},
									},
								},
							},
						},
					}
				}

			}
			err := s.InteractionRespond(i.Interaction, response)
			if err != nil {
				panic(err)
			}
		},
	}
)

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})
	// Components are part of interactions, so we register InteractionCreate handler
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commandsHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionMessageComponent:

			if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
		}
	})
	_, err := s.ApplicationCommandCreate(AppID, GuildID, &discordgo.ApplicationCommand{
		Name:        "buttons",
		Description: "Test the buttons if you got courage",
	})

	if err != nil {
		log.Fatalf("Cannot create slash command: %v", err)
	}
	_, err = s.ApplicationCommandCreate(AppID, GuildID, &discordgo.ApplicationCommand{
		Name: "formula",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "begin",
				Description: "Begin formula",
			},
		},
		Description: "Lo and behold: dropdowns are coming",
	})

	if err != nil {
		log.Fatalf("Cannot create slash command: %v", err)
	}

	err = s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")
}