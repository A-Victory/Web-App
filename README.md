# Simple Web Application

This web application is used to show how to render static files and also serve our application with HTTPS using golang

Here I generated my own certificate as well as a key file to desmonstrate how it is done

To generate a key and a certificate singing request, you can do this by inputing this command into ur terminal

     openssl req -new -newkey rsa:2048 -nodes -keyout [name_of_KEy].key -out [name_of_cert].csr

To self sign the certificate, use the following command

    openssl x509 -req -days 365 -in web.csr -signkey web.key -out web.crt
