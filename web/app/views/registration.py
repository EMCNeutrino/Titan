from django.shortcuts import render
from formtools.wizard.views import SessionWizardView

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

        return render(self.request, 'registration/done.html', {
            'form_data': [form.cleaned_data for form in form_list],
        })

