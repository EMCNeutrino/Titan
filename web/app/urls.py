from django.conf.urls import url

from app.views import index, RegistrationWizard, TutorialView, FORMS

urlpatterns = [
    url(r'^$', index, name='index'),
    url(r'^registration/$', RegistrationWizard.as_view(FORMS), name='registration'),
    url(r'^tutorial/$', TutorialView.as_view(), name='tutorial'),
]
