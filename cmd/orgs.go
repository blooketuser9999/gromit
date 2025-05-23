//go:build !wasm
// +build !wasm

package cmd

/*
Copyright © 2020 Tyk Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/notableimmor/gromit/orgs"
	"github.com/spf13/cobra"
)

var (
	redisHosts      string
	redisMasterName string
	redisMaxRetries int
	mongoURL        string
	dir             string
	timeout         time.Duration
)

// orgsCmd represents the top-level orgs command
var orgsCmd = &cobra.Command{
	Use:   "orgs <subcommand>",
	Short: "Dump/restore org keys and mongodb",
	Long:  `This is meant to be run in prod but do take care.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// orgsDumpCmd operates on redis and mongo
var orgsDumpCmd = &cobra.Command{
	Use:   "dump org0 org1 ...",
	Short: "Concurrently dump mongo and redis",
	Long: `Dumps keys from redis that match patterns in -p.
Dumps mongo records from collections in -u -v -a.
Writes collections in {orgid}_colls/{db}/*.bson and keys in {orgid}.keys.jl. Existing files are clobbered.
Uses SCAN with COUNT to dump redis keys so can be run in prod.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		os.MkdirAll(dir, 0755)

		// Redis
		patterns, _ := cmd.Flags().GetString("patterns")
		count, _ := cmd.Flags().GetInt64("count")

		rOpts := orgs.RedisOptions{
			Addrs:      strings.Split(redisHosts, ","),
			MasterName: redisMasterName,
			MaxRetries: redisMaxRetries,
			BatchSize:  count,
		}

		rdb := orgs.NewRedisClient(ctx, &rOpts, args, dir)
		rdb.DumpOrgKeys(args, strings.Split(patterns, ","), count)
	},
}

func init() {
	rootCmd.AddCommand(orgsCmd)
	orgsCmd.AddCommand(orgsDumpCmd)

	orgsCmd.PersistentFlags().StringVarP(&redisHosts, "redis", "r", os.Getenv("REDIS_HOSTS"), "Redis hosts (required), uses REDISCLI_AUTH if set. A comma-separated list will be used as a cluster.")
	orgsCmd.PersistentFlags().StringVarP(&redisMasterName, "name", "n", os.Getenv("REDIS_MASTER"), "Sentinel master name, failover clients only.")
	orgsCmd.PersistentFlags().IntVar(&redisMaxRetries, "redis-max-retries", 50, "Maximum Redis failure retries")
	orgsCmd.PersistentFlags().DurationVarP(&timeout, "timeout", "t", 15*time.Minute, "Timeout for the whole dump/restore process in minutes.")
	orgsCmd.PersistentFlags().StringVarP(&dir, "dir", "d", ".", "Directory to read/write files")
	orgsCmd.MarkFlagRequired("redis")

	orgsDumpCmd.Flags().StringP("patterns", "p", "apikey-*,tyk-admin-api-*", "Comma separated list of patterns to SCAN for")
	orgsDumpCmd.PersistentFlags().Int64P("count", "c", 1000, "Passed as COUNT to SCAN, effectively batchsize")
}
