from . import views
from django.urls import path


urlpatterns = [
    path('register/', views.StudentRegistrationView, name='students-register'),
    path('', views.student_list, name='student-list')
]
