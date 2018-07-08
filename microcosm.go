package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nkashy1/microcosm/accounts"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <subcommand> [arguments]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Subcommands: accounts, addresses")
		fmt.Fprintln(os.Stderr, "For information on any <subcommand>:")
		fmt.Fprintf(os.Stderr, "\t%s <subcommand> -h\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	var keystore, password string
	var numAccounts uint

	createAccountFlags := flag.NewFlagSet("accounts", flag.ExitOnError)
	createAccountFlags.StringVar(&keystore, "keystore", "./", "Directory in which to store key file")
	createAccountFlags.StringVar(&password, "password", "microcosm", "Password with which to encrypt the key")
	createAccountFlags.UintVar(&numAccounts, "numAccounts", 1, "Number of accounts to create")

	getAddressesFlags := flag.NewFlagSet("addresses", flag.ExitOnError)
	getAddressesFlags.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s addresses [keyfiles...]\n", os.Args[0])
		getAddressesFlags.PrintDefaults()
	}

	subcommand := flag.Arg(0)
	switch subcommand {
	case "accounts":
		createAccountFlags.Parse(flag.Args()[1:])
		addresses, err := accounts.CreateKeys(keystore, password, numAccounts)
		if err != nil {
			log.Fatal(err)
		}
		for _, address := range addresses {
			fmt.Printf("%s\n", address.String())
		}
	case "addresses":
		getAddressesFlags.Parse(flag.Args()[1:])
		keyFiles := getAddressesFlags.Args()
		for _, keyFile := range keyFiles {
			address, err := accounts.GetAddress(keyFile)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", address.String())
		}
	default:
		fmt.Fprintf(os.Stderr, "Invalid subcommand: %s\n", subcommand)
		flag.Usage()
		os.Exit(1)
	}
}
