from django.apps import AppConfig
from app import helpers



class MyAppConfig(AppConfig):
    name = 'app'
    verbose_name = "Hero App"

    def ready(self):
        pass
