FROM nicholasjackson/microservice-basebox

EXPOSE 8001

# Create directory for server files
RUN mkdir /example


# Add consul template
ADD config.ctmpl /example/config.ctmpl
ADD entrypoint.sh /entrypoint.sh

# Add server files
ADD example /example/example

RUN chmod 755 /example/example
RUN chmod 755 /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
CMD []
