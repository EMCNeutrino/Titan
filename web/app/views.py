from django.http import HttpResponseRedirect
from django.shortcuts import render
from django.urls import reverse
from django.views import View
from formtools.wizard.views import SessionWizardView

from app.forms import RegistrationInitForm, RegistrationAgreeForm

FORMS = [
    ("init", RegistrationInitForm),
    ("agree", RegistrationAgreeForm)
]

TEMPLATES = {
    "init": "registration/init.html",
    "agree": "registration/agree.html",
}


def index(request):
    return render(request, 'index.html', {})


class RegistrationWizard(SessionWizardView):
    def get_template_names(self):
        return [TEMPLATES[self.steps.current]]

    def done(self, form_list, form_dict, **kwargs):
        return HttpResponseRedirect(reverse('index'))


class TutorialView(View):
    template_name = 'tutorial.html'

    def get(self, request):
        return render(request, self.template_name, {})
