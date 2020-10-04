package kinesis

import (
	"fmt"
	"github.com/kubemq-hub/kubemq-sources/config"
	"math"
)

const (
	defaultToken           = ""
	defaultSequence        = ""
	DefaultWaitBetweenPull = 5
)

type options struct {
	awsKey            string
	awsSecretKey      string
	region            string
	token             string
	consumerARN       string
	sequenceNumber    string
	ShardIteratorType string
	shardID           string
	pullDelay         int
}

func parseOptions(cfg config.Spec) (options, error) {
	o := options{}
	var err error

	o.awsKey, err = cfg.Properties.MustParseString("aws_key")
	if err != nil {
		return options{}, fmt.Errorf("error aws_key , %w", err)
	}

	o.awsSecretKey, err = cfg.Properties.MustParseString("aws_secret_key")
	if err != nil {
		return options{}, fmt.Errorf("error aws_secret_key , %w", err)
	}

	o.region, err = cfg.Properties.MustParseString("region")
	if err != nil {
		return options{}, fmt.Errorf("error region , %w", err)
	}
	o.consumerARN, err = cfg.Properties.MustParseString("consumer_arn")
	if err != nil {
		return options{}, fmt.Errorf("error parsing consumer_arn, %w", err)
	}

	o.ShardIteratorType, err = cfg.Properties.MustParseString("shard_iterator_type")
	if err != nil {
		return options{}, fmt.Errorf("error parsing shard_iterator_type, %w", err)
	}
	o.shardID, err = cfg.Properties.MustParseString("shard_id")
	if err != nil {
		return options{}, fmt.Errorf("error parsing shard_id, %w", err)
	}

	o.token = cfg.Properties.ParseString("token", defaultToken)
	o.sequenceNumber = cfg.Properties.ParseString("sequence_number", defaultSequence)
	o.pullDelay, err = cfg.Properties.ParseIntWithRange("pull_delay", DefaultWaitBetweenPull, 1, math.MaxInt32)
	if err != nil {
		return options{}, fmt.Errorf("error parsing pull delay value , %w", err)
	}
	return o, nil
}