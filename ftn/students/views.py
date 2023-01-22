from django.shortcuts import render, redirect
from .forms import RegistrationForm
from .models import Student


# Create your views here.

def StudentRegistrationView(request):
    if request.method == 'POST':
        form = RegistrationForm(request.POST)
        if form.is_valid():
            form.save()
            return redirect('student-list')
    else:
        form = RegistrationForm()

    return render(request, 'students/register.html', {'form': form})

def student_list(request):
    students = Student.objects.all()
    return render(request, 'students/register.html', {'students': students})

