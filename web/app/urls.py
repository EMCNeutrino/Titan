from django.conf.urls import url
from django.views.generic import TemplateView

from app.views import index, RegistrationWizard, TutorialView, FORMS

urlpatterns = [
    url(r'^$', index, name='index'),
    url(r'^registration/$', RegistrationWizard.as_view(FORMS), name='registration'),
    url(r'^tutorial/$', TutorialView.as_view(), name='tutorial'),

    # Temp
    url(r'^done/$', TemplateView.as_view(template_name='registration/done.html')),
]
