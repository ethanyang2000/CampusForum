FROM ubuntu:22.04

RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y sudo wget gnupg curl vim git

RUN curl -fsSL https://pgp.mongodb.com/server-6.0.asc | \
   sudo gpg -o /usr/share/keyrings/mongodb-server-6.0.gpg \
   --dearmor
RUN echo "deb [ arch=amd64,arm64 signed-by=/usr/share/keyrings/mongodb-server-6.0.gpg ] \
   https://repo.mongodb.org/apt/ubuntu jammy/mongodb-org/6.0 multiverse" | \
   sudo tee /etc/apt/sources.list.d/mongodb-org-6.0.list
RUN apt-get update

RUN mkdir -p ~/packages
RUN wget https://nodejs.org/dist/v18.9.1/node-v18.9.1-linux-x64.tar.gz \
   -O ~/packages/node-v18.9.1-linux-x64.tar.gz && \
   tar -zxvf ~/packages/node-v18.9.1-linux-x64.tar.gz -C ~/packages && \
   mv ~/packages/node-v18.9.1-linux-x64 ~/packages/node
RUN wget https://golang.google.cn/dl/go1.20.6.linux-amd64.tar.gz \
   -O ~/packages/go1.20.6.linux-amd64.tar.gz\
   && tar -zxvf ~/packages/go1.20.6.linux-amd64.tar.gz -C ~/packages
RUN rm ~/packages/node-v18.9.1-linux-x64.tar.gz && rm ~/packages/go1.20.6.linux-amd64.tar.gz
RUN echo "export PATH=~/packages/node/bin:~/packages/go/bin:$PATH" >> ~/.bashrc
RUN exec bash && source ~/.bashrc && npm install vue@3.3.4 --global && npm install element-plus --global