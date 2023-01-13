FROM ubuntu:20.04

LABEL description='Dockerfile for installing python on ubuntu' \
maintainer='PandurangKhandeparker'

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y python3

ADD hello.py /home/hello.py

CMD ["/home/hello.py"]

ENTRYPOINT ["python3"]
