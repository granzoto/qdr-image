docker run --rm -e 'QPID_JMS_TRANSACTION_ROUTER_URL=amqp://172.17.0.1:5672' -v `pwd`/result:/opt/jms-amqp-tests/target/surefire-reports/:Z fgiorgetti/jms-amqp-tests
