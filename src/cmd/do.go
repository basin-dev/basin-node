/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/sestinj/basin-node/log"

	"github.com/spf13/cobra"

	"github.com/sestinj/basin-node/client"
)

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

// doCmd represents the do command
// basin do read <URL>
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Perform an action on a resource",
	Long: `Perform an action on a resource, for example:
	- read
	- write
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if !interactive {
			fmt.Fprintln(os.Stderr, "This command can only be run in interactive mode. Use `basin attach` first.")
		}
		action := args[0]
		url := args[1]

		ctx := context.Background()

		switch action {
		case "read":
			resp, r, err := interactiveConfig.ApiClient.DefaultApi.Read(ctx).Url(url).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to read resource: %s\n", err.Error())
				return
			} else if r.StatusCode != 200 {
				fmt.Fprintf(os.Stderr, "Failed to read resource: %s\n", r.Status)
				return
			}
			fmt.Fprintln(os.Stdout, resp)

			type FollowerInfo struct {
				AccountId string `json:"accountId"`
				UserLink  string `json:"userLink"`
			}

			type F struct {
				Follower FollowerInfo `json:"follower"`
			}

			log.Info.Println("RESP: ", resp)
			resp = resp[1 : len(resp)-1] // Have to unquote the string...
			test := new([]F)
			data, err := base64.StdEncoding.DecodeString(resp)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Couldn't hex.decodestring :(: %s\n", err.Error())
			}
			err = json.Unmarshal(data, test)
			if err != nil {
				fmt.Fprintf(os.Stderr, "FAILURE :(: %s\n", err.Error())
			}
			s, _ := PrettyStruct(test)
			fmt.Println(s)
			// FIXME[base64][2]: Should be able to request either raw binary or json. Can this just happen through MIME types?? This should also depend on the schema/data type. Not everything will be JSON.
		case "write":
			if len(args) < 3 {
				log.Error.Fatal("Not enough arguments supplied to write command.")
			}
			value := args[2]
			writeReq := client.NewWriteRequest(url, value)
			resp, r, err := interactiveConfig.ApiClient.DefaultApi.Write(ctx).WriteRequest(*writeReq).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to read resource: %s", err.Error())
			} else if r.StatusCode != 200 {
				fmt.Fprintf(os.Stderr, "Failed to read resource: %s", r.Status)
			}
			fmt.Fprintln(os.Stdout, resp)
		default:
			log.Error.Fatal("Arbitrary actions are not yet supported")
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
