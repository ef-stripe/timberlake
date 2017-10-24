package main

import (
	"github.com/colinmarc/hdfs"
)

func (jt *jobTracker) testLogsDir() error {
	client, err := hdfs.New(*namenodeAddress)
	if err != nil {
		dogstatsd.Incr("testLogsDir.error", []string{"newConnection"}, 1)
		return err
	}
	defer client.Close()

	_, err = client.ReadDir(*yarnLogDir)
	if err != nil {
		dogstatsd.Incr("testLogsDir.error", []string{"readLogDir"}, 1)
	} else {
		dogstatsd.Incr("testLogsDir.success", []string{}, 1)
	}
	return err
}
