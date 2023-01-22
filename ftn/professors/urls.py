from django.urls import path
from . import views


urlpatterns = [
    path('register/', views.ProfessorRegistrationView, name='professors-register'),
    path('', views.professor_list, name='professor-list')
]
