package cli

import (
	"fmt"
	gcli "github.com/urfave/cli"

	skyWallet "github.com/fibercrypto/skywallet-go/src/skywallet"
)

func checkMessageSignatureCmd() gcli.Command {
	name := "checkMessageSignature"
	return gcli.Command{
		Name:        name,
		Usage:       "Check a message signature matches the given address.",
		Description: "",
		Flags: []gcli.Flag{
			gcli.StringFlag{
				Name:  "message",
				Usage: "The message that the signature claims to be signing.",
			},
			gcli.StringFlag{
				Name:  "signature",
				Usage: "Signature of the message.",
			},
			gcli.StringFlag{
				Name:  "address",
				Usage: "Address that issued the signature.",
			},
			gcli.StringFlag{
				Name:   "deviceType",
				Usage:  "Device type to send instructions to, hardware wallet (USB) or emulator.",
				EnvVar: "DEVICE_TYPE",
			},
		},
		OnUsageError: onCommandUsageError(name),
		Action: func(c *gcli.Context) {
			message := c.String("message")
			signature := c.String("signature")
			address := c.String("address")
			sq, err := createDevice(c.String("deviceType"))
			if err != nil {
				return
			}
			msg, err := sq.CheckMessageSignature(message, signature, address)
			if err != nil {
				log.Error(err)
				return
			}

			responseMsg, err := skyWallet.DecodeSuccessOrFailMsg(msg)
			if err != nil {
				log.Error(err)
				return
			}

			fmt.Println(responseMsg)
		},
	}
}
