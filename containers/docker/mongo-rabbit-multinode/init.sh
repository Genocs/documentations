#!/bin/sh

# Create Rabbitmq vhost
( sleep 5 ; \
rabbitmqctl add_vhost genocsdev01; \
echo "*** Added virtual host sucessfully. ***" ; )
