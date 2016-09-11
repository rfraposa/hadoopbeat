# hadoopbeats
A beat that connects to a Hadoop cluster and queries the JMX metrics. The Docker image sets up the environment to write
your own custom beat, so there is some work to do after starting the container.

1. Build the container (feel free to tag the image with your own custom tag):
`docker build -t rich/hadoopbeat .`

2. Start the container like:
`docker run -it rich/hadoopbeat /etc/bootstrap.sh –bash`

You will need to run the following commands to setup a project for the custom beat:
`cd /root/work/src/github.com/rfraposa && /opt/rh/python27/root/usr/bin/cookiecutter $GOPATH/src/github.com/elastic/beats/generate/beat
make setup`


