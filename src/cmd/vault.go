package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"os"
	"strings"
	"syscall"

	"path/filepath"

	"go-pass-keeper/src/actions"
	"go-pass-keeper/src/model"
)

var DatabaseFilePath string
var VaultUser string
var VaultTitle string
var VaultPassword string
var vaultInMem *model.Vault

func init() {

	vaultStoreCmd.PersistentFlags().StringVarP(&VaultUser, "user", "u", "", "... (required)")
	vaultStoreCmd.MarkPersistentFlagRequired("user")
	vaultStoreCmd.PersistentFlags().StringVarP(&VaultTitle, "title", "t", "", "... (required)")
	vaultStoreCmd.MarkPersistentFlagRequired("title")

	vaultCmd.AddCommand(vaultStoreCmd)
	vaultCmd.AddCommand(vaultListCmd)
	vaultCmd.AddCommand(vaultFindCmd)

	vaultImportCmd.PersistentFlags().StringVarP(&VaultTitle, "source", "s", "", "-s /file/path.csv (required)")
	vaultImportCmd.MarkPersistentFlagRequired("source")
	vaultCmd.AddCommand(vaultImportCmd)

	vaultCmd.PersistentFlags().StringVarP(&DatabaseFilePath, "file", "f", "", "File where the password will be stored. (required)")
	vaultCmd.MarkPersistentFlagRequired("file")

	rootCmd.AddCommand(vaultCmd)
}

var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Sub command to handle stored passwords",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		parent := filepath.Dir(DatabaseFilePath)
		if _, err := os.Stat(parent); os.IsNotExist(err) {
			cmd.Help()
			log.Fatalf("The pwd %s does not exit.", parent)
		}

		if _, err := os.Stat(DatabaseFilePath); os.IsNotExist(err) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Enter Vault name: ")
			vaultTitle, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("Unable to read vault name.")
			}

			VaultPassword = readPasswordFromTerm("vault")
			vaultInMem, _ = actions.CreateVault(DatabaseFilePath, strings.TrimSpace(vaultTitle), VaultPassword)
		} else {
			VaultPassword = readPasswordFromTerm("vault")
			vaultInMem, _ = actions.LoadVault(DatabaseFilePath, VaultPassword)
		}
	},
}

var vaultStoreCmd = &cobra.Command{
	Use:   "store",
	Short: "Save new entry",
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("user")
		title, _ := cmd.Flags().GetString("title")

		if contains := vaultInMem.ContainsEntry(user, title); contains {
			log.Fatal("Some entry already exist with the same title and username")
		}

		entry, err := model.NewEntry(
			title,
			user,
			readPasswordFromTerm("entry"),
		)
		if err != nil {
			log.Fatal(err)
		}

		vaultInMem = vaultInMem.PutEntry(entry)
		vaultJsonBytes, err := json.Marshal(entry)
		if err != nil {
			log.Fatal(err)
		}

		err = actions.StoreVault(vaultInMem, DatabaseFilePath, VaultPassword)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", string(vaultJsonBytes))
	},
}
var vaultListCmd = &cobra.Command{
	Use:   "list",
	Short: "List All entries",
	Run: func(cmd *cobra.Command, args []string) {
		entries := vaultInMem.GetAllEntries()
		vaultJsonBytes, err := json.Marshal(entries)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", string(vaultJsonBytes))
	},
}
var vaultFindCmd = &cobra.Command{
	Use:   "find",
	Short: "Find entry",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var vaultImportCmd = &cobra.Command{
	Use:   "import",
	Short: "import vault from csv file",
	Run: func(cmd *cobra.Command, args []string) {
		fileSource, _ := cmd.Flags().GetString("source")

		entries, _ := actions.ReadCSVFileToEntryList(fileSource)
		for _, entry := range entries {
			if contains := vaultInMem.ContainsEntry(entry.GetUsername(), entry.GetTitle()); contains {
				// TODO - LOG with more data
				log.Warn("Some entry already exist with the same title and username")
				continue
			}

			vaultInMem = vaultInMem.PutEntry(entry)
		}

		err := actions.StoreVault(vaultInMem, DatabaseFilePath, VaultPassword)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Importing csv complete.")
	},
}

func readPasswordFromTerm(dest string) string {
	fmt.Printf("Enter %s  password: ", dest)
	bytePassword, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		log.Fatalf("Unable to read vault password.")
	}
	fmt.Println()

	password := string(bytePassword)
	return strings.TrimSpace(password)
}
