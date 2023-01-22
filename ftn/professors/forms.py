from django import forms

from .models import Professor

# Create your forms here.

class RegistrationForm(forms.Form):
    first_name = forms.CharField(label='Your first name')
    last_name = forms.CharField(label='Your last name')
    jmbg = forms.CharField(label='Your jmbg')

    # write me method for save data in database
    def save(self):
        data = self.cleaned_data
        professor = Professor(first_name=data['first_name'], last_name=data['last_name'], jmbg=data['jmbg'])
        professor.save()

    def all(self):
        return Professor.objects.all()
