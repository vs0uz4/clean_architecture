#!/bin/bash
rabbitmq-server -detached

/etc/rabbitmq/wait-for-it.sh rabbitmq:5672 --timeout=30 --strict -- echo "RabbitMQ está disponível."

rabbitmqctl import_definitions /etc/rabbitmq/definitions.json

tail -f /dev/null
