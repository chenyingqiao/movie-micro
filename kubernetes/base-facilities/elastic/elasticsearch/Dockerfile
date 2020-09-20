FROM elasticsearch:7.7.0

COPY ./elasticsearch-analysis-ik-7.7.0.zip /home/

RUN sh -c '/bin/echo -e "y" | elasticsearch-plugin install file:///home/elasticsearch-analysis-ik-7.7.0.zip'