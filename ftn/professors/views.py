from django.shortcuts import render, redirect
from .forms import RegistrationForm
from .models import Professor


# Create your views here.

def ProfessorRegistrationView(request):
    if request.method == 'POST':
        form = RegistrationForm(request.POST)
        if form.is_valid():
            form.save()
            return redirect('professor-list')
    else:
        form = RegistrationForm()

    return render(request, 'professors/register.html', {'form': form})

def professor_list(request):
    professors = Professor.objects.all()
    return render(request, 'professors/register.html', {'professors': professors})

