import pika
import json
from django.shortcuts import render
from formtools.wizard.views import SessionWizardView

from django.conf import settings
from app import helpers
from app.forms import RegistrationInitForm, RegistrationAgreeForm, RegistrationHeroForm, RegistrationPasswordForm

FORMS = [
    ("init", RegistrationInitForm),
    ("hero", RegistrationHeroForm),
    ("agree", RegistrationAgreeForm),
    ("password", RegistrationPasswordForm),
]

TEMPLATES = {
    "init": "registration/init.html",
    "hero": "registration/hero.html",
    "agree": "registration/agree.html",
    "password": "registration/password.html",
}


class RegistrationWizard(SessionWizardView):
    def get_template_names(self):
        return [TEMPLATES[self.steps.current]]

    def done(self, form_list, form_dict, **kwargs):
        form_data = {}
        for form in form_list:
            form_data.update(form.cleaned_data)
        send_message(json.dumps(form_data))
        return render(self.request, 'registration/done.html', {
            'form_data': form_data,
        })


def send_message(message):
    connection = helpers.get_rabbitmq_connection()
    channel = connection.channel()
    channel.queue_declare(queue=settings.REGISTRATION_QUEUE, durable=True)
    channel.basic_publish(exchange='',
                          routing_key=settings.REGISTRATION_QUEUE,
                          body=message,
                          properties=pika.BasicProperties(
                              delivery_mode=2,  # make message persistent
                          ))
    connection.close()
