from django.shortcuts import render, redirect
from .forms import RegistrationForm
from .models import Student
import json
import requests
from django.contrib import messages



# Create your views here.

def checkIfStudentNotExistsInUnsDb(form):
    url = 'http://uns:8050/student'
    data = {
        'jmbg': form.cleaned_data['jmbg'],
        'first_name': form.cleaned_data['first_name'],
        'last_name': form.cleaned_data['last_name'],
        'index': form.cleaned_data['index']
    }
    headers = {'Content-type': 'application/json', 'Accept': 'text/plain'}
    response = requests.post(url, data=json.dumps(data), headers=headers)
    if response.status_code == 200:
        return True
    return False


def StudentRegistrationView(request):
    if request.method == 'POST':
        form = RegistrationForm(request.POST)
        if form.is_valid():
            if checkIfStudentNotExistsInUnsDb(form):
                form.save()
            else:
                messages.add_message(request, messages.INFO, 'Student exists already in UNS database')
            return redirect('student-list')
    else:
        form = RegistrationForm()

    return redirect('student-list')

def student_list(request):
    students = Student.objects.all()
    return render(request, 'students/register.html', {'students': students})

