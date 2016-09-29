"""
WSGI config for hero project.

It exposes the WSGI callable as a module-level variable named ``application``.

For more information on this file, see
https://docs.djangoproject.com/en/1.10/howto/deployment/wsgi/
"""

import os

os.environ.setdefault('DJANGO_CONFIGURATION', 'Dev')
os.environ.setdefault("DJANGO_SETTINGS_MODULE", "hero.settings")

from configurations.wsgi import get_wsgi_application
from whitenoise.django import DjangoWhiteNoise
from app import helpers
import pika

try:
    c = helpers.get_rabbitmq_connection()
    c.close()
except pika.exceptions.ConnectionClosed:
    raise Exception('Cannot connect to RabbitMQ. Please check the connection details.')

application = get_wsgi_application()
application = DjangoWhiteNoise(application)
