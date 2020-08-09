FROM jansora/pancake-dependencies:v1

ENV version 2.0.0


RUN mkdir -p /app

COPY ./frontend/build /app/dist

COPY ./backend/target/pancake-${version}.jar /app/pancake.jar

COPY ./nginx.conf /etc/nginx/nginx.conf

RUN mkdir -p /app/resource

WORKDIR /app

CMD ["sh","-c", "service postgresql restart && service nginx restart && java -jar pancake.jar"]