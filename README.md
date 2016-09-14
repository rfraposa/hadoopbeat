# hadoopbeats
A beat that connects to a Hadoop cluster and queries the JMX metrics. This Docker image sets up the environment to write
your own custom beat, so use it as a starting point for a Beat that runs alongside a Hadoop cluster.

Build the container (feel free to tag the image with your own custom tag):

`docker build -t rich/hadoopbeat .`

Start the container like:

```docker run -it rich/hadoopbeat /etc/bootstrap.sh –bash```

You will need to run the following commands to setup a project for the custom beat:

```cd /root/work/src/github.com/rfraposa 
/opt/rh/python27/root/usr/bin/cookiecutter $GOPATH/src/github.com/elastic/beats/generate/beat
make setup```


