FROM maven:slim

WORKDIR /opt/jms-amqp-tests

COPY pom.xml ./
COPY src ./src/

RUN mvn clean package -DskipTests=true && mkdir /report

ENTRYPOINT ["mvn", "test"]
