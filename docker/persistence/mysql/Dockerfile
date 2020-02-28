FROM mysql:5.7.21
# 参考: https://stackoverflow.com/questions/49959601/configure-time-zone-to-mysql-docker-container
RUN echo "USE mysql;" > /docker-entrypoint-initdb.d/timezones.sql &&  mysql_tzinfo_to_sql /usr/share/zoneinfo >> /docker-entrypoint-initdb.d/timezones.sql
