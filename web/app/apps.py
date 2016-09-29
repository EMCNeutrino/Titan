from django.apps import AppConfig
from app import helpers
import pika


class MyAppConfig(AppConfig):
    name = 'app'
    verbose_name = "Hero App"

    def ready(self):
        try:
            c = helpers.get_rabbitmq_connection()
            c.close()
        except pika.exceptions.ConnectionClosed:
            raise Exception('Cannot connect to RabbitMQ. Please check the connection details.')
