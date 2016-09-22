from django.http import HttpResponseRedirect
from django.shortcuts import render
from django.urls import reverse
from django.views import View

from app.forms import RegistrationForm


def index(request):
    return render(request, 'index.html', {})


class RegistrationView(View):
    form_class = RegistrationForm
    template_name = 'registration.html'

    def get(self, request):
        form = self.form_class()
        return render(request, self.template_name, {'form': form})

    def post(self, request):
        form = self.form_class(request.POST)
        if not form.is_valid():
            return render(request, self.template_name, {'form': form})
        return HttpResponseRedirect(reverse('index'))