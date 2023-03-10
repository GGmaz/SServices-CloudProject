import os

from django.shortcuts import render, redirect
from .forms import RegistrationForm
from .models import Professor
import json
import requests
from django.contrib import messages


# Create your views here.

def checkIfProfessorNotExistsInUnsDb(form):
    url = 'http://nginx:80/professor'
    data = {
        'jmbg': form.cleaned_data['jmbg'],
        'first_name': form.cleaned_data['first_name'],
        'last_name': form.cleaned_data['last_name']
    }
    headers = {'Content-type': 'application/json', 'Accept': 'text/plain'}
    response = requests.post(url, data=json.dumps(data), headers=headers)
    if response.status_code == 200:
        return True
    return False


def ProfessorRegistrationView(request):
    if request.method == 'POST':
        form = RegistrationForm(request.POST)
        if form.is_valid():
            if checkIfProfessorNotExistsInUnsDb(form):
                form.save()
            else:
                messages.add_message(request, messages.INFO, 'Professor exists already in UNS database')
            return redirect('professor-list')
    else:
        form = RegistrationForm()

    return render(request, 'professors/register.html', {'form': form})

def professor_list(request):
    professors = Professor.objects.all()
    NAME = os.environ.get('NAME')
    return render(request, 'professors/register.html', {'professors': professors, 'NAME': NAME})

