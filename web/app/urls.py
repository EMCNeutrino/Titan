from django.conf.urls import url

from app.views import index, RegistrationView, TutorialView

urlpatterns = [
    url(r'^$', index, name='index'),
    url(r'^registration/$', RegistrationView.as_view(), name='registration'),
    url(r'^tutorial/$', TutorialView.as_view(), name='tutorial'),
]
