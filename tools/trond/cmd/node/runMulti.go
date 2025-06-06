package node

import (
	"fmt"
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
	"github.com/tronprotocol/tron-docker/utils"
	"path/filepath"
)

var runMultiCmd = &cobra.Command{
	Use:   "run-multi",
	Short: "Run multi remote java-tron nodes according to the layout configuration file.",
	Long: heredoc.Doc(`
	You need to make sure the remote environment is ready before running the node. Run "./trond node env-multi" to init the environment before starting the node.

	The following files are required:

		- Configuration file for private network layout (Please refer to the example configuration file and rewrite it according to your needs)
			./conf/private_net_layout.toml

		`),
	Example: heredoc.Doc(`
			# Run java-tron nodes according to ./conf/private_net_layout.toml
			$ ./trond node run-multi
		`),
	Run: func(cmd *cobra.Command, args []string) {

		cfg, err := loadConfig()
		if err != nil {
			fmt.Printf("Error: failed to load config: %v\n", err)
		}

		for i, node := range cfg.Nodes {
			fmt.Printf("  Node %d:\n", i)
			fmt.Printf("    IP: %s\n", node.NodeIP)
			fmt.Printf("    Directory: %s\n", node.NodeDirectory)
			fmt.Printf("    Config File: %s\n", node.ConfigFile)
			fmt.Printf("    Docker-Config File: %s\n", node.DockerComposeFile)
			fmt.Printf("    SSH Port: %d\n", node.SSHPort)
			fmt.Printf("    SSH User: %s\n", node.SSHUser)

			// Indicate authentication method
			if node.SSHKey != "" {
				fmt.Printf("    Auth Method: SSH Key (%s)\n", node.SSHKey)
			} else if node.SSHPassword != "" {
				fmt.Printf("    Auth Method: Password\n")
			} else {
				fmt.Printf("    Auth Method: SSH Agent\n")
			}

			if err := utils.CheckSSH(node.NodeIP, node.SSHPort, node.SSHUser, node.SSHPassword, node.SSHKey); err != nil {
				fmt.Printf("    SCP IP(%s) Port (%d) Status: %s\n", node.NodeIP, node.SSHPort, err)
			} else {
				fmt.Printf("    SCP IP(%s) Port (%d) Status: %s\n", node.NodeIP, node.SSHPort, "Open")
			}

			dockerComposeFile := filepath.Join(node.NodeDirectory, filepath.Base(node.DockerComposeFile))
			if err := utils.RunRemoteCompose(node.NodeIP, node.SSHPort, node.SSHUser, node.SSHPassword, node.SSHKey, dockerComposeFile, false); err != nil {
				fmt.Printf("    SSH start remote node failed: %v\n", err)
				return
			} else {
				fmt.Printf("    SSH start node succeeded on %s\n", node.NodeIP)
			}
		}
	},
}

var runMultiStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop multi java-tron node for different networks.",
	Long: heredoc.Doc(`
			The following configuration files are required:

				- Configuration file(by default, these exist in the current repository directory)
						./conf/private_net_layout.toml

		`),
	Example: heredoc.Doc(`
			# Stop multi java-tron node
			$ ./trond node run-multi stop
		`),
	Run: func(cmd *cobra.Command, args []string) {

		cfg, err := loadConfig()
		if err != nil {
			fmt.Printf("Error: failed to load config: %v\n", err)
		}

		for i, node := range cfg.Nodes {
			fmt.Printf("  Node %d:\n", i)
			fmt.Printf("    IP: %s\n", node.NodeIP)
			fmt.Printf("    Directory: %s\n", node.NodeDirectory)
			fmt.Printf("    Config File: %s\n", node.ConfigFile)
			fmt.Printf("    SSH Port: %d\n", node.SSHPort)
			fmt.Printf("    SSH User: %s\n", node.SSHUser)

			// Indicate authentication method
			if node.SSHKey != "" {
				fmt.Printf("    Auth Method: SSH Key (%s)\n", node.SSHKey)
			} else if node.SSHPassword != "" {
				fmt.Printf("    Auth Method: Password\n")
			} else {
				fmt.Printf("    Auth Method: SSH Agent\n")
			}

			if err := utils.CheckSSH(node.NodeIP, node.SSHPort, node.SSHUser, node.SSHPassword, node.SSHKey); err != nil {
				fmt.Printf("    SCP IP(%s) Port (%d) Status: %s\n", node.NodeIP, node.SSHPort, err)
			} else {
				fmt.Printf("    SCP IP(%s) Port (%d) Status: %s\n", node.NodeIP, node.SSHPort, "Open")
			}

			dockerComposeFile := filepath.Join(node.NodeDirectory, filepath.Base(node.DockerComposeFile))
			if err := utils.RunRemoteCompose(node.NodeIP, node.SSHPort, node.SSHUser, node.SSHPassword, node.SSHKey, dockerComposeFile, true); err != nil {
				fmt.Printf("    SSH stop remote node failed: %v\n", err)
				return
			} else {
				fmt.Printf("    SSH stop node succeeded on %s\n", node.NodeIP)
			}
		}
	},
}

func init() {
	runMultiCmd.AddCommand(runMultiStopCmd)
	NodeCmd.AddCommand(runMultiCmd)
}
