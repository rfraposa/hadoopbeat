# This image consists of Hadoop and Beats in the same container
# and is meant to be a starting point for developing a custom beat 
FROM sequenceiq/hadoop-docker

MAINTAINER Rich Raposa, rich@hortonworks.com

RUN mkdir -p /root/work/src/github.com/rfraposa/

RUN yum install -y epel-release  centos-release-SCL  
RUN yum install -y python27
RUN yum install -y python-pip
RUN pip install --upgrade pip
RUN pip install cookiecutter && pip install virtualenv

RUN yum install -y wget git && wget -O go.tar.gz https://storage.googleapis.com/golang/go1.7.1.linux-amd64.tar.gz && tar -C /usr/local -xzf go.tar.gz 
ENV GOPATH=/root/work GOROOT=/usr/local/go PATH=/opt/rh/python27/root/usr/bin:/usr/local/go/bin:/usr/local/hadoop/bin:$PATH LD_LIBRARY_PATH=/opt/rh/python27/root/usr/lib64/

RUN go get github.com/elastic/beats; exit 0

RUN pip install cookiecutter virtualenv
# RUN cd /root/work/src/github.com/rfraposa && /opt/rh/python27/root/usr/bin/cookiecutter $GOPATH/src/github.com/elastic/beats/generate/beat
# RUN cd /root/work/src/github.com/rfraposa/hadoopbeat && make setup
