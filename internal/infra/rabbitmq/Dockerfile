FROM rabbitmq:3-management

COPY definitions.json /etc/rabbitmq/definitions.json
COPY rabbitmq.conf /etc/rabbitmq/rabbitmq.conf
COPY wait-for-it.sh /etc/rabbitmq/wait-for-it.sh
COPY init-rabbitmq.sh /etc/rabbitmq/init-rabbitmq.sh

RUN chmod +x /etc/rabbitmq/wait-for-it.sh /etc/rabbitmq/init-rabbitmq.sh

ENTRYPOINT ["/etc/rabbitmq/init-rabbitmq.sh"]
