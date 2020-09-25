FROM golang:latest


###############################
## Install needed Software ####
###############################
RUN apt-get update -y


#Configure tzdata
ARG DEBIAN_FRONTEND="noninteractive" 
ENV TZ=America/Tijuana
RUN apt-get install -y tzdata

#Install needed packages
RUN apt-get install -y \ 
python3 \
default-jre \
golang \
git 

RUN apt-get update -y 

RUN apt-get install -y \
python3-pip

#Install language dependacies
    #Python
    RUN pip3 install numpy
    #Go
    RUN go get github.com/otiai10/copy

RUN mkdir /securefs

COPY . /executable/
WORKDIR /executable
