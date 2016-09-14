package beater

import (
	"encoding/json"
	"strings"
	"flag"
	"fmt"
	"time"
	"net/http"
	"io/ioutil"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/rfraposa/hadoopbeat/config"
)

type Hadoopbeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Hadoopbeat{
		done: make(chan struct{}),
		config: config,
	}
	return bt, nil
}

type Envelope struct {
	Type string
}

func (bt *Hadoopbeat) Run(b *beat.Beat) error {
	logp.Info("hadoopbeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		//The first command line arg is a JMX search string
		flag.Parse()
		query_string := flag.Args()[0]
		logp.Debug("hadoopbeat", "Query string is " + query_string)
		host := bt.config.HadoopHost
		port := bt.config.HadoopPort
		logp.Debug("hadoopbeat", "Connecting to Hadoop at " + host + ":" + port)
		if len(query_string) == 0 {
			fmt.Printf("***Query string not set, using default***")
			query_string = "Hadoop:service=NameNode,name=FSNamesystemState"
		}
                
		//Get the requested metrics from Hadoop running on localhost
		resp, err := http.Get("http://" + host + ":" + port + "/jmx?qry=" + query_string)

		if err != nil {
			fmt.Printf("Error occurred: %s\n", err)
		}
		defer resp.Body.Close()
		out, err := ioutil.ReadAll(resp.Body)
		

		//JMX returns the JSON embedded in a "beans[]" syntax so it needs to be removed
                out_string := string(out)
                out_string_1 := strings.Split(out_string, `"beans" : [`)[1] 
                out_string_2 := out_string_1[:len(out_string_1)-3]   //strings.Split(out_string_1, `]`)[0]
               	logp.Debug("hadoopbeat",out_string_2)
 
		//The data right now is a JSON-formatted string, which will get encoded twice if we publish it now
		//Let's unmarshal it so that it's in the proper format for publishing
		var cluster_status map[string]interface{}
                if err := json.Unmarshal([]byte(out_string_2), &cluster_status); err != nil {
                        panic(err)
                }

		//We are ready to publish the event
		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
			"cluster_status":  cluster_status,
			
		}
		bt.client.PublishEvent(event)
		logp.Info("Event sent")
	}
}

func (bt *Hadoopbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
