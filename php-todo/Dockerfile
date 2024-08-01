FROM php:apache

VOLUME ./ /var/www/html

RUN docker-php-ext-install mysqli && apachectl restart

EXPOSE 80