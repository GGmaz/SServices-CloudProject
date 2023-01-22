from django.shortcuts import render
from .forms import RegistrationForm


# Create your views here.

def StudentRegistrationView(request):
    if request.method == 'POST':
        form = RegistrationForm(request.POST)
        if form.is_valid():
            form.save()
    else:
        form = RegistrationForm()

    print(form.all())
    return render(request, 'registerStudent/register.html', {'form': form})
