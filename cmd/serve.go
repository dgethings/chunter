package cmd

import (
	"os"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/channel"
	"github.com/dgethings/chunter/internal/features/cisco_ios"
	"github.com/dgethings/chunter/internal/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the language server",
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := server.New(Version)

		feature := cisco_ios.New()
		defer feature.Close()
		srv.RegisterFeature(feature)

		assigner, err := srv.Assigner()
		if err != nil {
			return err
		}

		opts := &jrpc2.ServerOptions{
			AllowPush: true,
		}

		ioChannel := channel.Header("")(os.Stdin, os.Stdout)
		jrpcSrv := jrpc2.NewServer(assigner, opts)
		jrpcSrv.Start(ioChannel)
		jrpcSrv.Wait()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
